import { test, expect } from '@playwright/test'

const BASE_URL = 'http://localhost:5176'
const API_URL = 'http://localhost:9080/api/v1'

test.describe('MOM3.0 M02 主数据模块测试', () => {
  let token: string

  test.beforeAll(async ({ request }) => {
    // Login
    const loginRes = await request.post(`${API_URL}/auth/login`, {
      data: { username: 'admin', password: 'admin123' }
    })
    const loginData = await loginRes.json()
    token = loginData.data.access_token
  })

  test('1. 班次管理 CRUD 测试', async ({ request }) => {
    // Create
    const createRes = await request.post(`${API_URL}/mdm/mdm-shift`, {
      headers: { Authorization: `Bearer ${token}` },
      data: {
        shift_code: 'TEST_SHIFT_PW_001',
        shift_name: '测试班次',
        start_time: '08:00',
        end_time: '17:00',
        work_hours: 8,
        is_night: 0
      }
    })
    expect(createRes.ok()).toBeTruthy()
    const createData = await createRes.json()
    const shiftId = createData.data?.id

    // List
    const listRes = await request.get(`${API_URL}/mdm/mdm-shift/list`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    expect(listRes.ok()).toBeTruthy()
    const listData = await listRes.json()
    expect(listData.data.list.length).toBeGreaterThan(0)

    // Update
    const updateRes = await request.put(`${API_URL}/mdm/mdm-shift/${shiftId}`, {
      headers: { Authorization: `Bearer ${token}` },
      data: { shift_name: '测试班次-已更新' }
    })
    expect(updateRes.ok()).toBeTruthy()

    // Delete
    const deleteRes = await request.delete(`${API_URL}/mdm/mdm-shift/${shiftId}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    expect(deleteRes.ok()).toBeTruthy()
  })

  test('2. 工序管理 CRUD 测试', async ({ request }) => {
    // Create
    const createRes = await request.post(`${API_URL}/mdm/operation`, {
      headers: { Authorization: `Bearer ${token}` },
      data: {
        operation_code: 'TEST_OP_PW_001',
        operation_name: '测试工序',
        standard_worktime: 60,
        is_key_process: 1,
        is_qc_point: 1
      }
    })
    expect(createRes.ok()).toBeTruthy()
    const createData = await createRes.json()
    const opId = createData.data?.id

    // List
    const listRes = await request.get(`${API_URL}/mdm/operation/list`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    expect(listRes.ok()).toBeTruthy()
    const listData = await listRes.json()
    expect(listData.data.list.length).toBeGreaterThan(0)

    // Update
    const updateRes = await request.put(`${API_URL}/mdm/operation/${opId}`, {
      headers: { Authorization: `Bearer ${token}` },
      data: { operation_name: '测试工序-已更新' }
    })
    expect(updateRes.ok()).toBeTruthy()

    // Delete
    const deleteRes = await request.delete(`${API_URL}/mdm/operation/${opId}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    expect(deleteRes.ok()).toBeTruthy()
  })

  test('3. BOM管理 CRUD 测试', async ({ request }) => {
    // Create BOM
    const createRes = await request.post(`${API_URL}/mdm/bom`, {
      headers: { Authorization: `Bearer ${token}` },
      data: {
        bom_code: 'TEST_BOM_PW_001',
        bom_name: '测试BOM',
        material_id: 1,
        material_code: 'MAT001',
        material_name: '测试物料',
        version: 'V1',
        status: 'DRAFT',
        items: [
          {
            material_id: 2,
            material_code: 'MAT002',
            material_name: '子物料',
            quantity: 2,
            unit: 'PCS',
            scrap_rate: 0
          }
        ]
      }
    })
    expect(createRes.ok()).toBeTruthy()

    // List
    const listRes = await request.get(`${API_URL}/mdm/bom/list`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    expect(listRes.ok()).toBeTruthy()
    const listData = await listRes.json()
    expect(listData.data.list.length).toBeGreaterThan(0)

    // Get BOM with items
    const bomId = listData.data.list[0].id
    const getRes = await request.get(`${API_URL}/mdm/bom/${bomId}/items`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    expect(getRes.ok()).toBeTruthy()

    // Delete
    const deleteRes = await request.delete(`${API_URL}/mdm/bom/${bomId}`, {
      headers: { Authorization: `Bearer ${token}` }
    })
    expect(deleteRes.ok()).toBeTruthy()
  })

  test('4. 前端页面访问测试', async ({ page }) => {
    // Test frontend is accessible
    const response = await page.goto(BASE_URL)
    expect(response?.status()).toBe(200)

    // Check for console errors on frontend
    const errors: string[] = []
    page.on('console', msg => {
      if (msg.type() === 'error' && !msg.text().includes('favicon')) {
        errors.push(msg.text())
      }
    })
    await page.waitForTimeout(2000)
    // Allow some non-critical errors
    expect(errors.filter(e => !e.includes('net::'))).toHaveLength(0)
  })
})
