<template>
  <div class="ai-chat-button" @click="toggleChat">
    <el-badge :is-dot="hasUnread" :hidden="!hasUnread">
      <el-button
        :icon="isOpen ? 'Close' : 'ChatDotRound'"
        type="primary"
        circle
        size="large"
        class="chat-btn"
      />
    </el-badge>
  </div>
</template>

<script setup lang="ts">
import { useAIChatStore } from '@/stores/ai-chat'
import { storeToRefs } from 'pinia'

const store = useAIChatStore()
const { isOpen, hasUnread } = storeToRefs(store)
const { toggleChat } = store
</script>

<style scoped>
.ai-chat-button {
  position: fixed;
  bottom: 80px;
  right: 24px;
  z-index: 9999;
}

.chat-btn {
  width: 56px;
  height: 56px;
  font-size: 24px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}
</style>
