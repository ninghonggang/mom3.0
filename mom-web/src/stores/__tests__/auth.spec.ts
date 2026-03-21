import { describe, it, expect, vi, beforeEach } from 'vitest'
import { setActivePinia, createPinia } from 'pinia'
import { useAuthStore } from '../auth'

// Mock router
vi.mock('@/router', () => ({
  default: {
    push: vi.fn()
  }
}))

// Mock api
vi.mock('@/api/auth', () => ({
  login: vi.fn().mockResolvedValue({
    data: {
      access_token: 'test-token',
      refresh_token: 'test-refresh-token'
    }
  }),
  logout: vi.fn().mockResolvedValue({}),
  getUserInfo: vi.fn().mockResolvedValue({
    data: {
      id: 1,
      username: 'admin',
      nickname: 'Admin',
      roles: ['admin'],
      perms: ['*:*:*']
    }
  }),
  refreshToken: vi.fn().mockResolvedValue({
    data: { access_token: 'new-token' }
  })
}))

describe('auth store', () => {
  beforeEach(() => {
    localStorage.clear()
    setActivePinia(createPinia())
  })

  it('should initialize with empty token', () => {
    const authStore = useAuthStore()
    expect(authStore.token).toBe('')
    expect(authStore.isLoggedIn).toBe(false)
  })

  it('should login successfully', async () => {
    const authStore = useAuthStore()
    await authStore.loginAction('admin', 'password')

    expect(authStore.token).toBe('test-token')
    expect(authStore.isLoggedIn).toBe(true)
    expect(localStorage.getItem('token')).toBe('test-token')
  })

  it('should check permissions correctly', async () => {
    const authStore = useAuthStore()
    await authStore.loginAction('admin', 'password')

    expect(authStore.hasPermission('system:user:list')).toBe(true)
  })

  it('should check roles correctly', async () => {
    const authStore = useAuthStore()
    await authStore.loginAction('admin', 'password')

    expect(authStore.hasRole('admin')).toBe(true)
    expect(authStore.hasRole('user')).toBe(false)
  })

  it('should logout successfully', async () => {
    const authStore = useAuthStore()
    await authStore.loginAction('admin', 'password')
    expect(authStore.isLoggedIn).toBe(true)

    await authStore.logoutAction()
    expect(authStore.token).toBe('')
    expect(authStore.isLoggedIn).toBe(false)
    expect(localStorage.getItem('token')).toBeNull()
  })
})
