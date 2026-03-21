<template>
  <div class="error-404">
    <div class="content">
      <h1>404</h1>
      <p>抱歉，您访问的页面不存在</p>
      <el-button type="primary" @click="goHome">返回首页</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const goHome = () => {
  router.push('/')
}
</script>

<style scoped lang="scss">
.error-404 {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: #f5f7fa;

  .content {
    text-align: center;

    h1 {
      font-size: 120px;
      color: #409eff;
      margin-bottom: 20px;
    }

    p {
      font-size: 18px;
      color: #606266;
      margin-bottom: 30px;
    }
  }
}
</style>
