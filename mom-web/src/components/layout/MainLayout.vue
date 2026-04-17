<template>
  <el-container class="main-layout">
    <!-- Left Sidebar Menu -->
    <el-aside :style="{ width: sidebarWidth }" class="sidebar">
      <div class="logo" :style="{ backgroundColor: '#0f172a' }">
        <img v-if="!isCollapse" :src="settingsStore.logoUrl || '/favicon.svg'" alt="logo" class="logo-img" @error="handleLogoError">
        <span v-if="!isCollapse" class="logo-text">MOM</span>
        <div class="home-btn" @click="router.push('/dashboard')" title="首页">
          <el-icon><House /></el-icon>
        </div>
      </div>

      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :router="true"
        class="sidebar-menu"
        :background-color="settingsStore.sidebarBgColor"
        :text-color="settingsStore.sidebarTextColor"
        :active-text-color="settingsStore.sidebarActiveTextColor"
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
      <!-- Top Header -->
      <el-header class="top-header" :style="{ backgroundColor: settingsStore.headerBgColor }">
        <div class="tabs-container">
          <!-- Collapse Toggle -->
          <el-tooltip :content="isCollapse ? '展开菜单' : '折叠菜单'" placement="bottom">
            <span class="collapse-btn" @click="toggleCollapse">
              <el-icon :size="18"><component :is="isCollapse ? Expand : Fold" /></el-icon>
            </span>
          </el-tooltip>
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
          <!-- Language Switch -->
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

          <!-- Menu Search -->
          <el-popover placement="bottom-end" :width="320" trigger="click">
            <template #reference>
              <span class="header-btn">
                <el-icon><Search /></el-icon>
              </span>
            </template>
            <div class="menu-search">
              <el-input v-model="searchQuery" placeholder="搜索菜单..." clearable size="small" @input="handleSearchMenu">
                <template #prefix><el-icon><Search /></el-icon></template>
              </el-input>
              <div class="search-results" v-if="searchResults.length > 0">
                <div v-for="item in searchResults" :key="item.id" class="search-item" @click="handleSearchClick(item)">
                  <el-icon><component :is="getIcon(item.icon)" /></el-icon>
                  <span>{{ item.menu_name }}</span>
                </div>
              </div>
              <el-empty v-else-if="searchQuery && searchResults.length === 0" description="未找到菜单" :image-size="60" />
            </div>
          </el-popover>

          <!-- Theme Quick Switch -->
          <el-dropdown @command="handleThemeChange" trigger="click">
            <span class="header-btn">
              <el-icon><Brush /></el-icon>
              <span>主题</span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>
                  <span>主题模式</span>
                </el-dropdown-item>
                <el-dropdown-item @click.stop="settingsStore.setTheme('light')">
                  <el-icon><Sunny /></el-icon> 浅色
                  <el-icon v-if="settingsStore.theme === 'light'" class="check-icon"><Check /></el-icon>
                </el-dropdown-item>
                <el-dropdown-item @click.stop="settingsStore.setTheme('dark')">
                  <el-icon><Moon /></el-icon> 深色
                  <el-icon v-if="settingsStore.theme === 'dark'" class="check-icon"><Check /></el-icon>
                </el-dropdown-item>
                <el-dropdown-item @click.stop="settingsStore.setTheme('custom')">
                  <el-icon><Brush /></el-icon> 自定义
                  <el-icon v-if="settingsStore.theme === 'custom'" class="check-icon"><Check /></el-icon>
                </el-dropdown-item>
                <el-dropdown-item divided command="system-config">系统设置</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

          <!-- User Menu -->
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

      <!-- Main Content -->
      <el-main class="main-content" :style="{ backgroundColor: settingsStore.pageBgColor }">
        <router-view />
      </el-main>

      <!-- AI Chat -->
      <AiChatButton />
      <AiChatWindow />
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useSettingsStore } from '@/stores/settings'
import AiChatButton from '@/components/ai-chat/AiChatButton.vue'
import AiChatWindow from '@/components/ai-chat/AiChatWindow.vue'
import { ElMessageBox } from 'element-plus'
import {
  House, Setting, Box, List, Monitor, House as WarehouseIcon,
  CircleCheck, Calendar, Search, Lightning, Fold, Expand, ArrowDown, Brush, Switch as SwitchIcon, Sunny, Moon, Check as CheckIcon,
  User, ShoppingCart, HomeFilled, Location, PriceTag, OfficeBuilding, Key, Connection, Grid,
  Bell, Document, Tools, Tickets, Top, DataAnalysis, Download, Refresh,
  DocumentCopy, TrendCharts, Clock, DataLine, Warning, Upload, DocumentChecked,
  Files, Postcard, Operation, Edit, Close, Collection, Folder, ChatDotRound, Checked, Guide, Menu, SetUp
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const settingsStore = useSettingsStore()

const isCollapse = ref(false)
const sidebarWidth = computed(() => isCollapse.value ? '64px' : '200px')
const userInfo = computed(() => authStore.userInfo)
const currentLang = ref(localStorage.getItem('language') || 'zh')
const tabsScroll = ref<HTMLElement | null>(null)
const containerWidth = ref(window.innerWidth)
const searchQuery = ref('')
const searchResults = ref<any[]>([])

const handleResize = () => {
  containerWidth.value = window.innerWidth
}

onMounted(() => {
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})

// Toggle sidebar collapse
const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}

// Tab name mapping
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

const getTabName = (tab: any) => {
  if (currentLang.value === 'en' && tabNameMap[tab.menu_name]) {
    return tabNameMap[tab.menu_name]
  }
  return tab.menu_name
}

const handleLanguageChange = (lang: string) => {
  currentLang.value = lang
  localStorage.setItem('language', lang)
}

const handleThemeChange = (command: string) => {
  if (command === 'system-config') {
    router.push('/system/config')
  }
}

const getMenuPath = (menu: any) => {
  if (menu.path && menu.path.startsWith('/')) {
    return menu.path
  }

  let parentPath = ''
  if (menu.parent_id && menu.parent_id > 0) {
    const parentMenu = allMenus.value.find((m: any) => m.id === menu.parent_id)
    if (parentMenu) {
      parentPath = getMenuPath(parentMenu)
    }
  }

  const currentPath = menu.path || ''
  if (!currentPath) return '/'
  return parentPath ? `${parentPath}/${currentPath}` : `/${currentPath}`
}

const allMenus = computed(() => {
  if (!authStore.userInfo?.menus) return []
  return authStore.userInfo.menus
})

const topLevelMenus = computed(() => {
  // 只显示模块菜单(M)，首页已移至侧边栏快捷按钮
  const moduleMenus = allMenus.value.filter((m: any) => m.menu_type === 'M')
  return moduleMenus
})

const activeTab = ref<number>(0)

const updateActiveTab = () => {
  const path = route.path
  // 首先检查是否是首页
  if (path === '/dashboard' || path === '/') {
    // 默认激活第一个顶级菜单
    if (topLevelMenus.value.length > 0) {
      activeTab.value = topLevelMenus.value[0].id
    }
    return
  }
  for (const tab of topLevelMenus.value) {
    if (path.startsWith('/' + tab.path)) {
      activeTab.value = tab.id
      return
    }
  }
  // 如果没找到匹配，默认激活第一个
  if (topLevelMenus.value.length > 0 && !activeTab.value) {
    activeTab.value = topLevelMenus.value[0].id
  }
}

const currentSubMenus = computed(() => {
  const currentTopMenu = topLevelMenus.value.find((m: any) => m.id === activeTab.value)
  if (!currentTopMenu || !currentTopMenu.children) return []
  return currentTopMenu.children.filter((m: any) => m.menu_type === 'C')
})

const handleTabClick = (tab: any) => {
  activeTab.value = tab.id
  // 如果是首页，直接跳转到dashboard并收起侧边栏
  if (tab.path === '/dashboard') {
    if (!isCollapse.value) {
      isCollapse.value = true
    }
    router.push('/dashboard')
    return
  }
  const firstChild = currentSubMenus.value[0]
  if (firstChild) {
    router.push(getMenuPath(firstChild))
  }
}

const handleMoreTabClick = (tabId: number) => {
  const tab = topLevelMenus.value.find((m: any) => m.id === tabId)
  if (tab) {
    handleTabClick(tab)
  }
}

const iconMap: Record<string, any> = {
  'House': House,
  'Setting': Setting,
  'Box': Box,
  'List': List,
  'Monitor': Monitor,
  'Warehouse': WarehouseIcon,
  'CircleCheck': CircleCheck,
  'Calendar': Calendar,
  'Search': Search,
  'Lightning': Lightning,
  'User': User,
  'ShoppingCart': ShoppingCart,
  'HomeFilled': HomeFilled,
  'Location': Location,
  'PriceTag': PriceTag,
  'OfficeBuilding': OfficeBuilding,
  'Key': Key,
  'Connection': Connection,
  'Grid': Grid,
  'Bell': Bell,
  'Document': Document,
  'Check': CheckIcon,
  'Tools': Tools,
  'Tickets': Tickets,
  'Top': Top,
  'DataAnalysis': DataAnalysis,
  'Download': Download,
  'Switch': SwitchIcon,
  'Refresh': Refresh,
  'DocumentCopy': DocumentCopy,
  'TrendCharts': TrendCharts,
  'Clock': Clock,
  'DataLine': DataLine,
  'Warning': Warning,
  'Upload': Upload,
  'DocumentChecked': DocumentChecked,
  'Files': Files,
  'Postcard': Postcard,
  'Operation': Operation,
  'Edit': Edit,
  'Close': Close,
  'Collection': Collection,
  'Folder': Folder,
  'ChatDotRound': ChatDotRound,
  'Checked': Checked,
  'Guide': Guide,
  'Menu': Menu,
  'SetUp': SetUp,
}

const getIcon = (iconName?: string) => {
  if (!iconName) return House
  return iconMap[iconName] || House // 如果找不到图标，返回默认图标
}

const handleSearchMenu = () => {
  if (!searchQuery.value.trim()) {
    searchResults.value = []
    return
  }
  const query = searchQuery.value.toLowerCase()
  const flatMenus = flattenMenus(allMenus.value)
  searchResults.value = flatMenus
    .filter((m: any) => m.menu_name && m.menu_name.toLowerCase().includes(query))
    .slice(0, 10) // 最多显示10条
}

const flattenMenus = (menus: any[]): any[] => {
  const result: any[] = []
  menus.forEach(m => {
    result.push(m)
    if (m.children && m.children.length > 0) {
      result.push(...flattenMenus(m.children))
    }
  })
  return result
}

const handleSearchClick = (menu: any) => {
  searchQuery.value = ''
  searchResults.value = []
  // 找到对应的顶级菜单并激活
  const topMenu = findTopMenu(allMenus.value, menu.id)
  if (topMenu) {
    activeTab.value = topMenu.id
  }
  router.push(getMenuPath(menu))
}

const findTopMenu = (menus: any[], targetId: number): any => {
  const parentMap = new Map<number, any>()
  const buildMap = (items: any[]) => {
    items.forEach(m => {
      if (m.children) {
        m.children.forEach((c: any) => parentMap.set(c.id, m))
        buildMap(m.children)
      }
    })
  }
  buildMap(menus)
  return parentMap.get(targetId)
}

const activeMenu = computed(() => route.path)

const handleLogoError = () => {
  settingsStore.setLogoUrl('')
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

watch(() => route.path, updateActiveTab, { immediate: true })

const tabItemWidth = 110 // 每个tab的固定宽度
const moreButtonWidth = 60 // 更多按钮宽度

const sidebarWidthValue = computed(() => isCollapse.value ? 64 : 200)

const sidebarActualWidth = computed(() => {
  // 强制依赖isCollapse，确保侧边栏变化时重新计算
  void isCollapse.value
  return isCollapse.value ? 64 : 200
})

const visibleTabs = computed(() => {
  const width = containerWidth.value - sidebarActualWidth.value - 180
  const availableWidth = Math.max(100, width - moreButtonWidth)
  const maxTabs = Math.max(1, Math.floor(availableWidth / tabItemWidth))
  return topLevelMenus.value.slice(0, maxTabs)
})

const hiddenTabs = computed(() => {
  const width = containerWidth.value - sidebarActualWidth.value - 180
  const availableWidth = Math.max(100, width - moreButtonWidth)
  const maxTabs = Math.max(1, Math.floor(availableWidth / tabItemWidth))
  return topLevelMenus.value.slice(maxTabs)
})
</script>

<style scoped lang="scss">
.main-layout {
  height: 100vh;
}

.sidebar {
  background-color: v-bind('settingsStore.sidebarBgColor');
  transition: width 0.3s, background-color 0.3s;

  .logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 50px;
    background: #0f172a;
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);

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

    .home-btn {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 32px;
      height: 32px;
      margin-left: 12px;
      color: rgba(255, 255, 255, 0.75);
      cursor: pointer;
      border-radius: 6px;
      transition: all 0.2s;

      &:hover {
        background: rgba(255, 255, 255, 0.1);
        color: #fff;
      }

      .el-icon {
        font-size: 18px;
      }
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
        background-color: rgba(37, 99, 235, 0.1) !important;
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
  background: v-bind('settingsStore.headerBgColor');
  border-bottom: none;
  padding: 0 12px;
  height: 50px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.15);
  transition: background-color 0.3s;

  .tabs-container {
    display: flex;
    align-items: center;
    height: 100%;
    overflow: hidden;

    .collapse-btn {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 32px;
      height: 32px;
      margin-right: 8px;
      color: rgba(255, 255, 255, 0.85);
      cursor: pointer;
      border-radius: 4px;
      transition: all 0.25s ease;

      &:hover {
        background: rgba(255, 255, 255, 0.15);
        color: #fff;
      }
    }

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
      color: rgba(255, 255, 255, 0.65);
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
        background: rgba(255, 255, 255, 0.15);

        &::after {
          content: '';
          position: absolute;
          bottom: 0;
          left: 50%;
          transform: translateX(-50%);
          width: 70%;
          height: 3px;
          background: #2563eb;
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
    border-left: 1px solid rgba(255, 255, 255, 0.15);

    .header-btn {
      display: flex;
      align-items: center;
      gap: 5px;
      cursor: pointer;
      padding: 5px 10px;
      border-radius: 4px;
      color: rgba(255, 255, 255, 0.85);
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
        background: rgba(37, 99, 235, 0.8);
        color: #fff;
        font-size: 12px;
      }
    }
  }
}

.main-content {
  background: v-bind('settingsStore.pageBgColor');
  padding: 12px;
  height: calc(100vh - 50px);
  overflow-y: auto;
  transition: background-color 0.3s;
}

.check-icon {
  margin-left: auto;
  color: #2563eb;
}

.menu-search {
  .search-results {
    max-height: 400px;
    overflow-y: auto;
    margin-top: 8px;

    .search-item {
      display: flex;
      align-items: center;
      gap: 8px;
      padding: 8px 12px;
      cursor: pointer;
      border-radius: 4px;
      transition: background 0.2s;

      &:hover {
        background: #f5f7fa;
      }

      .el-icon {
        color: #909399;
      }
    }
  }
}
</style>

<style lang="scss">
// Global theme class
.dark {
  .main-content {
    background: #0f172a;
  }
}
</style>
