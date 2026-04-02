<template>
  <el-container class="main-layout">
    <!-- 侧边栏 -->
    <el-aside :width="isCollapse ? '64px' : '200px'" class="sidebar">
      <div class="logo">
        <img v-if="!isCollapse" :src="settingsStore.logoUrl || '/favicon.svg'" alt="logo" class="logo-img" @error="handleLogoError">
        <span v-if="!isCollapse" class="logo-text">MOM3.0</span>
      </div>

      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :router="true"
        class="sidebar-menu"
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409eff"
      >
        <template v-for="menu in dynamicMenus" :key="menu.id">
          <el-sub-menu v-if="menu.children && menu.children.length > 0" :index="menu.path">
            <template #title>
              <el-icon><component :is="getIcon(menu.icon)" /></el-icon>
              <span>{{ menu.menu_name }}</span>
            </template>
            <el-menu-item v-for="child in menu.children" :key="child.id" :index="child.path">
              {{ child.menu_name }}
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item v-else :index="menu.path">
            <el-icon><component :is="getIcon(menu.icon)" /></el-icon>
            <template #title>{{ menu.menu_name }}</template>
          </el-menu-item>
        </template>
      </el-menu>
    </el-aside>

    <el-container>
      <!-- 头部 -->
      <el-header class="header">
        <div class="header-left">
          <el-icon class="collapse-icon" @click="isCollapse = !isCollapse">
            <Fold v-if="!isCollapse" />
            <Expand v-else />
          </el-icon>
        </div>

        <div class="header-right">
          <!-- 快捷设置 -->
          <el-dropdown @command="handleQuickSettings" trigger="click">
            <span class="quick-settings">
              <el-icon><Brush /></el-icon>
              <span>设置</span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>
                  <span>主题</span>
                  <el-radio-group v-model="settingsStore.theme" size="small" @click.stop>
                    <el-radio-button value="light">浅色</el-radio-button>
                    <el-radio-button value="dark">深色</el-radio-button>
                  </el-radio-group>
                </el-dropdown-item>
                <el-dropdown-item>
                  <span>字体</span>
                  <el-select v-model="settingsStore.fontFamily" size="small" style="width: 100px" @click.stop>
                    <el-option value="default" label="默认" />
                    <el-option value="helvetica" label="Helvetica" />
                    <el-option value="microsoft" label="微软雅黑" />
                  </el-select>
                </el-dropdown-item>
                <el-dropdown-item command="system-config" divided>系统设置</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <el-dropdown @command="handleCommand">
            <span class="user-info">
              <el-avatar :size="32" :src="userInfo?.avatar">
                {{ userInfo?.nickname?.charAt(0) || 'U' }}
              </el-avatar>
              <span class="username">{{ userInfo?.nickname || userInfo?.username }}</span>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人中心</el-dropdown-item>
                <el-dropdown-item command="password">修改密码</el-dropdown-item>
                <el-dropdown-item divided command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 内容区 -->
      <el-main class="main-content">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useSettingsStore } from '@/stores/settings'
import { ElMessageBox } from 'element-plus'
import {
  House, Setting, Box, List, Monitor, House as Warehouse,
  CircleCheck, Calendar, Search, Lightning, Fold, Expand, ArrowDown, Brush
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const settingsStore = useSettingsStore()

const isCollapse = ref(false)
const userInfo = computed(() => authStore.userInfo)

// 动态菜单 - 从用户信息获取
const dynamicMenus = computed(() => {
  if (!authStore.userInfo?.menus) return []
  return authStore.userInfo.menus.filter((m: any) => m.menu_type === 'M' || m.menu_type === 'C')
})

// 图标映射
const iconMap: Record<string, any> = {
  'House': House,
  'Setting': Setting,
  'Box': Box,
  'List': List,
  'Monitor': Monitor,
  'Warehouse': Warehouse,
  'CircleCheck': CircleCheck,
  'Calendar': Calendar,
  'Search': Search,
  'Lightning': Lightning
}

const getIcon = (iconName?: string) => {
  if (!iconName) return House
  return iconMap[iconName] || House
}

const activeMenu = computed(() => route.path)

const handleLogoError = () => {
  settingsStore.setLogoUrl('')
}

const handleQuickSettings = (command: string) => {
  if (command === 'system-config') {
    router.push('/system/config')
  }
}

const handleCommand = async (command: string) => {
  switch (command) {
    case 'logout':
      await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
        type: 'warning'
      })
      authStore.logoutAction()
      break
    case 'profile':
      router.push('/profile')
      break
    case 'password':
      router.push('/password')
      break
  }
}
</script>

<style scoped lang="scss">
.main-layout {
  height: 100vh;
}

.sidebar {
  background-color: #304156;
  transition: width 0.3s;

  .logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 60px;
    background: #2b3a4b;

    .logo-img {
      width: 32px;
      height: 32px;
    }

    .logo-text {
      margin-left: 8px;
      font-size: 18px;
      font-weight: 600;
      color: #fff;
    }
  }

  .sidebar-menu {
    border-right: none;
    height: calc(100vh - 60px);
    overflow-y: auto;
  }
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
  border-bottom: 1px solid #e4e7ed;
  padding: 0 16px;

  .header-left {
    .collapse-icon {
      font-size: 20px;
      cursor: pointer;
      &:hover {
        color: #409eff;
      }
    }
  }

  .header-right {
    display: flex;
    align-items: center;
    gap: 16px;

    .quick-settings {
      display: flex;
      align-items: center;
      cursor: pointer;
      padding: 4px 8px;
      border-radius: 4px;
      color: #666;

      &:hover {
        background: #f5f7fa;
        color: var(--el-color-primary);
      }
    }

    .user-info {
      display: flex;
      align-items: center;
      cursor: pointer;

      .username {
        margin: 0 8px;
      }
    }
  }
}

.main-content {
  background: #f5f7fa;
  padding: 16px;
}
</style>
