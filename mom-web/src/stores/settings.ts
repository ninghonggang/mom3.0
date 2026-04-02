import { defineStore } from 'pinia'
import { ref, watch, computed } from 'vue'

export type Theme = 'light' | 'dark' | 'custom'
export type FontFamily = 'default' | 'helvetica' | 'pingfang' | 'microsoft'

const STORAGE_KEY = 'app_settings'

export const useSettingsStore = defineStore('settings', () => {
  const stored = localStorage.getItem(STORAGE_KEY)
  const defaultSettings = stored ? JSON.parse(stored) : {
    theme: 'light' as Theme,
    customThemeColor: '#409eff',
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
    return theme.value === 'custom' ? customThemeColor.value : getDefaultThemeColor(theme.value)
  })

  const getDefaultThemeColor = (t: Theme) => {
    return t === 'dark' ? '#1f1f1f' : '#409eff'
  }

  const sidebarBgColor = computed(() => {
    if (theme.value === 'custom') return '#2b3a4b'
    return isDark.value ? '#1f1f1f' : '#304156'
  })

  const sidebarTextColor = computed(() => {
    return isDark.value ? '#e0e0e0' : '#bfcbd9'
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
    // Apply custom theme color
    document.documentElement.style.setProperty('--el-color-primary', effectiveThemeColor.value)
  }

  const applyFontFamily = () => {
    document.documentElement.style.setProperty('--el-font-family', fontFamilyCSS.value)
  }

  // 初始化
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
    setTheme,
    setCustomThemeColor,
    setFontFamily,
    setLogoUrl,
    toggleSidebar,
    applyTheme,
    applyFontFamily
  }
})
