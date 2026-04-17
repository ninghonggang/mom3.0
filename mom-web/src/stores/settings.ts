import { defineStore } from 'pinia'
import { ref, watch, computed } from 'vue'

export type Theme = 'light' | 'dark' | 'custom'
export type FontFamily = 'default' | 'helvetica' | 'pingfang' | 'microsoft'

const STORAGE_KEY = 'app_settings'

// Design spec colors
const DESIGN_COLORS = {
  primary: '#2563eb',
  primaryLight3: '#60a5fa',
  primaryLight5: '#93c5fd',
  primaryLight7: '#bfdbfe',
  primaryLight8: '#dbeafe',
  primaryLight9: '#eff6ff',
  primaryDark2: '#1d4ed8',
  success: '#10b981',
  warning: '#f59e0b',
  danger: '#ef4444',
  info: '#6b7280',
  headerBg: '#1e293b',
  sidebarBg: '#1f2937',
  pageBg: '#f1f5f9',
  cardBg: '#ffffff',
  borderColor: '#e2e8f0'
} as const

export const useSettingsStore = defineStore('settings', () => {
  const stored = localStorage.getItem(STORAGE_KEY)
  const defaultSettings = stored ? JSON.parse(stored) : {
    theme: 'light' as Theme,
    customThemeColor: '#2563eb',
    fontFamily: 'default' as FontFamily,
    sidebarCollapsed: false,
    logoUrl: ''
  }

  const theme = ref<Theme>(defaultSettings.theme)
  const customThemeColor = ref(defaultSettings.customThemeColor)
  const fontFamily = ref<FontFamily>(defaultSettings.fontFamily)
  const sidebarCollapsed = ref(defaultSettings.sidebarCollapsed)
  const logoUrl = ref(defaultSettings.logoUrl)

  watch([theme, customThemeColor, fontFamily, sidebarCollapsed, logoUrl], () => {
    localStorage.setItem(STORAGE_KEY, JSON.stringify({
      theme: theme.value,
      customThemeColor: customThemeColor.value,
      fontFamily: fontFamily.value,
      sidebarCollapsed: sidebarCollapsed.value,
      logoUrl: logoUrl.value
    }))
    applyTheme()
    applyFontFamily()
  }, { deep: true })

  const isDark = computed(() => theme.value === 'dark')

  const fontFamilyCSS = computed(() => {
    const map: Record<FontFamily, string> = {
      default: "'Helvetica Neue', 'Helvetica', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif",
      helvetica: "'Helvetica Neue', Helvetica, sans-serif",
      pingfang: "'PingFang SC', 'Hiragino Sans GB', sans-serif",
      microsoft: "'Microsoft YaHei', sans-serif"
    }
    return map[fontFamily.value]
  })

  const effectiveThemeColor = computed(() => {
    return theme.value === 'custom' ? customThemeColor.value : DESIGN_COLORS.primary
  })

  const sidebarBgColor = computed(() => {
    if (theme.value === 'custom') return '#2b3a4b'
    return theme.value === 'dark' ? '#0f172a' : DESIGN_COLORS.sidebarBg
  })

  const sidebarTextColor = computed(() => {
    return theme.value === 'dark' ? '#e2e8f0' : '#94a3b8'
  })

  const sidebarActiveTextColor = computed(() => {
    return DESIGN_COLORS.primary
  })

  const headerBgColor = computed(() => {
    if (theme.value === 'custom') return '#1e293b'
    return theme.value === 'dark' ? '#0f172a' : DESIGN_COLORS.headerBg
  })

  const headerTextColor = computed(() => {
    return theme.value === 'dark' ? '#f1f5f9' : '#ffffff'
  })

  const pageBgColor = computed(() => {
    return theme.value === 'dark' ? '#0f172a' : DESIGN_COLORS.pageBg
  })

  const cardBgColor = computed(() => {
    return theme.value === 'dark' ? '#1e293b' : DESIGN_COLORS.cardBg
  })

  const setTheme = (newTheme: Theme) => {
    theme.value = newTheme
  }

  const setCustomThemeColor = (color: string) => {
    customThemeColor.value = color
    theme.value = 'custom'
  }

  const setFontFamily = (font: FontFamily) => {
    fontFamily.value = font
  }

  const setLogoUrl = (url: string) => {
    logoUrl.value = url
  }

  const toggleSidebar = () => {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  const applyTheme = () => {
    document.documentElement.setAttribute('data-theme', theme.value)
    document.body.classList.toggle('dark', isDark.value)

    // Apply primary color
    document.documentElement.style.setProperty('--el-color-primary', effectiveThemeColor.value)
    document.documentElement.style.setProperty('--el-color-primary-light-3', DESIGN_COLORS.primaryLight3)
    document.documentElement.style.setProperty('--el-color-primary-light-5', DESIGN_COLORS.primaryLight5)
    document.documentElement.style.setProperty('--el-color-primary-light-7', DESIGN_COLORS.primaryLight7)
    document.documentElement.style.setProperty('--el-color-primary-light-8', DESIGN_COLORS.primaryLight8)
    document.documentElement.style.setProperty('--el-color-primary-light-9', DESIGN_COLORS.primaryLight9)
    document.documentElement.style.setProperty('--el-color-primary-dark-2', DESIGN_COLORS.primaryDark2)

    // Apply semantic colors
    document.documentElement.style.setProperty('--el-color-success', DESIGN_COLORS.success)
    document.documentElement.style.setProperty('--el-color-warning', DESIGN_COLORS.warning)
    document.documentElement.style.setProperty('--el-color-danger', DESIGN_COLORS.danger)
    document.documentElement.style.setProperty('--el-color-info', DESIGN_COLORS.info)

    // Apply layout colors
    document.documentElement.style.setProperty('--header-bg', headerBgColor.value)
    document.documentElement.style.setProperty('--sidebar-bg', sidebarBgColor.value)
    document.documentElement.style.setProperty('--page-bg', pageBgColor.value)
    document.documentElement.style.setProperty('--card-bg', cardBgColor.value)
    document.documentElement.style.setProperty('--border-color', DESIGN_COLORS.borderColor)
  }

  const applyFontFamily = () => {
    document.documentElement.style.setProperty('--el-font-family', fontFamilyCSS.value)
  }

  // Initialize theme on load
  applyTheme()
  applyFontFamily()

  return {
    theme,
    customThemeColor,
    fontFamily,
    sidebarCollapsed,
    logoUrl,
    isDark,
    fontFamilyCSS,
    effectiveThemeColor,
    sidebarBgColor,
    sidebarTextColor,
    sidebarActiveTextColor,
    headerBgColor,
    headerTextColor,
    pageBgColor,
    cardBgColor,
    setTheme,
    setCustomThemeColor,
    setFontFamily,
    setLogoUrl,
    toggleSidebar,
    applyTheme,
    applyFontFamily
  }
})
