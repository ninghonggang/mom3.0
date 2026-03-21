import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login, logout, getUserInfo, refreshToken } from '@/api/auth'
import router from '@/router'

export interface UserInfo {
  id: number
  username: string
  nickname: string
  avatar?: string
  email?: string
  phone?: string
  roles: string[]
  perms: string[]
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string>(localStorage.getItem('token') || '')
  const refreshTokenVal = ref<string>(localStorage.getItem('refresh_token') || '')
  const userInfo = ref<UserInfo | null>(null)

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => userInfo.value?.roles.includes('admin'))

  // 登录
  const loginAction = async (username: string, password: string) => {
    const res = await login({ username, password })
    token.value = res.data.access_token
    refreshTokenVal.value = res.data.refresh_token
    localStorage.setItem('token', token.value)
    localStorage.setItem('refresh_token', refreshTokenVal.value)
    await getUserInfoAction()
  }

  // 获取用户信息
  const getUserInfoAction = async () => {
    if (!token.value) return
    try {
      const res = await getUserInfo()
      userInfo.value = res.data
    } catch (error) {
      logoutAction()
    }
  }

  // 登出
  const logoutAction = async () => {
    try {
      await logout()
    } catch (error) {
      // ignore
    }
    token.value = ''
    refreshTokenVal.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('refresh_token')
    router.push('/login')
  }

  // 检查权限
  const hasPermission = (perm: string) => {
    if (!userInfo.value) return false
    if (userInfo.value.roles.includes('admin')) return true
    return userInfo.value.perms.includes(perm)
  }

  // 检查角色
  const hasRole = (role: string) => {
    if (!userInfo.value) return false
    return userInfo.value.roles.includes(role)
  }

  return {
    token,
    userInfo,
    isLoggedIn,
    isAdmin,
    loginAction,
    getUserInfoAction,
    logoutAction,
    hasPermission,
    hasRole
  }
})
