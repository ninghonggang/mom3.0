<template>
  <div class="page-container">
    <div class="page-header">
      <h2 class="page-title">系统设置</h2>
    </div>

    <div class="settings-grid">
      <!-- Theme Settings Card -->
      <el-card class="settings-card">
        <template #header>
          <div class="card-header">
            <el-icon><Brush /></el-icon>
            <span>主题设置</span>
          </div>
        </template>

        <el-form label-position="top">
          <el-form-item label="主题模式">
            <el-radio-group v-model="settingsStore.theme" class="theme-radio-group">
              <el-radio value="light">
                <div class="theme-option">
                  <el-icon><Sunny /></el-icon>
                  <span>浅色</span>
                </div>
              </el-radio>
              <el-radio value="dark">
                <div class="theme-option">
                  <el-icon><Moon /></el-icon>
                  <span>深色</span>
                </div>
              </el-radio>
              <el-radio value="custom">
                <div class="theme-option">
                  <el-icon><Brush /></el-icon>
                  <span>自定义</span>
                </div>
              </el-radio>
            </el-radio-group>
          </el-form-item>

          <el-form-item v-if="settingsStore.theme === 'custom'" label="主题颜色">
            <div class="color-picker-row">
              <el-color-picker v-model="settingsStore.customThemeColor" @change="settingsStore.applyTheme" />
              <span class="color-value">{{ settingsStore.customThemeColor }}</span>
            </div>
          </el-form-item>

          <el-form-item label="字体">
            <el-select v-model="settingsStore.fontFamily" @change="settingsStore.applyFontFamily" style="width: 200px;">
              <el-option value="default" label="系统默认字体" />
              <el-option value="helvetica" label="Helvetica" />
              <el-option value="pingfang" label="苹方字体" />
              <el-option value="microsoft" label="微软雅黑" />
            </el-select>
          </el-form-item>

          <el-form-item label="侧边栏">
            <el-switch
              v-model="settingsStore.sidebarCollapsed"
              @change="settingsStore.toggleSidebar"
              active-text="折叠"
              inactive-text="展开"
            />
          </el-form-item>
        </el-form>
      </el-card>

      <!-- Logo Settings Card -->
      <el-card class="settings-card">
        <template #header>
          <div class="card-header">
            <el-icon><Picture /></el-icon>
            <span>Logo设置</span>
          </div>
        </template>

        <el-form label-position="top">
          <el-form-item label="系统Logo">
            <el-input
              v-model="settingsStore.logoUrl"
              placeholder="请输入Logo图片URL"
              clearable
              style="width: 100%; max-width: 400px;"
            />
          </el-form-item>

          <el-form-item label="Logo预览">
            <div class="logo-preview">
              <div class="logo-box" :style="{ backgroundColor: settingsStore.headerBgColor }">
                <img
                  v-if="settingsStore.logoUrl"
                  :src="settingsStore.logoUrl"
                  alt="logo"
                  class="preview-logo"
                  @error="handleLogoError"
                >
                <img v-else src="/favicon.svg" alt="default logo" class="preview-logo">
                <span class="logo-text" :style="{ color: settingsStore.headerTextColor }">MOM</span>
              </div>
            </div>
          </el-form-item>

          <el-form-item label="预设Logo">
            <div class="preset-logos">
              <div class="preset-logo-item" @click="settingsStore.setLogoUrl('')">
                <img src="/favicon.svg" alt="默认" class="preset-logo">
                <span>默认</span>
              </div>
            </div>
          </el-form-item>
        </el-form>
      </el-card>

      <!-- Preview Card -->
      <el-card class="settings-card preview-card">
        <template #header>
          <div class="card-header">
            <el-icon><View /></el-icon>
            <span>效果预览</span>
          </div>
        </template>

        <div class="preview-frame">
          <div class="preview-topbar" :style="{ backgroundColor: settingsStore.headerBgColor }">
            <div class="preview-logo-mini">
              <img src="/favicon.svg" alt="logo" class="preview-logo-mini-img">
              <span :style="{ color: settingsStore.headerTextColor }">MOM</span>
            </div>
            <div class="preview-tabs">
              <span class="preview-tab active">生产执行</span>
              <span class="preview-tab">质量管理</span>
            </div>
          </div>
          <div class="preview-body" :style="{ backgroundColor: settingsStore.pageBgColor }">
            <div class="preview-sidebar" :style="{ backgroundColor: settingsStore.sidebarBgColor }">
              <div class="preview-menu-item active" :style="{ color: settingsStore.sidebarActiveTextColor }">
                <span>生产订单</span>
              </div>
              <div class="preview-menu-item" :style="{ color: settingsStore.sidebarTextColor }">
                <span>工序报工</span>
              </div>
              <div class="preview-menu-item" :style="{ color: settingsStore.sidebarTextColor }">
                <span>包装条码</span>
              </div>
            </div>
            <div class="preview-content" :style="{ backgroundColor: settingsStore.cardBgColor }">
              <p class="preview-text" :style="{ fontFamily: settingsStore.fontFamilyCSS }">
                这是一段预览文本，用于查看主题效果。
              </p>
              <p class="preview-text-en">The quick brown fox jumps over the lazy dog.</p>
            </div>
          </div>
        </div>
      </el-card>

      <!-- System Info Card -->
      <el-card class="settings-card">
        <template #header>
          <div class="card-header">
            <el-icon><InfoFilled /></el-icon>
            <span>系统信息</span>
          </div>
        </template>

        <el-descriptions :column="2" border>
          <el-descriptions-item label="系统版本">v1.0.0</el-descriptions-item>
          <el-descriptions-item label="构建时间">2026-04-02</el-descriptions-item>
          <el-descriptions-item label="前端框架">Vue 3 + Element Plus</el-descriptions-item>
          <el-descriptions-item label="设计规范">MOM3.0 UI规范 V1.0</el-descriptions-item>
        </el-descriptions>
      </el-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Sunny, Moon, Brush, Picture, View, InfoFilled } from '@element-plus/icons-vue'
import { useSettingsStore } from '@/stores/settings'

const settingsStore = useSettingsStore()

const handleLogoError = () => {
  settingsStore.setLogoUrl('')
}
</script>

<script lang="ts">
export default { name: 'SystemConfig' }
</script>

<style scoped lang="scss">
.page-container {
  padding: 16px;
}

.page-header {
  margin-bottom: 20px;
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  margin: 0;
}

.settings-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;

  @media (max-width: 1200px) {
    grid-template-columns: 1fr;
  }
}

.settings-card {
  :deep(.el-card__header) {
    padding: 12px 16px;
    border-bottom: 1px solid var(--border-color, #e2e8f0);
  }

  :deep(.el-card__body) {
    padding: 20px;
  }
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
  color: var(--el-text-color-primary);

  .el-icon {
    color: var(--el-color-primary);
  }
}

.theme-radio-group {
  display: flex;
  gap: 16px;

  :deep(.el-radio) {
    margin-right: 0;
  }

  :deep(.el-radio__label) {
    padding-left: 8px;
  }
}

.theme-option {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 4px 0;
}

.color-picker-row {
  display: flex;
  align-items: center;
  gap: 12px;

  .color-value {
    font-family: monospace;
    color: var(--el-text-color-secondary);
    font-size: 13px;
  }
}

.logo-preview {
  margin-top: 8px;
}

.logo-box {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-radius: 8px;
  width: fit-content;
  transition: background-color 0.3s;

  .preview-logo {
    width: 32px;
    height: 32px;
  }

  .logo-text {
    font-size: 16px;
    font-weight: 600;
    letter-spacing: 2px;
    transition: color 0.3s;
  }
}

.preset-logos {
  display: flex;
  gap: 12px;
  margin-top: 8px;

  .preset-logo-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    cursor: pointer;
    padding: 8px 16px;
    border-radius: 6px;
    border: 1px solid var(--border-color, #e2e8f0);
    transition: all 0.2s;

    &:hover {
      border-color: var(--el-color-primary);
      background: var(--el-color-primary-light-9);
    }

    .preset-logo {
      width: 28px;
      height: 28px;
      margin-bottom: 4px;
    }

    span {
      font-size: 12px;
      color: var(--el-text-color-secondary);
    }
  }
}

.preview-card {
  grid-column: span 2;

  @media (max-width: 1200px) {
    grid-column: span 1;
  }
}

.preview-frame {
  border: 1px solid var(--border-color, #e2e8f0);
  border-radius: 8px;
  overflow: hidden;
  height: 280px;
}

.preview-topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 12px;
  height: 40px;
  transition: background-color 0.3s;

  .preview-logo-mini {
    display: flex;
    align-items: center;
    gap: 6px;

    .preview-logo-mini-img {
      width: 20px;
      height: 20px;
    }

    span {
      font-size: 12px;
      font-weight: 600;
      letter-spacing: 1px;
    }
  }

  .preview-tabs {
    display: flex;
    gap: 4px;

    .preview-tab {
      padding: 4px 12px;
      font-size: 11px;
      color: rgba(255, 255, 255, 0.6);
      border-radius: 4px;
      cursor: pointer;

      &.active {
        color: #fff;
        background: rgba(255, 255, 255, 0.15);
      }
    }
  }
}

.preview-body {
  display: flex;
  height: calc(100% - 40px);
  transition: background-color 0.3s;
}

.preview-sidebar {
  width: 140px;
  padding: 8px;
  transition: background-color 0.3s;

  .preview-menu-item {
    padding: 8px 12px;
    font-size: 12px;
    border-radius: 4px;
    cursor: pointer;
    transition: all 0.2s;

    &.active {
      background: rgba(37, 99, 235, 0.15);
    }

    &:hover:not(.active) {
      background: rgba(255, 255, 255, 0.05);
    }
  }
}

.preview-content {
  flex: 1;
  padding: 16px;
  transition: background-color 0.3s;

  .preview-text {
    font-size: 13px;
    color: var(--el-text-color-primary);
    margin: 0 0 8px 0;
    line-height: 1.6;
  }

  .preview-text-en {
    font-size: 12px;
    color: var(--el-text-color-secondary);
    margin: 0;
  }
}
</style>
