import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
import router from '@/router'

// 创建axios实例
const service: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api/v1',
  timeout: 30000
})

// 请求队列
let isRefreshing = false
let requestQueue: ((token: string) => void)[] = []

// 请求拦截器
service.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    // 租户ID
    const tenantId = localStorage.getItem('tenantId')
    if (tenantId) {
      config.headers['X-Tenant-ID'] = tenantId
    }

    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse) => {
    const { code, message, data } = response.data

    // 成功
    if (code === 200) {
      return response.data
    }

    // Token过期
    if (code === 40101) {
      handleUnauthorized()
      return Promise.reject(new Error(message || '未授权'))
    }

    // 其他错误
    ElMessage.error(message || '请求失败')
    return Promise.reject(new Error(message))
  },
  (error) => {
    // 网络错误
    if (!error.response) {
      ElMessage.error('网络错误，请检查服务器')
      return Promise.reject(error)
    }

    const { status, data } = error.response

    if (status === 401) {
      handleUnauthorized()
      return Promise.reject(new Error('未授权'))
    }

    if (status === 403) {
      ElMessage.error('没有权限')
      return Promise.reject(new Error('没有权限'))
    }

    if (status === 404) {
      ElMessage.error('请求资源不存在')
      return Promise.reject(new Error('请求资源不存在'))
    }

    if (status >= 500) {
      ElMessage.error('服务器错误')
      return Promise.reject(new Error('服务器错误'))
    }

    ElMessage.error(data?.message || '请求失败')
    return Promise.reject(error)
  }
)

// 处理未授权
function handleUnauthorized() {
  if (!isRefreshing) {
    isRefreshing = true
    ElMessageBox.confirm('登录已过期，请重新登录', '提示', {
      confirmButtonText: '重新登录',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(() => {
        localStorage.removeItem('token')
        localStorage.removeItem('refresh_token')
        router.push('/login')
      })
      .finally(() => {
        isRefreshing = false
        requestQueue = []
      })
  }
}

export default service

// 扩展AxiosRequestConfig类型
declare module 'axios' {
  interface AxiosRequestConfig {
    showLoading?: boolean
  }
}
