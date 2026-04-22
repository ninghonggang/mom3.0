import { test, expect } from '@playwright/test'

// 完整路由列表
const allRoutes = [
  // 首页
  '/dashboard',

  // 系统管理
  '/system/user', '/system/role', '/system/menu', '/system/dept',
  '/system/dict', '/system/post', '/system/tenant',
  '/system/login-log', '/system/oper-log', '/system/config',

  // MDM
  '/mdm/material', '/mdm/material-category', '/mdm/workshop',
  '/mdm/line', '/mdm/workstation', '/mdm/customer', '/mdm/supplier',

  // 生产执行
  '/production/sales-order', '/production/report', '/production/dispatch',
  '/production/order', '/production/issue', '/production/return',

  // 设备
  '/equipment', '/equipment/check', '/equipment/maintenance',
  '/equipment/repair', '/equipment/spare',

  // WMS
  '/wms/warehouse', '/wms/location', '/wms/inventory',
  '/wms/receive', '/wms/delivery', '/wms/transfer', '/wms/stock-check',

  // 质量
  '/quality/iqc', '/quality/ipqc', '/quality/fqc', '/quality/oqc',
  '/quality/ncr', '/quality/spc',

  // APS
  '/aps/mps', '/aps/mrp', '/aps/schedule',
  '/aps/rolling-config', '/aps/material-shortage',

  // SCP
  '/scp/purchase', '/scp/rfq', '/scp/supplier-quote',
  '/scp/sales-order', '/scp/supplier-kpi', '/scp/customer-inquiry',

  // BPM
  '/bpm/process', '/bpm/instance', '/bpm/task',

  // EAM
  '/eam/factory', '/eam/equipment-org', '/eam/downtime',

  // Alert
  '/alert/rules', '/alert/records', '/alert/escalation',

  // Report
  '/report/production-daily', '/report/quality-weekly', '/report/oee',

  // Fin
  '/fin/payment-request', '/fin/purchase-settlement', '/fin/sales-settlement',

  // Integration
  '/integration/interface-config', '/integration/execution-log',

  // AGV
  '/agv/task', '/agv/device', '/agv/location',

  // MES
  '/mes/team', '/mes/process-routes', '/mes/offline', '/mes/person-skill',
]

test.describe('L4 全量冒烟测试', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[type="text"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"]')
    await page.waitForURL('**/dashboard', { timeout: 10000 })
  })

  for (const route of allRoutes) {
    test(`页面加载: ${route}`, async ({ page }) => {
      const errors: string[] = []
      page.on('console', msg => {
        if (msg.type() === 'error' && msg.text().includes('500')) {
          errors.push(msg.text())
        }
      })

      await page.goto(route, { waitUntil: 'networkidle' })
      await page.waitForTimeout(1000)

      const criticalErrors = errors.filter(e => e.includes('500'))
      expect(criticalErrors).toHaveLength(0)
    })
  }
})
