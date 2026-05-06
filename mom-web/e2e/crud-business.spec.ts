/**
 * MOM3.0 CRUD & Business Flow Tests
 * 测试基础CRUD操作和业务流程
 */

const BASE_URL = 'http://localhost:9081/api/v1'
const WEB_URL = 'http://localhost:5175'

// 测试结果收集
const results = {
  passed: [],
  failed: []
}

function log(name, passed, msg = '') {
  const status = passed ? '✅ PASS' : '❌ FAIL'
  console.log(`${status}: ${name}${msg ? ' - ' + msg : ''}`)
  if (passed) results.passed.push(name)
  else results.failed.push({ name, msg })
}

async function request(method, path, body = null, token = null) {
  const opts = {
    method,
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { 'Authorization': `Bearer ${token}` } : {})
    }
  }
  if (body) opts.body = JSON.stringify(body)

  try {
    const res = await fetch(`${BASE_URL}${path}`, opts)
    const data = await res.json()
    return { status: res.status, data }
  } catch (e) {
    return { status: 0, error: e.message }
  }
}

async function login() {
  const res = await request('POST', '/auth/login', { username: 'admin', password: 'admin123' })
  if (res.status === 200 && res.data?.data?.access_token) {
    log('登录', true)
    return res.data.data.access_token
  }
  log('登录', false, `status: ${res.status}`)
  return null
}

async function runCRUDTests(token) {
  console.log('\n========== CRUD 测试 ==========')

  // 1. MDM - 物料管理
  console.log('\n--- MDM 物料管理 ---')
  let res = await request('POST', '/mdm/material', {
    material_code: `TEST-M-${Date.now()}`,
    material_name: '测试物料',
    material_type: '原材料',
    unit: 'PCS',
    spec: '10*10',
    status: 1
  }, token)
  log('创建物料', res.status === 200, `code: ${res.status}`)
  const materialId = res.data?.data?.id

  if (materialId) {
    res = await request('GET', `/mdm/material/${materialId}`, null, token)
    log('查询物料', res.status === 200)

    res = await request('PUT', `/mdm/material/${materialId}`, {
      material_name: '测试物料-已修改'
    }, token)
    log('修改物料', res.status === 200)
  }

  // 2. MDM - 供应商管理
  console.log('\n--- MDM 供应商管理 ---')
  res = await request('POST', '/mdm/supplier', {
    supplier_code: `TEST-S-${Date.now()}`,
    supplier_name: '测试供应商',
    contact_person: '张三',
    contact_phone: '13800138000',
    status: 1
  }, token)
  log('创建供应商', res.status === 200)
  const supplierId = res.data?.data?.id

  // 3. MDM - 客户管理
  console.log('\n--- MDM 客户管理 ---')
  res = await request('POST', '/mdm/customer', {
    customer_code: `TEST-C-${Date.now()}`,
    customer_name: '测试客户',
    contact_person: '李四',
    contact_phone: '13900139000',
    status: 1
  }, token)
  log('创建客户', res.status === 200)
  const customerId = res.data?.data?.id

  // 4. MDM - 工序管理
  console.log('\n--- MDM 工序管理 ---')
  res = await request('POST', '/mdm/operation', {
    operation_code: `OP-${Date.now()}`,
    operation_name: '测试工序',
    description: '测试工序描述',
    status: 1
  }, token)
  log('创建工序', res.status === 200)

  // 5. MDM - BOM管理
  console.log('\n--- MDM BOM管理 ---')
  res = await request('POST', '/mdm/bom', {
    bom_code: `BOM-${Date.now()}`,
    product_code: materialId || 'TEST',
    version: '1.0',
    status: 1
  }, token)
  log('创建BOM', res.status === 200)

  // 6. 仓库管理
  console.log('\n--- WMS 仓库管理 ---')
  res = await request('POST', '/wms/warehouse', {
    warehouse_code: `WH-${Date.now()}`,
    warehouse_name: '测试仓库',
    warehouse_type: '原材料仓',
    status: 1
  }, token)
  log('创建仓库', res.status === 200)
  const warehouseId = res.data?.data?.id

  // 7. 库位管理
  console.log('\n--- WMS 库位管理 ---')
  if (warehouseId) {
    res = await request('POST', '/wms/location', {
      location_code: `LOC-${Date.now()}`,
      location_name: '测试库位',
      warehouse_id: warehouseId,
      status: 1
    }, token)
    log('创建库位', res.status === 200)
  }

  // 8. 质量管理 - IQC
  console.log('\n--- Quality IQC ---')
  res = await request('POST', '/quality/iqc', {
    iqc_no: `IQC-${Date.now()}`,
    supplier_id: supplierId || 1,
    inspection_type: '常规',
    status: 1
  }, token)
  log('创建IQC', res.status === 200)

  return { materialId, supplierId, customerId, warehouseId }
}

async function runBusinessFlowTests(token) {
  console.log('\n========== 业务流程测试 ==========')

  // 1. SCM - 询价流程
  console.log('\n--- SCM 询价流程 ---')
  let res = await request('POST', '/scp/rfq', {
    rfq_no: `RFQ-${Date.now()}`,
    title: '测试询价单',
    supplier_id: 1,
    required_date: new Date().toISOString(),
    status: 'pending'
  }, token)
  log('创建询价单', res.status === 200)
  const rfqId = res.data?.data?.id

  if (rfqId) {
    // 询价回复通过 supplier-quotes API (直接POST)
    res = await request('POST', '/scp/supplier-quotes', {
      rfq_id: rfqId,
      supplier_id: 1,
      unit_price: 100,
      delivery_date: new Date().toISOString(),
      valid_until: new Date(Date.now() + 30*24*3600*1000).toISOString()
    }, token)
    log('添加询价回复', res.status === 200, `code: ${res.status}`)

    // Award通过 supplier-quotes 的 award 接口
    const quoteId = res.data?.data?.id
    if (quoteId) {
      res = await request('POST', `/scp/supplier-quotes/rfq/${rfqId}/award`, {
        quote_id: quoteId
      }, token)
      log('询价单Award', res.status === 200, `code: ${res.status}`)
    }
  }

  // 2. SCM - 采购订单
  console.log('\n--- SCM 采购订单 ---')
  res = await request('POST', '/scp/purchase-orders', {
    po_no: `PO-${Date.now()}`,
    title: '测试采购订单',
    supplier_id: 1,
    order_date: new Date().toISOString(),
    status: 'draft'
  }, token)
  log('创建采购订单', res.status === 200)
  const poId = res.data?.data?.id

  if (poId) {
    res = await request('POST', `/scp/purchase-orders/${poId}/submit`, {}, token)
    log('提交采购订单', res.status === 200, `code: ${res.status}`)

    // 入库通过采购订单的receive接口
    res = await request('POST', `/scp/purchase-orders/${poId}/receive`, {
      receive_no: `REC-${Date.now()}`,
      supplier_id: 1
    }, token)
    log('创建入库单', res.status === 200, `code: ${res.status}`)
  }

  // 4. 生产工单
  console.log('\n--- Production 工单 ---')
  res = await request('POST', '/production/order', {
    order_no: `WO-${Date.now()}`,
    product_code: 'TEST-PROD',
    quantity: 100,
    planned_start: new Date().toISOString(),
    planned_end: new Date(Date.now() + 7*24*3600*1000).toISOString(),
    status: 'pending'
  }, token)
  log('创建生产工单', res.status === 200)
  const woId = res.data?.data?.id

  if (woId) {
    res = await request('POST', `/production/order/${woId}/release`, {}, token)
    log('发布工单', res.status === 200, `code: ${res.status}`)
  }

  // 5. 销售订单
  console.log('\n--- SCP 销售订单 ---')
  res = await request('POST', '/scp/sales-orders', {
    so_no: `SO-${Date.now()}`,
    customer_id: 1,
    order_date: new Date().toISOString(),
    delivery_date: new Date(Date.now() + 14*24*3600*1000).toISOString(),
    status: 'draft'
  }, token)
  log('创建销售订单', res.status === 200)
  const soId = res.data?.data?.id

  if (soId) {
    res = await request('POST', `/scp/sales-orders/${soId}/submit`, {}, token)
    log('提交销售订单', res.status === 200, `code: ${res.status}`)
  }

  // 6. 设备管理
  console.log('\n--- EAM 设备 ---')
  res = await request('POST', '/equipment', {
    equipment_code: `EQ-${Date.now()}`,
    equipment_name: '测试设备',
    equipment_type: '加工中心',
    status: 1
  }, token)
  log('创建设备', res.status === 200)

  // 7. 报工记录
  console.log('\n--- MES 报工 ---')
  res = await request('POST', '/production/report', {
    work_order_id: woId || 1,
    operation_id: 1,
    quantity: 10,
    report_time: new Date().toISOString(),
    operator: 'admin'
  }, token)
  log('创建报工记录', res.status === 200, `code: ${res.status}`)
}

async function runListTests(token) {
  console.log('\n========== 列表查询测试 ==========')

  // 列表查询 - 注意API路径是复数形式
  const listApis = [
    '/mdm/material/list',
    '/mdm/supplier/list',
    '/mdm/customer/list',
    '/mdm/workshop/list',
    '/mdm/line/list',
    '/wms/warehouse/list',
    '/wms/location/list',
    '/wms/inventory/list',
    '/production/order/list',
    '/scp/purchase-orders/list',   // 复数
    '/scp/sales-orders/list',      // 复数
    '/scp/rfq/list',
    '/quality/iqc/list',
    '/equipment/list',
    '/mes/team/list'
  ]

  for (const api of listApis) {
    const res = await request('GET', api, null, token)
    const passed = res.status === 200
    log(`列表${api}`, passed, `code: ${res.status}`)
  }
}

async function run() {
  console.log('🚀 MOM3.0 CRUD & Business Flow Tests')
  console.log('====================================')

  const token = await login()
  if (!token) {
    console.log('❌ 无法获取登录token，测试终止')
    return
  }

  await runCRUDTests(token)
  await runBusinessFlowTests(token)
  await runListTests(token)

  console.log('\n========== 测试结果汇总 ==========')
  console.log(`✅ 通过: ${results.passed.length}`)
  console.log(`❌ 失败: ${results.failed.length}`)

  if (results.failed.length > 0) {
    console.log('\n失败详情:')
    results.failed.forEach(f => console.log(`  - ${f.name}: ${f.msg}`))
  }

  console.log('\n测试完成!')
}

run().catch(console.error)