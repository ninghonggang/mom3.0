import request from '@/utils/request'

export interface LoginParams {
  username: string
  password: string
}

export interface LoginResponse {
  access_token: string
  refresh_token: string
  expires_in: number
  user: {
    id: number
    username: string
    nickname: string
    avatar?: string
    email?: string
    phone?: string
    roles: string[]
    perms: string[]
  }
}

export const login = (data: LoginParams) => {
  return request.post<any, any>('/auth/login', data)
}

export const logout = () => {
  return request.post('/auth/logout')
}

export const refreshToken = (refresh_token: string) => {
  return request.post('/auth/refresh', { refresh_token })
}

export const getUserInfo = () => {
  return request.get('/auth/info')
}

export const changePassword = (data: { old_password: string; new_password: string }) => {
  return request.put('/auth/password', data)
}
