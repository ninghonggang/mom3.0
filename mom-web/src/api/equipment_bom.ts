import request from '@/utils/request'

// 设备BOM管理
export const getEquipmentBomList = (params?: any) => {
  return request.get('/equipment/bom/page', { params })
}

export const getEquipmentBomByEquipment = (equipmentId: number) => {
  return request.get('/equipment/bom/list', { params: { equipment_id: equipmentId } })
}

export const getEquipmentBom = (id: number) => {
  return request.get(`/equipment/bom/${id}`)
}

export const createEquipmentBom = (data: any) => {
  return request.post('/equipment/bom', data)
}

export const updateEquipmentBom = (id: number, data: any) => {
  return request.put(`/equipment/bom/${id}`, data)
}

export const deleteEquipmentBom = (id: number) => {
  return request.delete(`/equipment/bom/${id}`)
}