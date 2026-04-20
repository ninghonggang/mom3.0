import request from '@/utils/request'

// ==================== 告警规则 ====================
export const getAlertRulesList = (params?: any) => request.get('/alert/rule/list', { params })
export const getAlertRule = (id: number) => request.get(`/alert/rule/${id}`)
export const createAlertRule = (data: any) => request.post('/alert/rule', data)
export const updateAlertRule = (id: number, data: any) => request.put(`/alert/rule/${id}`, data)
export const deleteAlertRule = (id: number) => request.delete(`/alert/rule/${id}`)
export const enableAlertRule = (id: number) => request.post(`/alert/rule/${id}/enable`)
export const disableAlertRule = (id: number) => request.post(`/alert/rule/${id}/disable`)

// ==================== 告警记录 ====================
export const getAlertRecordsList = (params?: any) => request.get('/alert/record/list', { params })
export const getAlertRecord = (id: number) => request.get(`/alert/record/${id}`)
export const acknowledgeAlertRecord = (id: number, data?: any) => request.post(`/alert/record/${id}/ack`, data)
export const resolveAlertRecord = (id: number, data?: any) => request.post(`/alert/record/${id}/resolve`, data)
export const closeAlertRecord = (id: number) => request.post(`/alert/record/${id}/close`)

// ==================== 告警统计 ====================
export const getAlertStatistics = () => request.get('/alert/statistics')

// ==================== 升级规则 ====================
export const getAlertEscalationList = (params?: any) => request.get('/alert/escalation/list', { params })
export const getAlertEscalation = (id: number) => request.get(`/alert/escalation/${id}`)
export const createAlertEscalation = (data: any) => request.post('/alert/escalation', data)
export const updateAlertEscalation = (id: number, data: any) => request.put(`/alert/escalation/${id}`, data)
export const deleteAlertEscalation = (id: number) => request.delete(`/alert/escalation/${id}`)

// ==================== 通知渠道配置 ====================
export const getNotificationChannels = (params?: any) => request.get('/alert/channel/list', { params })
export const getNotificationChannel = (id: number) => request.get(`/alert/channel/${id}`)
export const createNotificationChannel = (data: any) => request.post('/alert/channel', data)
export const updateNotificationChannel = (id: number, data: any) => request.put(`/alert/channel/${id}`, data)
export const deleteNotificationChannel = (id: number) => request.delete(`/alert/channel/${id}`)
export const enableNotificationChannel = (id: number) => request.post(`/alert/channel/${id}/enable`)
export const disableNotificationChannel = (id: number) => request.post(`/alert/channel/${id}/disable`)

// ==================== 通知记录 ====================
export const getNotificationList = (params?: any) => request.get('/alert/notify/logs', { params })
export const getNotification = (id: number) => request.get(`/alert/notify/${id}`)
export const getNotificationStatistics = () => request.get('/alert/notify/statistics')
export const markNotificationAsRead = (id: number) => request.put(`/alert/notify/${id}/read`)

// ==================== 发送通知 ====================
export const sendNotification = (data: any) => request.post('/alert/send', data)
export const sendNotificationToChannel = (data: any) => request.post('/alert/channel/send', data)
