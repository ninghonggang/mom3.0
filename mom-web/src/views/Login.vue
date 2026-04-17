<template>
  <div class="login-page">
    <!-- Left Panel - Decorative -->
    <div class="login-panel-left">
      <div class="panel-content">
        <div class="brand">
          <img src="/favicon.svg" alt="logo" class="brand-logo">
          <h1 class="brand-name">MOM3.0</h1>
        </div>
        <p class="brand-tagline">制造运营管理系统</p>

        <div class="feature-list">
          <div class="feature-item">
            <div class="feature-icon">
              <el-icon><Box /></el-icon>
            </div>
            <div class="feature-text">
              <h3>生产执行</h3>
              <p>实时掌控生产进度</p>
            </div>
          </div>
          <div class="feature-item">
            <div class="feature-icon">
              <el-icon><Monitor /></el-icon>
            </div>
            <div class="feature-text">
              <h3>设备管理</h3>
              <p>OEE设备效能分析</p>
            </div>
          </div>
          <div class="feature-item">
            <div class="feature-icon">
              <el-icon><CircleCheck /></el-icon>
            </div>
            <div class="feature-text">
              <h3>质量管理</h3>
              <p>SPC全流程质量管控</p>
            </div>
          </div>
        </div>

        <div class="panel-footer">
          <span>闻荫科技 © 2026</span>
        </div>
      </div>
    </div>

    <!-- Right Panel - Login Form -->
    <div class="login-panel-right">
      <div class="form-container">
        <div class="form-header">
          <h2>欢迎登录</h2>
          <p class="form-subtitle">请输入您的账号信息</p>
        </div>

        <el-form
          ref="formRef"
          :model="loginForm"
          :rules="rules"
          class="login-form"
          @submit.prevent="handleLogin"
        >
          <el-form-item prop="username">
            <el-input
              v-model="loginForm.username"
              placeholder="请输入用户名"
              size="large"
              :prefix-icon="User"
            />
          </el-form-item>

          <el-form-item prop="password">
            <el-input
              v-model="loginForm.password"
              type="password"
              placeholder="请输入密码"
              size="large"
              :prefix-icon="Lock"
              show-password
              @keyup.enter="handleLogin"
            />
          </el-form-item>

          <el-form-item>
            <div class="form-options">
              <el-checkbox v-model="rememberMe">记住密码</el-checkbox>
              <span class="forgot-link">忘记密码？</span>
            </div>
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              size="large"
              :loading="loading"
              class="login-btn"
              @click="handleLogin"
            >
              登 录
            </el-button>
          </el-form-item>
        </el-form>

        <div class="login-tip">
          <el-alert type="info" :closable="false" show-icon>
            <template #title>
              <span>演示账号：admin / admin123</span>
            </template>
          </el-alert>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage, FormInstance, FormRules } from 'element-plus'
import { User, Lock, Box, Monitor, CircleCheck } from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const formRef = ref<FormInstance>()
const loading = ref(false)
const rememberMe = ref(false)

const loginForm = reactive({
  username: 'admin',
  password: 'admin123'
})

const rules: FormRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

const handleLogin = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (!valid) return

    loading.value = true
    try {
      await authStore.loginAction(loginForm.username, loginForm.password)
      ElMessage.success('登录成功')

      const redirect = route.query.redirect as string
      router.push(redirect || '/')
    } catch (error: any) {
      ElMessage.error(error.message || '登录失败')
    } finally {
      loading.value = false
    }
  })
}
</script>

<style scoped lang="scss">
.login-page {
  display: flex;
  min-height: 100vh;
}

// Left Panel
.login-panel-left {
  flex: 1;
  background: linear-gradient(135deg, #1e293b 0%, #0f172a 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: -50%;
    left: -50%;
    width: 200%;
    height: 200%;
    background: radial-gradient(circle, rgba(37, 99, 235, 0.15) 0%, transparent 50%);
    animation: pulse 8s ease-in-out infinite;
  }

  &::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    height: 40%;
    background: linear-gradient(to top, rgba(37, 99, 235, 0.1), transparent);
  }
}

@keyframes pulse {
  0%, 100% { transform: scale(1); opacity: 1; }
  50% { transform: scale(1.1); opacity: 0.8; }
}

.panel-content {
  position: relative;
  z-index: 1;
  text-align: center;
  padding: 48px;
  max-width: 480px;
}

.brand {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
  margin-bottom: 16px;

  .brand-logo {
    width: 56px;
    height: 56px;
  }

  .brand-name {
    font-size: 42px;
    font-weight: 700;
    color: #fff;
    letter-spacing: 4px;
    margin: 0;
  }
}

.brand-tagline {
  font-size: 18px;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 64px;
  letter-spacing: 8px;
}

.feature-list {
  display: flex;
  flex-direction: column;
  gap: 24px;
  text-align: left;
}

.feature-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px 20px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;

  &:hover {
    background: rgba(255, 255, 255, 0.1);
    transform: translateX(8px);
  }

  .feature-icon {
    width: 44px;
    height: 44px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: rgba(37, 99, 235, 0.3);
    border-radius: 10px;
    color: #60a5fa;
    font-size: 20px;
  }

  .feature-text {
    h3 {
      font-size: 16px;
      font-weight: 600;
      color: #fff;
      margin: 0 0 4px 0;
    }

    p {
      font-size: 13px;
      color: rgba(255, 255, 255, 0.5);
      margin: 0;
    }
  }
}

.panel-footer {
  position: absolute;
  bottom: 32px;
  left: 0;
  right: 0;
  text-align: center;
  color: rgba(255, 255, 255, 0.4);
  font-size: 13px;
}

// Right Panel
.login-panel-right {
  width: 520px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f1f5f9;
  padding: 48px;
}

.form-container {
  width: 100%;
  max-width: 360px;
}

.form-header {
  margin-bottom: 40px;

  h2 {
    font-size: 28px;
    font-weight: 600;
    color: #1f2937;
    margin: 0 0 8px 0;
  }

  .form-subtitle {
    font-size: 14px;
    color: #6b7280;
    margin: 0;
  }
}

.login-form {
  .login-btn {
    width: 100%;
    height: 48px;
    font-size: 16px;
    font-weight: 500;
    border-radius: 8px;
    background: #2563eb;
    border-color: #2563eb;

    &:hover {
      background: #1d4ed8;
      border-color: #1d4ed8;
    }
  }
}

.form-options {
  display: flex;
  justify-content: space-between;
  width: 100%;

  .forgot-link {
    font-size: 14px;
    color: #2563eb;
    cursor: pointer;

    &:hover {
      color: #1d4ed8;
    }
  }
}

.login-tip {
  margin-top: 24px;

  :deep(.el-alert) {
    border-radius: 8px;
  }
}

// Responsive
@media (max-width: 1024px) {
  .login-panel-left {
    display: none;
  }

  .login-panel-right {
    width: 100%;
  }
}
</style>
