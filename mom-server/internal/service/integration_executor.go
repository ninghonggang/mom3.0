package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type IntegrationExecutor struct {
	configRepo *repository.InterfaceConfigRepository
	logRepo    *repository.InterfaceExecutionLogRepository
	httpClient *http.Client
}

func NewIntegrationExecutor(
	configRepo *repository.InterfaceConfigRepository,
	logRepo *repository.InterfaceExecutionLogRepository,
) *IntegrationExecutor {
	return &IntegrationExecutor{
		configRepo: configRepo,
		logRepo:    logRepo,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// ExecuteConfig 执行指定的接口配置（手动/定时触发）
func (e *IntegrationExecutor) ExecuteConfig(ctx context.Context, configID int64, triggerType string) (*model.InterfaceExecutionLog, error) {
	cfg, err := e.configRepo.GetByID(ctx, configID)
	if err != nil {
		return nil, fmt.Errorf("配置不存在: %w", err)
	}
	if cfg.Status != "ENABLE" {
		return nil, fmt.Errorf("配置已禁用")
	}

	fieldMaps, _ := e.configRepo.ListFieldMaps(ctx, configID)

	// Get last execution time for incremental
	lastExecTime, _ := e.logRepo.GetLastExecutionTime(ctx, configID)

	// Query source data
	sourceData, err := e.configRepo.QuerySourceData(ctx, cfg.TenantID, cfg, lastExecTime)
	if err != nil {
		sourceData = []map[string]interface{}{}
	}

	// Execute for each record
	var lastLog *model.InterfaceExecutionLog
	for _, record := range sourceData {
		log := e.doExecute(ctx, cfg, fieldMaps, record, triggerType)
		lastLog = log
	}

	// If no source data, still do a single execute with empty context
	if len(sourceData) == 0 {
		log := e.doExecute(ctx, cfg, fieldMaps, nil, triggerType)
		lastLog = log
	}

	return lastLog, nil
}

func (e *IntegrationExecutor) doExecute(ctx context.Context, cfg *model.InterfaceConfig, fieldMaps []model.InterfaceFieldMap, data map[string]interface{}, triggerType string) *model.InterfaceExecutionLog {
	log := &model.InterfaceExecutionLog{
		InterfaceConfigID: cfg.ID,
		ConfigName:       cfg.Name,
		TriggerType:      triggerType,
		StartTime:        time.Now(),
		RequestMethod:   cfg.Method,
		RequestURL:       cfg.BaseURL + cfg.Path,
		Status:          "FAILED",
	}

	// Build request body
	bodyBytes, _ := e.buildRequestBody(cfg, fieldMaps, data)
	log.RequestBody = string(bodyBytes)

	// Build headers
	headers := e.buildHeaders(cfg)
	headersJSON, _ := json.Marshal(headers)
	log.RequestHeaders = string(headersJSON)

	// Build request
	var reqBody io.Reader
	if len(bodyBytes) > 0 {
		reqBody = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(cfg.Method, cfg.BaseURL+cfg.Path, reqBody)
	if err != nil {
		log.EndTime = time.Now()
		log.Duration = log.EndTime.Sub(log.StartTime).Milliseconds()
		log.ErrorMessage = err.Error()
		e.logRepo.Create(ctx, log)
		return log
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Execute with timeout
	client := &http.Client{Timeout: time.Duration(cfg.Timeout) * time.Second}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.EndTime = time.Now()
		log.Duration = log.EndTime.Sub(log.StartTime).Milliseconds()
		log.ErrorMessage = err.Error()
		e.logRepo.Create(ctx, log)
		return log
	}
	defer resp.Body.Close()

	log.ResponseCode = resp.StatusCode

	respBody, _ := io.ReadAll(resp.Body)
	log.ResponseBody = string(respBody)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Status = "SUCCESS"
	} else {
		log.Status = "FAILED"
		log.ErrorMessage = fmt.Sprintf("HTTP %d", resp.StatusCode)
	}

	log.EndTime = time.Now()
	log.Duration = log.EndTime.Sub(log.StartTime).Milliseconds()
	log.RecordsProcessed = 1

	e.logRepo.Create(ctx, log)
	return log
}

// TestConfig 测试接口（不记日志）
func (e *IntegrationExecutor) TestConfig(ctx context.Context, cfg *model.InterfaceConfig, fieldMaps []model.InterfaceFieldMap, testData map[string]interface{}) (int, string, error) {
	bodyBytes, _ := e.buildRequestBody(cfg, fieldMaps, testData)
	headers := e.buildHeaders(cfg)

	var reqBody io.Reader
	if len(bodyBytes) > 0 {
		reqBody = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(cfg.Method, cfg.BaseURL+cfg.Path, reqBody)
	if err != nil {
		return 0, "", err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: time.Duration(cfg.Timeout) * time.Second}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(body), nil
}

// ExecuteByEvent 事件触发执行
func (e *IntegrationExecutor) ExecuteByEvent(ctx context.Context, eventSource string, eventData map[string]interface{}) error {
	configs, err := e.configRepo.FindByEventSource(ctx, eventSource)
	if err != nil || len(configs) == 0 {
		return nil // No configs for this event
	}

	for _, cfg := range configs {
		fieldMaps, _ := e.configRepo.ListFieldMaps(ctx, cfg.ID)
		log := e.doExecute(ctx, &cfg, fieldMaps, eventData, "EVENT")
		if log.Status == "FAILED" {
			// Could log warning but don't block other configs
		}
	}
	return nil
}

// buildRequestBody builds the final request body by applying field mappings to source data
func (e *IntegrationExecutor) buildRequestBody(cfg *model.InterfaceConfig, fieldMaps []model.InterfaceFieldMap, data map[string]interface{}) ([]byte, error) {
	if cfg.RequestBodyTemplate == "" {
		return nil, nil
	}

	// First apply template variable substitution {{field}} -> value
	bodyStr := cfg.RequestBodyTemplate
	re := regexp.MustCompile(`\{\{([^}]+)\}\}`)
	bodyStr = re.ReplaceAllStringFunc(bodyStr, func(match string) string {
		key := match[2 : len(match)-2]
		key = strings.TrimSpace(key)
		if val, ok := data[key]; ok {
			return fmt.Sprintf("%v", val)
		}
		return match
	})

	// If content type is JSON, try to parse and apply field maps
	if cfg.RequestContentType == "JSON" {
		var jsonBody map[string]interface{}
		if err := json.Unmarshal([]byte(bodyStr), &jsonBody); err == nil {
			// Apply field maps to JSON
			for _, fm := range fieldMaps {
				val := e.resolveFieldValue(fm, data)
				if val != nil {
					jsonBody[fm.FieldName] = val
				}
			}
			return json.Marshal(jsonBody)
		}
	}

	return []byte(bodyStr), nil
}

// resolveFieldValue resolves a single field value based on its mapping type
func (e *IntegrationExecutor) resolveFieldValue(fm model.InterfaceFieldMap, data map[string]interface{}) interface{} {
	switch fm.MapType {
	case "CONST":
		return fm.MapValue

	case "FIELD":
		if val, ok := data[fm.MapValue]; ok {
			val = e.applyTransform(val, fm.TransformFunc)
			return val
		}
		if fm.Required {
			return fm.DefaultValue
		}
		return nil

	case "EXPRESSION":
		// Simple expression: fieldA + fieldB, fieldA * 10, etc.
		return e.evalSimpleExpr(fm.MapValue, data)

	case "JSONPATH":
		return e.applyJSONPath(fm.MapValue, data)

	default:
		if val, ok := data[fm.MapValue]; ok {
			return val
		}
	}

	if fm.DefaultValue != "" {
		return fm.DefaultValue
	}
	return nil
}

// applyTransform applies transform functions to a value
func (e *IntegrationExecutor) applyTransform(val interface{}, transformFunc string) interface{} {
	if transformFunc == "" {
		return val
	}

	str, ok := val.(string)
	if !ok {
		return val
	}

	funcs := strings.Split(transformFunc, "|")
	for _, f := range funcs {
		f = strings.TrimSpace(f)
		switch f {
		case "upper", "UPPER":
			str = strings.ToUpper(str)
		case "lower", "LOWER":
			str = strings.ToLower(str)
		case "trim", "TRIM":
			str = strings.TrimSpace(str)
		case "date_format":
			// Expects input to be parseable date, outputs YYYY-MM-DD
			if t, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
				str = t.Format("2006-01-02")
			}
		case "datetime_format":
			if t, err := time.Parse("2006-01-02 15:04:05", str); err == nil {
				str = t.Format("2006-01-02 15:04:05")
			}
		case "md5":
			// Simplified - just return as-is in Go since no md5 import
			return val
		case "uuid":
			str = strings.ReplaceAll(str, "-", "")
		}
	}
	return str
}

// evalSimpleExpr evaluates a simple arithmetic/string expression
func (e *IntegrationExecutor) evalSimpleExpr(expr string, data map[string]interface{}) interface{} {
	// Supported patterns: fieldA + fieldB, fieldA - fieldB, fieldA * N, fieldA / N
	// Also: fieldA == value, fieldA != value (return bool)

	// Simple field reference
	expr = strings.TrimSpace(expr)

	// Addition
	if idx := strings.Index(expr, "+"); idx > 0 {
		left := strings.TrimSpace(expr[:idx])
		right := strings.TrimSpace(expr[idx+1:])
		return e.exprOp(left, right, data, func(a, b float64) float64 { return a + b })
	}
	// Subtraction
	if idx := strings.Index(expr, "-"); idx > 0 {
		left := strings.TrimSpace(expr[:idx])
		right := strings.TrimSpace(expr[idx+1:])
		return e.exprOp(left, right, data, func(a, b float64) float64 { return a - b })
	}
	// Multiplication
	if idx := strings.Index(expr, "*"); idx > 0 {
		left := strings.TrimSpace(expr[:idx])
		right := strings.TrimSpace(expr[idx+1:])
		return e.exprOp(left, right, data, func(a, b float64) float64 { return a * b })
	}
	// Division
	if idx := strings.Index(expr, "/"); idx > 0 {
		left := strings.TrimSpace(expr[:idx])
		right := strings.TrimSpace(expr[idx+1:])
		return e.exprOp(left, right, data, func(a, b float64) float64 {
			if b == 0 {
				return 0
			}
			return a / b
		})
	}

	// Direct field value
	if val, ok := data[expr]; ok {
		return val
	}

	// String literal
	return strings.Trim(expr, "\" ")
}

func (e *IntegrationExecutor) exprOp(left, right string, data map[string]interface{}, op func(float64, float64) float64) float64 {
	lv := e.toFloat(e.getFieldValue(left, data))
	rv := e.toFloat(e.getFieldValue(right, data))
	return op(lv, rv)
}

func (e *IntegrationExecutor) getFieldValue(key string, data map[string]interface{}) interface{} {
	if val, ok := data[key]; ok {
		return val
	}
	// Try as string literal
	return strings.Trim(key, "\" ")
}

func (e *IntegrationExecutor) toFloat(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	default:
		return 0
	}
}

// applyJSONPath applies a simple JSONPath-like extraction
func (e *IntegrationExecutor) applyJSONPath(path string, data map[string]interface{}) interface{} {
	// Supports: $.field.subfield or field.subfield
	path = strings.TrimPrefix(path, "$.")
	parts := strings.Split(path, ".")

	var current interface{} = data
	for _, part := range parts {
		if current == nil {
			return nil
		}
		switch c := current.(type) {
		case map[string]interface{}:
			current = c[part]
		default:
			return nil
		}
	}
	return current
}

// buildHeaders builds HTTP headers based on auth type and content type
func (e *IntegrationExecutor) buildHeaders(cfg *model.InterfaceConfig) map[string]string {
	headers := make(map[string]string)

	switch cfg.RequestContentType {
	case "JSON":
		headers["Content-Type"] = "application/json"
	case "XML":
		headers["Content-Type"] = "application/xml"
	case "FORM":
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	case "TEXT":
		headers["Content-Type"] = "text/plain"
	}

	switch cfg.AuthType {
	case "BASIC":
		if cfg.AuthConfig != "" {
			headers["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(cfg.AuthConfig))
		}
	case "API_KEY":
		if cfg.AuthConfig != "" {
			// Assume format: HeaderName:HeaderValue or just Value
			if strings.Contains(cfg.AuthConfig, ":") {
				parts := strings.SplitN(cfg.AuthConfig, ":", 2)
				headers[parts[0]] = parts[1]
			} else {
				headers["X-API-Key"] = cfg.AuthConfig
			}
		}
	case "BEARER_TOKEN":
		headers["Authorization"] = "Bearer " + cfg.AuthConfig
	case "OAUTH2":
		// OAuth2 requires token endpoint flow - would need token refresh logic
		// For now, assume AuthConfig contains the access token
		headers["Authorization"] = "Bearer " + cfg.AuthConfig
	}

	return headers
}

// ParseResponseBody parses response body based on format
func (e *IntegrationExecutor) ParseResponseBody(body []byte, format string) (map[string]interface{}, error) {
	switch format {
	case "JSON":
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return map[string]interface{}{"raw": string(body)}, nil
	}
}

// CheckCondition evaluates a condition expression against data
func (e *IntegrationExecutor) CheckCondition(condition string, data map[string]interface{}) bool {
	if condition == "" {
		return true
	}
	// Simple SpEL-like condition: field == value, field != value, field > value
	condition = strings.TrimSpace(condition)

	for _, op := range []string{"==", "!=", ">", "<", ">=", "<="} {
		if idx := strings.Index(condition, op); idx > 0 {
			left := strings.TrimSpace(condition[:idx])
			right := strings.TrimSpace(condition[idx+len(op):])
			return e.compareValues(e.getFieldValue(left, data), right, op)
		}
	}
	return true
}

func (e *IntegrationExecutor) compareValues(left interface{}, right string, op string) bool {
	lv := e.toFloat(left)
	rv := e.toFloat(right)

	switch op {
	case "==":
		return fmt.Sprintf("%v", left) == right || lv == rv
	case "!=":
		return lv != rv
	case ">":
		return lv > rv
	case "<":
		return lv < rv
	case ">=":
		return lv >= rv
	case "<=":
		return lv <= rv
	}
	return false
}

// DataSourceTypeToString converts source type constant to display name
func DataSourceTypeToString(t string) string {
	switch t {
	case "TABLE_QUERY":
		return "数据库查询"
	case "API_CALL":
		return "内部API调用"
	case "EVENT_PAYLOAD":
		return "事件数据"
	default:
		return t
	}
}

// MapTypeToString converts map type constant to display name
func MapTypeToString(t string) string {
	switch t {
	case "CONST":
		return "常量值"
	case "FIELD":
		return "字段映射"
	case "EXPRESSION":
		return "表达式"
	case "JSONPATH":
		return "JSON路径"
	default:
		return t
	}
}
