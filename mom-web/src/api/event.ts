import request from '@/utils/request'

export interface EventLog {
  id: number
  event_type: string
  event_data: any
  created_at: string
}

export const eventApi = {
  list: () => {
    return request.get<any, any>('/event/logs')
  }
}
