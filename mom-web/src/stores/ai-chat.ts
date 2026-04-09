import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getConversations, getConversationHistory, sendMessage, executeOperation, type SendMessageResponse } from '@/api/ai-chat'

export interface Conversation {
  id: number
  tenant_id: number
  user_id: number
  session_id: string
  title?: string
  created_at: string
}

export interface Message {
  role: 'user' | 'assistant' | 'system'
  content: string
  result_type?: 'table' | 'chart' | 'text' | 'form' | 'confirmation' | 'error'
  result_data?: unknown
  operation_type?: string
  operation?: unknown
  needs_confirmation?: boolean
  message_id?: number
  status?: string
  created_at?: string
}

export const useAIChatStore = defineStore('ai-chat', () => {
  const conversations = ref<Conversation[]>([])
  const currentSessionId = ref<string>('')
  const messages = ref<Message[]>([])
  const isOpen = ref(false)
  const isLoading = ref(false)
  const pendingConfirmations = ref<{ message_id: number; operation: unknown }[]>([])

  const hasUnread = computed(() => pendingConfirmations.value.length > 0)

  async function loadConversations() {
    try {
      const res = await getConversations()
      if (res.data.code === 200) {
        conversations.value = res.data.data.list || []
      }
    } catch (error) {
      console.error('加载会话列表失败', error)
    }
  }

  async function loadHistory(sessionId: string) {
    try {
      const res = await getConversationHistory(sessionId)
      if (res.data.code === 200) {
        messages.value = res.data.data.list || []
        currentSessionId.value = sessionId
      }
    } catch (error) {
      console.error('加载会话历史失败', error)
    }
  }

  async function sendChatMessage(content: string) {
    isLoading.value = true
    try {
      const res = await sendMessage({
        session_id: currentSessionId.value || undefined,
        message: content
      })

      const data: SendMessageResponse = res.data
      if (data.session_id && !currentSessionId.value) {
        currentSessionId.value = data.session_id
        await loadConversations()
      }

      messages.value.push({
        role: 'user',
        content: content,
        created_at: new Date().toISOString()
      })

      messages.value.push({
        role: 'assistant',
        content: data.content,
        result_type: data.result_type,
        result_data: data.result_data,
        operation_type: data.operation_type,
        operation: data.operation,
        needs_confirmation: data.needs_confirmation,
        message_id: data.message_id,
        created_at: new Date().toISOString()
      })

      if (data.needs_confirmation) {
        pendingConfirmations.value.push({
          message_id: data.message_id,
          operation: data.operation
        })
      }

      return data
    } finally {
      isLoading.value = false
    }
  }

  async function confirmOperation(messageId: number, confirmed: boolean) {
    try {
      const res = await executeOperation({ message_id: messageId, confirmed })
      const msg = messages.value.find(m => m.message_id === messageId)
      if (msg) {
        msg.status = confirmed ? 'executed' : 'rejected'
        msg.result_data = res.data.data
      }
      pendingConfirmations.value = pendingConfirmations.value.filter(p => p.message_id !== messageId)
      return res.data
    } catch (error) {
      console.error('执行操作失败', error)
      throw error
    }
  }

  function toggleChat() {
    isOpen.value = !isOpen.value
  }

  function clearHistory() {
    messages.value = []
    currentSessionId.value = ''
  }

  return {
    conversations,
    currentSessionId,
    messages,
    isOpen,
    isLoading,
    pendingConfirmations,
    hasUnread,
    loadConversations,
    loadHistory,
    sendChatMessage,
    confirmOperation,
    toggleChat,
    clearHistory
  }
})
