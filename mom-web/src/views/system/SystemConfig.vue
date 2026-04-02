<template>
  <div class="system-config">
    <el-card>
      <el-form label-width="120px">
        <el-divider content-position="left">主题设置</el-divider>

        <el-form-item label="主题">
          <el-radio-group v-model="settingsStore.theme">
            <el-radio value="light">
              <el-icon><Sunny /></el-icon>浅色
            </el-radio>
            <el-radio value="dark">
              <el-icon><Moon /></el-icon>深色
            </el-radio>
            <el-radio value="custom">
              <el-icon><Brush /></el-icon>自定义
            </el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item v-if="settingsStore.theme === 'custom'" label="主题颜色">
          <el-color-picker v-model="settingsStore.customThemeColor" @change="settingsStore.applyTheme" />
          <span class="color-tip">选择您喜欢的颜色作为主题色</span>
        </el-form-item>

        <el-form-item label="字体">
          <el-select v-model="settingsStore.fontFamily" @change="settingsStore.applyFontFamily">
            <el-option value="default" label="系统默认字体" />
            <el-option value="helvetica" label="Helvetica" />
            <el-option value="pingfang" label="苹方字体" />
            <el-option value="microsoft" label="微软雅黑" />
          </el-select>
        </el-form-item>

        <el-form-item label="侧边栏折叠">
          <el-switch
            v-model="settingsStore.sidebarCollapsed"
            @change="settingsStore.toggleSidebar"
          />
        </el-form-item>

        <el-divider content-position="left">Logo设置</el-divider>

        <el-form-item label="系统Logo">
          <div class="logo-upload">
            <el-input v-model="settingsStore.logoUrl" placeholder="请输入Logo图片URL" clearable />
            <div class="logo-preview">
              <img v-if="settingsStore.logoUrl" :src="settingsStore.logoUrl" alt="logo" class="preview-logo" @error="handleLogoError">
              <img v-else src="/favicon.svg" alt="default logo" class="preview-logo">
              <span class="logo-tip">输入Logo图片URL后预览</span>
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

        <el-divider content-position="left">预览</el-divider>

        <el-form-item label="预览">
          <div class="preview-box">
            <div class="preview-sidebar" :style="{ backgroundColor: settingsStore.sidebarBgColor }">
              <div
                class="preview-sidebar-text"
                :style="{ color: settingsStore.sidebarTextColor }"
              >
                侧边栏
              </div>
            </div>
            <div class="preview-content">
              <div class="preview-text" :style="{ fontFamily: settingsStore.fontFamilyCSS }">
                <p>这是一段预览文本，用于查看字体效果。</p>
                <p>The quick brown fox jumps over the lazy dog.</p>
              </div>
            </div>
          </div>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="mt-16">
      <el-divider content-position="left">系统信息</el-divider>
      <el-descriptions :column="2" border>
        <el-descriptions-item label="系统版本">v1.0.0</el-descriptions-item>
        <el-descriptions-item label="最后更新时间">2026-04-02</el-descriptions-item>
      </el-descriptions>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { Sunny, Moon, Brush } from '@element-plus/icons-vue'
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
.system-config {
  .mt-16 {
    margin-top: 16px;
  }

  .color-tip {
    margin-left: 12px;
    color: #999;
    font-size: 12px;
  }

  .preview-box {
    display: flex;
    height: 200px;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    overflow: hidden;
  }

  .preview-sidebar {
    width: 200px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background-color 0.3s;
  }

  .preview-sidebar-text {
    transition: color 0.3s;
  }

  .preview-content {
    flex: 1;
    padding: 16px;
    background: #fff;
  }

  .preview-text {
    p {
      margin: 8px 0;
      line-height: 1.6;
    }
  }

  .logo-upload {
    width: 100%;

    .el-input {
      width: 400px;
      margin-bottom: 16px;
    }

    .logo-preview {
      display: flex;
      align-items: center;

      .preview-logo {
        width: 48px;
        height: 48px;
        margin-right: 12px;
        border-radius: 4px;
        border: 1px solid #dcdfe6;
        padding: 4px;
        background: #fff;
      }

      .logo-tip {
        color: #999;
        font-size: 12px;
      }
    }
  }

  .preset-logos {
    display: flex;
    gap: 12px;

    .preset-logo-item {
      display: flex;
      flex-direction: column;
      align-items: center;
      cursor: pointer;
      padding: 8px;
      border-radius: 4px;
      border: 1px solid #dcdfe6;

      &:hover {
        border-color: var(--el-color-primary);
      }

      .preset-logo {
        width: 32px;
        height: 32px;
        margin-bottom: 4px;
      }

      span {
        font-size: 12px;
        color: #666;
      }
    }
  }
}
</style>
