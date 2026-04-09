import request from '@/utils/request'

export interface AIConfig {
  id?: number
  tenant_id?: number
  config_name?: string
  provider: string
  endpoint: string
  api_version?: string
  model_name: string
  api_key?: string
  temperature: number
  max_tokens: number
  timeout: number
  enable: boolean
}

export interface ChatMessage {
  id?: number
  role: 'user' | 'assistant' | 'system'
  content: string
  intent_json?: string
  operation_type?: string
  status?: string
  tool_result?: string
  created_at?: string
}

export interface ChatConversation {
  id: number
  tenant_id: number
  user_id: number
  session_id: string
  title?: string
  created_at: string
}

export interface SendMessageRequest {
  session_id?: string
  message: string
}

export interface SendMessageResponse {
  session_id: string
  message_id: number
  role: string
  content: string
  result_type: 'table' | 'chart' | 'text' | 'form' | 'confirmation' | 'error'
  result_data?: unknown
  operation_type?: string
  operation?: unknown
  needs_confirmation: boolean
  metadata?: {
    tokens_used?: number
    model?: string
    confidence?: number
  }
}

export interface ExecuteRequest {
  message_id: number
  confirmed: boolean
}

// AI Config API
export function getAiConfig() {
  return request.get('/ai/config')
}

export function updateAiConfig(data: AIConfig) {
  return request.put('/ai/config', data)
}

export function testAiConfig(data: AIConfig) {
  return request.post('/ai/config/test', data)
}

export function getSchema() {
  return request.get('/ai/schema')
}

// Chat API
export function getConversations() {
  return request.get('/ai/chat/conversations')
}

export function getConversationHistory(sessionId: string) {
  return request.get(`/ai/chat/conversations/${sessionId}`)
}

export function deleteConversation(sessionId: string) {
  return request.delete(`/ai/chat/conversations/${sessionId}`)
}

export function sendMessage(data: SendMessageRequest) {
  return request.post('/ai/chat/send', data)
}

export function executeOperation(data: ExecuteRequest) {
  return request.post('/ai/chat/execute', data)
}
