<template>
  <el-card v-show="isOpen" class="ai-chat-window" :body-style="{ padding: '0' }">
    <template #header>
      <div class="chat-header">
        <div class="header-title">
          <el-icon><ChatDotRound /></el-icon>
          <span>AI 助手</span>
        </div>
        <el-select
          v-model="currentSession"
          placeholder="选择会话"
          size="small"
          style="width: 160px; margin-right: 8px"
          @change="handleSessionChange"
          clearable
        >
          <el-option
            v-for="conv in conversations"
            :key="conv.session_id"
            :label="conv.title || '新会话'"
            :value="conv.session_id"
          />
        </el-select>
        <el-button size="small" @click="handleNewChat">新建</el-button>
      </div>
    </template>

    <div class="chat-body" ref="chatBodyRef">
      <div class="message-list">
        <div
          v-for="(msg, index) in messages"
          :key="index"
          :class="['message', msg.role]"
        >
          <div class="message-content">
            <div class="message-bubble">
              <div v-if="!msg.result_type" class="text-content">
                {{ msg.content }}
              </div>

              <AiResultRenderer
                v-else
                :result-type="msg.result_type"
                :data="msg.result_data || msg.content"
                :operation="msg.operation"
                :message-id="msg.message_id"
                @confirm="handleConfirm"
                @cancel="handleCancel"
              />
            </div>
          </div>
        </div>

        <div v-if="isLoading" class="message assistant">
          <div class="message-content">
            <div class="message-bubble loading">
              <el-icon class="is-loading"><Loading /></el-icon>
              <span>AI思考中...</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="chat-input">
      <el-input
        v-model="inputText"
        type="textarea"
        :rows="2"
        placeholder="输入消息... (Enter发送，Shift+Enter换行)"
        @keydown.enter.exact.prevent="handleSend"
      />
      <el-button
        type="primary"
        :disabled="!inputText.trim() || isLoading"
        @click="handleSend"
        style="margin-top: 8px"
      >
        发送
      </el-button>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { ref, watch, nextTick, onMounted } from 'vue'
import { useAIChatStore } from '@/stores/ai-chat'
import { storeToRefs } from 'pinia'
import { ChatDotRound, Loading } from '@element-plus/icons-vue'
import AiResultRenderer from './AiResultRenderer.vue'

const store = useAIChatStore()
const {
  conversations,
  currentSessionId,
  messages,
  isOpen,
  isLoading
} = storeToRefs(store)

const chatBodyRef = ref<HTMLElement>()
const inputText = ref('')
const currentSession = ref('')

onMounted(() => {
  store.loadConversations()
})

watch(currentSessionId, (val) => {
  currentSession.value = val
})

watch(messages, async () => {
  await nextTick()
  if (chatBodyRef.value) {
    chatBodyRef.value.scrollTop = chatBodyRef.value.scrollHeight
  }
}, { deep: true })

async function handleSend() {
  const text = inputText.value.trim()
  if (!text) return

  inputText.value = ''
  await store.sendChatMessage(text)
}

function handleSessionChange(sessionId: string) {
  if (sessionId) {
    store.loadHistory(sessionId)
  }
}

function handleNewChat() {
  store.clearHistory()
  currentSession.value = ''
}

async function handleConfirm(messageId: number, confirmed: boolean) {
  await store.confirmOperation(messageId, confirmed)
}

async function handleCancel(messageId: number, confirmed: boolean) {
  await store.confirmOperation(messageId, false)
}
</script>

<style scoped>
.ai-chat-window {
  position: fixed;
  bottom: 150px;
  right: 24px;
  width: 420px;
  max-height: 600px;
  z-index: 9998;
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.15);
}

.chat-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
}

.chat-body {
  height: 400px;
  overflow-y: auto;
  padding: 16px;
  background: #f5f7fa;
}

.message-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.message {
  display: flex;
  max-width: 85%;
}

.message.user {
  align-self: flex-end;
}

.message.assistant {
  align-self: flex-start;
}

.message-content {
  display: flex;
  flex-direction: column;
}

.message.user .message-content {
  align-items: flex-end;
}

.message-bubble {
  padding: 10px 14px;
  border-radius: 12px;
  font-size: 14px;
  line-height: 1.5;
  word-break: break-word;
}

.message.user .message-bubble {
  background: #409eff;
  color: white;
  border-bottom-right-radius: 4px;
}

.message.assistant .message-bubble {
  background: white;
  color: #303133;
  border-bottom-left-radius: 4px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
}

.message-bubble.loading {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #909399;
}

.text-content {
  white-space: pre-wrap;
}

.chat-input {
  padding: 12px;
  border-top: 1px solid #ebeef5;
}
</style>
