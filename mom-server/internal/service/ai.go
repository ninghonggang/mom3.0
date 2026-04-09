package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
	"mom-server/internal/service/ai_schema"

	"github.com/google/uuid"
)

// AIIntent AI解析的操作意图
type AIIntent struct {
	OperationType string                 `json:"operation_type"` // query | write | analysis | navigation | unknown
	Module        string                 `json:"module"`          // e.g., production, wms, quality
	Action        string                 `json:"action"`          // e.g., list_orders, create_material
	Parameters    map[string]interface{} `json:"parameters"`
	NaturalDesc   string                 `json:"natural_desc"` // human-readable description
	SQLHint       string                 `json:"sql_hint,omitempty"`
	Confidence    float64                `json:"confidence"` // 0.0-1.0
}

// ChatResponse AI聊天响应
type ChatResponse struct {
	SessionID      string       `json:"session_id"`
	MessageID     int64        `json:"message_id"`
	Role          string       `json:"role"`
	Content       string       `json:"content"`
	ResultType    string       `json:"result_type"` // table | chart | text | form | confirmation | error
	ResultData    interface{}  `json:"result_data,omitempty"`
	OperationType string       `json:"operation_type,omitempty"`
	Operation     *AIIntent    `json:"operation,omitempty"`
	NeedsConfirm  bool         `json:"needs_confirmation"`
	Metadata      *ResponseMeta `json:"metadata,omitempty"`
}

// ResponseMeta 响应元数据
type ResponseMeta struct {
	TokensUsed int64   `json:"tokens_used,omitempty"`
	Model      string  `json:"model,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
}

// AIService AI代理服务
type AIService struct {
	configRepo *repository.AIConfigRepository
	convRepo   *repository.ConversationRepository
	msgRepo    *repository.MessageRepository
	httpClient *http.Client
	schemaGen  *ai_schema.SchemaGenerator
}

// NewAIService 创建AI服务
func NewAIService(
	configRepo *repository.AIConfigRepository,
	convRepo *repository.ConversationRepository,
	msgRepo *repository.MessageRepository,
) *AIService {
	return &AIService{
		configRepo: configRepo,
		convRepo:   convRepo,
		msgRepo:    msgRepo,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
		schemaGen: ai_schema.NewSchemaGenerator(),
	}
}

// SendMessage 处理用户发送的消息
func (s *AIService) SendMessage(ctx context.Context, tenantID, userID int64, sessionID, message string, history []*model.ChatMessage) (*ChatResponse, error) {
	// 获取或创建会话
	var conv *model.ChatConversation
	var err error

	if sessionID != "" {
		conv, err = s.convRepo.GetBySessionID(ctx, tenantID, sessionID)
		if err != nil {
			// 会话不存在，创建新会话
			conv = nil
		}
	}

	if conv == nil {
		// 创建新会话
		newSessionID := uuid.New().String()
		conv = &model.ChatConversation{
			TenantID:  tenantID,
			UserID:    userID,
			SessionID: newSessionID,
		}
		if err := s.convRepo.Create(ctx, conv); err != nil {
			return nil, fmt.Errorf("创建会话失败: %w", err)
		}
		sessionID = newSessionID
	}

	// 保存用户消息
	userMsg := &model.ChatMessage{
		TenantID:       tenantID,
		UserID:         userID,
		ConversationID: conv.ID,
		Role:           "user",
		Content:        message,
	}
	if err := s.msgRepo.Create(ctx, userMsg); err != nil {
		return nil, fmt.Errorf("保存用户消息失败: %w", err)
	}

	// 获取AI配置
	config, err := s.configRepo.GetByTenant(ctx, tenantID)
	if err != nil {
		return &ChatResponse{
			SessionID:   sessionID,
			MessageID:   userMsg.ID,
			Role:        "assistant",
			Content:     "AI服务未配置或不可用，请联系管理员配置AI模型。",
			ResultType:  "error",
			NeedsConfirm: false,
		}, nil
	}

	// 构建系统提示词
	systemPrompt := s.BuildSystemPrompt()

	// 构建对话历史
	chatHistory := s.buildChatHistory(history)

	// 调用AI模型
	aiResponse, err := s.CallModel(ctx, config, systemPrompt, message, chatHistory)
	if err != nil {
		return &ChatResponse{
			SessionID:   sessionID,
			MessageID:   userMsg.ID,
			Role:        "assistant",
			Content:     fmt.Sprintf("AI服务调用失败: %v", err),
			ResultType:  "error",
			NeedsConfirm: false,
		}, nil
	}

	// 解析AI响应
	intent, err := s.ParseAIResponse(aiResponse.Content)
	if err != nil {
		// 解析失败，返回原始文本
		intent = &AIIntent{
			OperationType: "unknown",
			Module:        "",
			Action:        "",
			Parameters:    nil,
			NaturalDesc:   aiResponse.Content,
			Confidence:    0.0,
		}
	}

	// 构建AI消息
	var operationType string
	var resultType string
	var needsConfirm bool

	if intent.OperationType == "query" {
		operationType = "query"
		resultType = "text"
		needsConfirm = false
	} else if intent.OperationType == "write" {
		operationType = "write"
		resultType = "confirmation"
		needsConfirm = true
	} else if intent.OperationType == "analysis" {
		operationType = "analysis"
		resultType = "text"
		needsConfirm = false
	} else if intent.OperationType == "chat" {
		// 对话模式：直接返回AI回复，不执行任何操作
		operationType = "chat"
		resultType = "text"
		needsConfirm = false
		// 提取纯文本回复（去掉JSON部分）
		plainContent := s.extractPlainReply(aiResponse.Content)
		aiResponse.Content = plainContent
	} else {
		operationType = "unknown"
		resultType = "text"
		needsConfirm = false
	}

	assistantMsg := &model.ChatMessage{
		TenantID:       tenantID,
		UserID:         userID,
		ConversationID: conv.ID,
		Role:           "assistant",
		Content:        aiResponse.Content,
		IntentJSON:     stringPtr(s.marshalIntent(intent)),
		OperationType:  &operationType,
		Status:         "pending",
	}
	if err := s.msgRepo.Create(ctx, assistantMsg); err != nil {
		return nil, fmt.Errorf("保存AI消息失败: %w", err)
	}

	return &ChatResponse{
		SessionID:      sessionID,
		MessageID:     assistantMsg.ID,
		Role:          "assistant",
		Content:       aiResponse.Content,
		ResultType:    resultType,
		ResultData:    intent.Parameters,
		OperationType: operationType,
		Operation:     intent,
		NeedsConfirm:  needsConfirm,
		Metadata: &ResponseMeta{
			TokensUsed: aiResponse.TokensUsed,
			Model:      config.ModelName,
			Confidence: intent.Confidence,
		},
	}, nil
}

// CallModel 调用AI模型
func (s *AIService) CallModel(ctx context.Context, config *model.AIConfig, systemPrompt, userMessage string, history []ChatMessage) (*ModelResponse, error) {
	// 构建请求
	messages := make([]ChatMessage, 0, len(history)+2)
	messages = append(messages, ChatMessage{
		Role:    "system",
		Content: systemPrompt,
	})
	messages = append(messages, history...)
	messages = append(messages, ChatMessage{
		Role:    "user",
		Content: userMessage,
	})

	var reqBody []byte
	var err error
	var endpoint string

	timeout := time.Duration(config.Timeout) * time.Second
	if timeout == 0 {
		timeout = 60 * time.Second
	}

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			Proxy: nil, // 禁用代理，直连AI服务
		},
	}

	switch config.Provider {
	case "openai", "ollama", "custom":
		// OpenAI / Ollama compatible API
		reqBody, err = json.Marshal(ChatCompletionsRequest{
			Model:       config.ModelName,
			Messages:    messages,
			Temperature: config.Temperature,
			MaxTokens:   config.MaxTokens,
		})
		endpoint = config.Endpoint
		if endpoint == "" {
			if config.Provider == "openai" {
				endpoint = "https://api.openai.com/v1/chat/completions"
			}
		}
		// Ensure endpoint has correct path for OpenAI-compatible APIs
		if !strings.Contains(endpoint, "/chat/completions") && !strings.Contains(endpoint, "/messages") {
			endpoint = strings.TrimSuffix(endpoint, "/") + "/chat/completions"
		}

	case "azure":
		// Azure OpenAI
		reqBody, err = json.Marshal(ChatCompletionsRequest{
			Model:       config.ModelName,
			Messages:    messages,
			Temperature: config.Temperature,
			MaxTokens:   config.MaxTokens,
		})
		// Azure endpoint format: https://{resource}.openai.azure.com/openai/deployments/{deployment}/chat/completions?api-version={version}
		endpoint = config.Endpoint
		if endpoint == "" {
			return nil, fmt.Errorf("Azure OpenAI endpoint未配置")
		}

	case "minimax":
		// MiniMax uses Anthropic API format
		reqBody, err = json.Marshal(struct {
			Model       string         `json:"model"`
			Messages    []ChatMessage `json:"messages"`
			MaxTokens   int           `json:"max_tokens"`
			Temperature float64       `json:"temperature,omitempty"`
		}{
			Model:       config.ModelName,
			Messages:    messages,
			MaxTokens:   config.MaxTokens,
			Temperature: config.Temperature,
		})
		endpoint = config.Endpoint
		if endpoint == "" {
			return nil, fmt.Errorf("MiniMax endpoint未配置")
		}
		// MiniMax API path
		if !strings.Contains(endpoint, "/messages") {
			endpoint = strings.TrimSuffix(endpoint, "/") + "/v1/messages"
		}

	default:
		return nil, fmt.Errorf("不支持的AI提供商: %s", config.Provider)
	}

	if err != nil {
		return nil, fmt.Errorf("构建请求失败: %w", err)
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	switch config.Provider {
	case "openai":
		req.Header.Set("Authorization", "Bearer "+config.APIKey)
	case "azure":
		req.Header.Set("api-key", config.APIKey)
	case "ollama", "custom":
		if !strings.HasPrefix(endpoint, "http://localhost") {
			req.Header.Set("Authorization", "Bearer "+config.APIKey)
		}
	case "minimax":
		req.Header.Set("Authorization", "Bearer "+config.APIKey)
		req.Header.Set("anthropic-version", "2023-06-01")
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求AI服务失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取AI响应失败: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("AI服务返回错误: status=%d, body=%s", resp.StatusCode, string(body))
	}

	// 解析响应
	var completionResp ChatCompletionsResponse
	var anthropicResp AnthropicResponse

	if config.Provider == "minimax" {
		// MiniMax uses Anthropic API format
		if err := json.Unmarshal(body, &anthropicResp); err != nil {
			return nil, fmt.Errorf("解析AI响应失败: %w", err)
		}
		if len(anthropicResp.Content) == 0 {
			return nil, fmt.Errorf("AI响应为空")
		}
		// Extract text content (skip thinking blocks)
		var textContent string
		for _, block := range anthropicResp.Content {
			if block.Type == "text" {
				textContent = block.Text
				break
			}
		}
		if textContent == "" {
			return nil, fmt.Errorf("AI响应为空")
		}
		return &ModelResponse{
			Content:    textContent,
			TokensUsed: int64(anthropicResp.Usage.OutputTokens),
		}, nil
	} else {
		// OpenAI compatible format
		if err := json.Unmarshal(body, &completionResp); err != nil {
			return nil, fmt.Errorf("解析AI响应失败: %w", err)
		}
		if len(completionResp.Choices) == 0 {
			return nil, fmt.Errorf("AI响应为空")
		}
		return &ModelResponse{
			Content:    completionResp.Choices[0].Message.Content,
			TokensUsed: int64(completionResp.Usage.TotalTokens),
		}, nil
	}
}

// BuildSystemPrompt 构建系统提示词
func (s *AIService) BuildSystemPrompt() string {
	var sb strings.Builder

	sb.WriteString("你是MOM3.0智能助手，一个专业的制造业执行系统AI助手。\n\n")
	sb.WriteString("你的职责是：\n")
	sb.WriteString("1. 理解用户的自然语言请求\n")
	sb.WriteString("2. 将请求解析为结构化的操作意图\n")
	sb.WriteString("3. 执行只读查询或建议写操作（需要用户确认）\n\n")

	sb.WriteString("## 操作分类\n")
	sb.WriteString("- query: 查询操作（GET请求），只读，不会修改数据\n")
	sb.WriteString("- write: 写操作（POST/PUT/DELETE请求），会修改数据，需要用户确认\n")
	sb.WriteString("- analysis: 分析操作，计算统计等\n")
	sb.WriteString("- navigation: 导航操作，切换页面等\n")
	sb.WriteString("- chat: 对话模式，用户只是闲聊、问问题、不需要执行系统操作\n")
	sb.WriteString("- unknown: 未知操作，无法解析\n\n")

	sb.WriteString("## 可用模块\n")
	sb.WriteString(s.schemaGen.Generate())

	sb.WriteString("\n## 判断规则\n")
	sb.WriteString("- 如果用户只是在聊天、问问题、问候、或表达情绪，但不需要执行任何MOM系统操作 → chat\n")
	sb.WriteString("- 如果用户要求查询、创建、修改、删除系统数据 → query/write/analysis\n")
	sb.WriteString("- 如果用户说\"帮我创建\"、\"查询\"、\"看看\"等明确操作意图 → query/write\n\n")

	sb.WriteString("\n## 输出格式\n")
	sb.WriteString("请以JSON格式回复，格式如下：\n")
	sb.WriteString(`{
  "operation_type": "query|write|analysis|navigation|chat|unknown",
  "module": "模块名(chat时可不填)",
  "action": "操作名(chat时可不填)",
  "parameters": {"键": "值"},
  "natural_desc": "人类可读的描述",
  "confidence": 0.95
}`)

	return sb.String()
}

// ParseAIResponse 解析AI响应为结构化意图
func (s *AIService) ParseAIResponse(content string) (*AIIntent, error) {
	// 尝试提取JSON
	jsonStr := s.extractJSON(content)
	if jsonStr == "" {
		return nil, fmt.Errorf("未找到有效的JSON")
	}

	var intent AIIntent
	if err := json.Unmarshal([]byte(jsonStr), &intent); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	// 验证必填字段
	if intent.OperationType == "" {
		intent.OperationType = "unknown"
	}
	if intent.Module == "" && intent.OperationType != "unknown" {
		return nil, fmt.Errorf("缺少module字段")
	}

	return &intent, nil
}

// extractJSON 从文本中提取JSON
func (s *AIService) extractJSON(text string) string {
	// 尝试找到JSON块
	start := strings.Index(text, "{")
	if start == -1 {
		return ""
	}

	// 找到最外层的闭合括号
	depth := 0
	for i := start; i < len(text); i++ {
		if text[i] == '{' {
			depth++
		} else if text[i] == '}' {
			depth--
			if depth == 0 {
				return text[start : i+1]
			}
		}
	}

	return ""
}

// extractPlainReply 从AI响应中提取纯文本（用于chat模式）
func (s *AIService) extractPlainReply(content string) string {
	// 去掉 <think>...</think> 标签
	content = strings.ReplaceAll(content, "<think>", "")
	content = strings.ReplaceAll(content, "</think>", "")

	// 找到 ``` 或 ```json 开始标记
	codeStart := strings.Index(content, "```")
	if codeStart > 0 {
		// 提取 ``` 之前的纯文本部分
		plainText := strings.TrimSpace(content[:codeStart])
		// 去掉 ```json 或 ``` 标记本身
		plainText = strings.TrimRight(plainText, "\n")
		if plainText != "" {
			return plainText
		}
	}

	// 如果没有代码块，直接返回清理后的内容
	// 尝试找到可能的JSON部分并去掉
	jsonStart := strings.Index(content, "{")
	if jsonStart > 0 {
		plainText := strings.TrimSpace(content[:jsonStart])
		plainText = strings.TrimRight(plainText, "\n")
		if plainText != "" {
			return plainText
		}
	}

	return strings.TrimSpace(content)
}

// buildChatHistory 构建对话历史
func (s *AIService) buildChatHistory(history []*model.ChatMessage) []ChatMessage {
	result := make([]ChatMessage, 0, len(history))
	for _, msg := range history {
		role := msg.Role
		if role == "system" {
			continue // 不发送系统消息
		}
		result = append(result, ChatMessage{
			Role:    role,
			Content: msg.Content,
		})
	}
	return result
}

// marshalIntent 将意图序列化为JSON字符串
func (s *AIService) marshalIntent(intent *AIIntent) string {
	data, _ := json.Marshal(intent)
	return string(data)
}

// GetConversations 获取用户的会话列表
func (s *AIService) GetConversations(ctx context.Context, tenantID, userID int64) ([]model.ChatConversation, error) {
	return s.convRepo.ListByUser(ctx, tenantID, userID, 50)
}

// GetConversationMessages 获取会话的消息列表
func (s *AIService) GetConversationMessages(ctx context.Context, tenantID int64, sessionID string, limit int) ([]model.ChatMessage, error) {
	conv, err := s.convRepo.GetBySessionID(ctx, tenantID, sessionID)
	if err != nil {
		return nil, err
	}
	if limit <= 0 {
		limit = 50
	}
	return s.msgRepo.ListByConversation(ctx, conv.ID, limit)
}

// DeleteConversation 删除会话
func (s *AIService) DeleteConversation(ctx context.Context, tenantID int64, sessionID string) error {
	conv, err := s.convRepo.GetBySessionID(ctx, tenantID, sessionID)
	if err != nil {
		return err
	}
	return s.convRepo.Delete(ctx, conv.ID)
}

// TestConnection 测试AI连接
func (s *AIService) TestConnection(ctx context.Context, config *model.AIConfig) error {
	testPrompt := "Hello, please respond with 'OK' if you receive this message."

	_, err := s.CallModel(ctx, config, testPrompt, "Hello", nil)
	return err
}

// ========== 内部类型 ==========

// ChatMessage 聊天消息结构
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatMessageEntry 对话历史条目
type ChatMessageEntry struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionsRequest OpenAI兼容的聊天补全请求
type ChatCompletionsRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Temperature float64       `json:"temperature,omitempty"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
}

// ChatCompletionsResponse 聊天补全响应
type ChatCompletionsResponse struct {
	ID      string   `json:"id"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice 选择
type Choice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

// Usage 使用量
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ModelResponse 模型响应
type ModelResponse struct {
	Content    string
	TokensUsed int64
}

// AnthropicResponse Anthropic API格式响应
type AnthropicResponse struct {
	ID      string             `json:"id"`
	Type    string             `json:"type"`
	Role    string             `json:"role"`
	Content []AnthropicContent  `json:"content"`
	Usage   AnthropicUsage     `json:"usage"`
}

// AnthropicContent 内容块
type AnthropicContent struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

// AnthropicUsage 使用量
type AnthropicUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

// stringPtr 返回字符串指针
func stringPtr(s string) *string {
	return &s
}
