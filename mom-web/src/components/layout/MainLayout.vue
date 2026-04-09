<template>
  <el-container class="main-layout">
    <!-- 左侧菜单 -->
    <el-aside :width="isCollapse ? '64px' : '200px'" class="sidebar">
      <div class="logo">
        <img v-if="!isCollapse" :src="settingsStore.logoUrl || '/favicon.svg'" alt="logo" class="logo-img" @error="handleLogoError">
        <span v-if="!isCollapse" class="logo-text">MOM</span>
      </div>

      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :router="true"
        class="sidebar-menu"
        background-color="#1a1a2e"
        text-color="#a0a0a0"
        active-text-color="#ffffff"
      >
        <template v-for="menu in currentSubMenus" :key="menu.id">
          <el-sub-menu v-if="menu.children && menu.children.length > 0" :index="getMenuPath(menu)">
            <template #title>
              <el-icon><component :is="getIcon(menu.icon)" /></el-icon>
              <span>{{ menu.menu_name }}</span>
            </template>
            <el-menu-item v-for="child in menu.children" :key="child.id" :index="getMenuPath(child)">
              {{ child.menu_name }}
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item v-else :index="getMenuPath(menu)">
            <el-icon><component :is="getIcon(menu.icon)" /></el-icon>
            <template #title>{{ menu.menu_name }}</template>
          </el-menu-item>
        </template>
      </el-menu>
    </el-aside>

    <el-container>
      <!-- 顶部Tab导航 -->
      <el-header class="top-header">
        <div class="tabs-container">
          <!-- 可滚动的Tab列表 -->
          <div class="tabs-scroll" ref="tabsScroll">
            <div
              v-for="tab in visibleTabs"
              :key="tab.id"
              :class="['tab-item', { active: activeTab === tab.id }]"
              @click="handleTabClick(tab)"
            >
              <el-icon><component :is="getIcon(tab.icon)" /></el-icon>
              <span>{{ getTabName(tab) }}</span>
            </div>
          </div>
          <!-- 更多Tab下拉 -->
          <el-dropdown @command="handleMoreTabClick" trigger="click" v-if="hiddenTabs.length > 0">
            <span class="tab-more">
              <span>...</span>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item v-for="tab in hiddenTabs" :key="tab.id" :command="tab.id">
                  <el-icon><component :is="getIcon(tab.icon)" /></el-icon>
                  {{ getTabName(tab) }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>

        <div class="header-right">
          <!-- 语言切换 -->
          <el-dropdown @command="handleLanguageChange" trigger="click">
            <span class="header-btn">
              <el-icon><Switch /></el-icon>
              <span>{{ currentLang === 'zh' ? '中文' : 'EN' }}</span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item :class="{ 'lang-active': currentLang === 'zh' }" command="zh">中文</el-dropdown-item>
                <el-dropdown-item :class="{ 'lang-active': currentLang === 'en' }" command="en">English</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <!-- 快捷设置 -->
          <el-dropdown @command="handleQuickSettings" trigger="click">
            <span class="header-btn">
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
                <el-dropdown-item command="system-config" divided>系统设置</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <el-dropdown @command="handleCommand">
            <span class="header-btn user-btn">
              <el-avatar :size="28" :src="userInfo?.avatar">
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

      <!-- AI Chat -->
      <AiChatButton />
      <AiChatWindow />
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useSettingsStore } from '@/stores/settings'
import AiChatButton from '@/components/ai-chat/AiChatButton.vue'
import AiChatWindow from '@/components/ai-chat/AiChatWindow.vue'
import { ElMessageBox } from 'element-plus'
import {
  House, Setting, Box, List, Monitor, House as Warehouse,
  CircleCheck, Calendar, Search, Lightning, Fold, Expand, ArrowDown, Brush, Switch
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const settingsStore = useSettingsStore()

const isCollapse = ref(false)
const userInfo = computed(() => authStore.userInfo)
const currentLang = ref(localStorage.getItem('language') || 'zh')
const tabsScroll = ref<HTMLElement | null>(null)

// Tab名称映射（中文→英文缩写）
const tabNameMap: Record<string, string> = {
  '系统管理': 'System',
  '主数据': 'MDM',
  '生产执行': 'Production',
  '设备管理': 'Equipment',
  '仓储管理': 'WMS',
  '质量管理': 'Quality',
  'APS计划': 'APS',
  '追溯管理': 'Trace',
  '能源管理': 'Energy'
}

// 获取Tab显示名称
const getTabName = (tab: any) => {
  if (currentLang.value === 'en' && tabNameMap[tab.menu_name]) {
    return tabNameMap[tab.menu_name]
  }
  return tab.menu_name
}

// 语言切换
const handleLanguageChange = (lang: string) => {
  currentLang.value = lang
  localStorage.setItem('language', lang)
}

// 获取完整菜单路径
const getMenuPath = (menu: any) => {
  // 如果有完整路径直接返回
  if (menu.path && menu.path.startsWith('/')) {
    return menu.path
  }

  // 查找父级菜单的路径
  let parentPath = ''
  if (menu.parent_id && menu.parent_id > 0) {
    const parentMenu = allMenus.value.find((m: any) => m.id === menu.parent_id)
    if (parentMenu) {
      // 递归获取父级路径
      parentPath = getMenuPath(parentMenu)
    }
  }

  // 构建完整路径：父路径 + 当前路径
  const currentPath = menu.path || ''
  if (!currentPath) return '/'
  return parentPath ? `${parentPath}/${currentPath}` : `/${currentPath}`
}

// 获取所有菜单（树结构）
const allMenus = computed(() => {
  if (!authStore.userInfo?.menus) return []
  return authStore.userInfo.menus
})

// 顶级菜单（一级菜单，用于Tab导航）
const topLevelMenus = computed(() => {
  return allMenus.value.filter((m: any) => m.menu_type === 'M')
})

// 当前选中的Tab
const activeTab = ref<number>(0)

// 根据当前路由自动选中对应的Tab
const updateActiveTab = () => {
  const path = route.path
  for (const tab of topLevelMenus.value) {
    if (path.startsWith('/' + tab.path)) {
      activeTab.value = tab.id
      return
    }
  }
  if (topLevelMenus.value.length > 0 && !activeTab.value) {
    activeTab.value = topLevelMenus.value[0].id
  }
}

// 当前Tab下的子菜单
const currentSubMenus = computed(() => {
  const currentTopMenu = topLevelMenus.value.find((m: any) => m.id === activeTab.value)
  if (!currentTopMenu || !currentTopMenu.children) return []
  return currentTopMenu.children.filter((m: any) => m.menu_type === 'C')
})

// Tab点击事件
const handleTabClick = (tab: any) => {
  activeTab.value = tab.id
  const firstChild = currentSubMenus.value[0]
  if (firstChild) {
    router.push(getMenuPath(firstChild))
  }
}

// 更多Tab点击
const handleMoreTabClick = (tabId: number) => {
  const tab = topLevelMenus.value.find((m: any) => m.id === tabId)
  if (tab) {
    handleTabClick(tab)
  }
}

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

// 监听路由变化，更新activeTab
watch(() => route.path, updateActiveTab, { immediate: true })

// 可见的Tabs（取前7个）和隐藏的Tabs
const visibleTabs = computed(() => topLevelMenus.value.slice(0, 7))
const hiddenTabs = computed(() => topLevelMenus.value.slice(7))
</script>

<style scoped lang="scss">
// 主题变量
$primary-color: #667eea;
$primary-dark: #764ba2;
$sidebar-bg: #1a1a2e;
$sidebar-text: #a0a0a0;
$sidebar-active: #ffffff;
$header-bg: linear-gradient(135deg, $primary-color 0%, $primary-dark 100%);
$header-text: rgba(255, 255, 255, 0.95);

.main-layout {
  height: 100vh;
}

.sidebar {
  background-color: $sidebar-bg;
  transition: width 0.3s;

  .logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 50px;
    background: #16162a;
    border-bottom: 1px solid #2a2a4a;

    .logo-img {
      width: 28px;
      height: 28px;
    }

    .logo-text {
      font-size: 16px;
      font-weight: 600;
      color: #fff;
      letter-spacing: 2px;
    }
  }

  .sidebar-menu {
    border-right: none;
    height: calc(100vh - 50px);
    overflow-y: auto;

    &.el-menu--collapse {
      width: 64px;
    }

    :deep(.el-menu-item),
    :deep(.el-sub-menu__title) {
      &:hover {
        background-color: #2a2a4a !important;
        color: $sidebar-active !important;
      }
    }

    :deep(.el-sub-menu .el-menu-item) {
      height: 40px;
      line-height: 40px;
      padding-left: 48px !important;
    }
  }
}

.top-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: $header-bg;
  border-bottom: none;
  padding: 0 12px;
  height: 50px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.15);

  .tabs-container {
    display: flex;
    align-items: center;
    height: 100%;
    overflow: hidden;

    .tabs-scroll {
      display: flex;
      align-items: center;
      height: 100%;
      gap: 2px;
      overflow-x: auto;
      scrollbar-width: none;
      -ms-overflow-style: none;

      &::-webkit-scrollbar {
        display: none;
      }
    }

    .tab-item {
      display: flex;
      align-items: center;
      gap: 6px;
      padding: 0 14px;
      height: 100%;
      color: rgba(255, 255, 255, 0.75);
      cursor: pointer;
      font-size: 13px;
      white-space: nowrap;
      transition: all 0.25s ease;
      position: relative;

      &:hover {
        color: #fff;
        background: rgba(255, 255, 255, 0.1);
      }

      &.active {
        color: #fff;
        background: rgba(255, 255, 255, 0.2);

        &::after {
          content: '';
          position: absolute;
          bottom: 0;
          left: 50%;
          transform: translateX(-50%);
          width: 70%;
          height: 3px;
          background: #fff;
          border-radius: 3px 3px 0 0;
        }
      }

      .el-icon {
        font-size: 14px;
      }
    }

    .tab-more {
      display: flex;
      align-items: center;
      gap: 4px;
      padding: 0 10px;
      height: 36px;
      color: rgba(255, 255, 255, 0.85);
      cursor: pointer;
      font-size: 13px;
      border-radius: 4px;
      transition: all 0.25s ease;

      &:hover {
        background: rgba(255, 255, 255, 0.15);
        color: #fff;
      }
    }
  }

  .header-right {
    display: flex;
    align-items: center;
    gap: 8px;
    padding-left: 12px;
    border-left: 1px solid rgba(255, 255, 255, 0.2);

    .header-btn {
      display: flex;
      align-items: center;
      gap: 5px;
      cursor: pointer;
      padding: 5px 10px;
      border-radius: 4px;
      color: $header-text;
      font-size: 12px;
      transition: all 0.25s ease;

      &:hover {
        background: rgba(255, 255, 255, 0.15);
        color: #fff;
      }
    }

    .user-btn {
      padding: 3px 8px;

      .username {
        max-width: 80px;
        overflow: hidden;
        text-overflow: ellipsis;
      }

      :deep(.el-avatar) {
        background: rgba(255, 255, 255, 0.3);
        color: #fff;
        font-size: 12px;
      }
    }
  }
}

.main-content {
  background: #f0f2f5;
  padding: 12px;
  height: calc(100vh - 50px);
  overflow-y: auto;
}
</style>

<style lang="scss">
// 全局样式
.lang-active {
  background: var(--el-color-primary-light-9) !important;
  color: var(--el-color-primary) !important;
}

// 暗色主题覆盖
.dark {
  .main-content {
    background: #1a1a2e;
  }

  .sidebar {
    background-color: #16162a;

    .logo {
      background: #0f0f1e;
    }
  }
}
</style>
