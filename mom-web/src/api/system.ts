import request from '@/utils/request'

// 用户相关
export const getUserList = (params?: any) => {
  return request.get('/system/user/list', { params })
}

export const getUserById = (id: number) => {
  return request.get(`/system/user/${id}`)
}

export const createUser = (data: any) => {
  return request.post('/system/user', data)
}

export const updateUser = (id: number, data: any) => {
  return request.put(`/system/user/${id}`, data)
}

export const deleteUser = (id: number) => {
  return request.delete(`/system/user/${id}`)
}

export const resetUserPassword = (id: number, password: string) => {
  return request.put(`/system/user/${id}/password`, { password })
}

export const getAllRoles = () => {
  return request.get('/system/role/list')
}

export const assignUserRoles = (id: number, roleIds: number[]) => {
  return request.put(`/system/user/${id}/roles`, { role_ids: roleIds })
}

// 角色相关
export const getRoleList = (params?: any) => {
  return request.get('/system/role/list', { params })
}

export const getRoleById = (id: number) => {
  return request.get(`/system/role/${id}`)
}

export const createRole = (data: any) => {
  return request.post('/system/role', data)
}

export const updateRole = (id: number, data: any) => {
  return request.put(`/system/role/${id}`, data)
}

export const deleteRole = (id: number) => {
  return request.delete(`/system/role/${id}`)
}

export const getRoleMenus = (id: number) => {
  return request.get(`/system/role/${id}/menus`)
}

export const assignRoleMenus = (id: number, menuIds: number[]) => {
  return request.put(`/system/role/${id}/menus`, { menu_ids: menuIds })
}

export const getRolePerms = (id: number) => {
  return request.get(`/system/role/${id}/perms`)
}

export const assignRolePerms = (id: number, perms: string[]) => {
  return request.put(`/system/role/${id}/perms`, { perms })
}

// 菜单相关
export const getMenuList = () => {
  return request.get('/system/menu/list')
}

export const getMenuTree = () => {
  return request.get('/system/menu/tree')
}

export const getMenuById = (id: number) => {
  return request.get(`/system/menu/${id}`)
}

export const createMenu = (data: any) => {
  return request.post('/system/menu', data)
}

export const updateMenu = (id: number, data: any) => {
  return request.put(`/system/menu/${id}`, data)
}

export const deleteMenu = (id: number) => {
  return request.delete(`/system/menu/${id}`)
}

// 部门相关
export const getDeptList = () => {
  return request.get('/system/dept/list')
}

export const getDeptTree = () => {
  return request.get('/system/dept/tree')
}

export const getDeptById = (id: number) => {
  return request.get(`/system/dept/${id}`)
}

export const createDept = (data: any) => {
  return request.post('/system/dept', data)
}

export const updateDept = (id: number, data: any) => {
  return request.put(`/system/dept/${id}`, data)
}

export const deleteDept = (id: number) => {
  return request.delete(`/system/dept/${id}`)
}

// 字典相关
export const getDictTypeList = (params?: any) => {
  return request.get('/system/dict/type/list', { params })
}

export const getDictDataList = (dictType: string) => {
  return request.get(`/system/dict/${dictType}/data`)
}

export const createDictType = (data: any) => {
  return request.post('/system/dict/type', data)
}

export const updateDictType = (id: number, data: any) => {
  return request.put(`/system/dict/type/${id}`, data)
}

export const deleteDictType = (id: number) => {
  return request.delete(`/system/dict/type/${id}`)
}

// 岗位相关
export const getPostList = (params?: any) => {
  return request.get('/system/post/list', { params })
}

export const createPost = (data: any) => {
  return request.post('/system/post', data)
}

export const updatePost = (id: number, data: any) => {
  return request.put(`/system/post/${id}`, data)
}

export const deletePost = (id: number) => {
  return request.delete(`/system/post/${id}`)
}

// 租户相关
export const getTenantList = (params?: any) => {
  return request.get('/system/tenant/list', { params })
}

export const getTenantById = (id: number) => {
  return request.get(`/system/tenant/${id}`)
}

export const createTenant = (data: any) => {
  return request.post('/system/tenant', data)
}

export const updateTenant = (id: number, data: any) => {
  return request.put(`/system/tenant/${id}`, data)
}

export const deleteTenant = (id: number) => {
  return request.delete(`/system/tenant/${id}`)
}

// 登录日志相关
export const getLoginLogList = (params?: any) => {
  return request.get('/system/loginlog/list', { params })
}

export const cleanLoginLog = () => {
  return request.delete('/system/loginlog/clean')
}

export const exportLoginLog = () => {
  return request.get('/system/loginlog/export')
}

// 操作日志相关
export const getOperLogList = (params?: any) => {
  return request.get('/system/operlog/list', { params })
}

export const cleanOperLog = () => {
  return request.delete('/system/operlog/clean')
}

export const exportOperLog = () => {
  return request.get('/system/operlog/export')
}
