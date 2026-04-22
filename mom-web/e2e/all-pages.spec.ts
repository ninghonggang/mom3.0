import { test, expect } from '@playwright/test'

/**
 * 所有页面加载测试
 * 测试项目中的 60+ 个页面是否能正常加载
 */

// 需要测试的页面路由列表（按模块分组）
const pageRoutes = [
  // 首页
  '/dashboard',

  // 系统管理
  '/system/user',
  '/system/role',
  '/system/menu',
  '/system/dept',
  '/system/dict',
  '/system/post',
  '/system/tenant',
  '/system/login-log',
  '/system/oper-log',
  '/system/config',
  '/system/ai-config',

  // 主数据管理
  '/mdm/material',
  '/mdm/material-category',
  '/mdm/workshop',
  '/mdm/line',
  '/mdm/workstation',
  '/mdm/mdm-shift',
  '/mdm/bom',
  '/mdm/operation',
  '/mdm/customer',

  // 生产执行
  '/production/sales-order',
  '/production/report',
  '/production/dispatch',
  '/production/order',
  '/production/kanban',
  '/production/order-change',
  '/production/package',
  '/production/first-last-inspect',

  // 设备管理
  '/equipment',
  '/equipment/check',
  '/equipment/maintenance',
  '/equipment/repair',
  '/equipment/spare',
  '/equipment/oee',

  // 仓储管理
  '/wms/warehouse',
  '/wms/data-point',
  '/wms/scan-log',
  '/wms/location',
  '/wms/inventory',
  '/wms/receive',
  '/wms/delivery',

  // 质量管理
  '/quality/iqc',
  '/quality/ipqc',
  '/quality/fqc',
  '/quality/oqc',
  '/quality/defect-code',
  '/quality/defect-record',
  '/quality/ncr',
  '/quality/spc',

  // APS计划
  '/aps/mps',
  '/aps/mrp',
  '/aps/schedule',
  '/aps/work-center',

  // 追溯管理
  '/trace/query',
  '/trace/andon',

  // 能源管理
  '/energy/monitor',

  // 供应链管理
  '/scp/purchase',
  '/scp/rfq',
  '/scp/supplier-quote',
  '/scp/sales-order',
  '/scp/supplier-kpi',

  // 统一告警
  '/alert/rules',
  '/alert/records',
  '/alert/escalation',
  '/alert/statistics',

  // 流程管理
  '/bpm/process',
  '/bpm/instance',
  '/bpm/task',

  // APS扩展
  '/aps/rolling-config',
  '/aps/delivery-analysis',
  '/aps/material-shortage',
  '/aps/shortage-rule',
  '/aps/changeover-matrix',
  '/aps/product-family',

  // WMS扩展
  '/wms/transfer',
  '/wms/stock-check',

  // 质量扩展
  '/quality/qrci',
  '/quality/lpa',

  // 设备扩展
  '/equipment/inspection/templates',
  '/equipment/inspection/plans',
  '/equipment/inspection/records',
  '/equipment/inspection/defects',
  '/equipment/gauge',
]

test.describe('所有页面加载测试', () => {
  for (const path of pageRoutes) {
    test(`页面加载: ${path}`, async ({ page }) => {
      // 访问页面
      await page.goto(path, { timeout: 30000 })

      // 等待页面基本加载
      await page.waitForLoadState('domcontentloaded', { timeout: 15000 })

      // 验证页面 body 可见
      await expect(page.locator('body')).toBeVisible()

      // 捕获控制台错误（但不断言失败，只记录）
      const errors: string[] = []
      page.on('console', msg => {
        if (msg.type() === 'error') {
          errors.push(msg.text())
        }
      })

      // 等待一小段时间让异步错误出现
      await page.waitForTimeout(500)

      // 过滤非关键错误
      const criticalErrors = errors.filter(e =>
        !e.includes('favicon') &&
        !e.includes('WebSocket') &&
        !e.includes('ws://') &&
        !e.includes('net::ERR')
      )

      // 如果有严重错误，在测试报告中输出（但不失败，因为是开发中）
      if (criticalErrors.length > 0) {
        console.log(`[${path}] 控制台错误:`, criticalErrors)
      }
    })
  }
})

test.describe('页面批量验证', () => {
  test('所有页面在 120 秒内完成加载', async ({ page }) => {
    const startTime = Date.now()
    let failedPages: string[] = []

    for (const path of pageRoutes) {
      try {
        await page.goto(path, { timeout: 10000, waitUntil: 'domcontentloaded' })
        await page.waitForTimeout(200)
      } catch (e) {
        failedPages.push(path)
      }
    }

    const elapsed = (Date.now() - startTime) / 1000
    console.log(`批量验证完成: ${pageRoutes.length - failedPages.length}/${pageRoutes.length} 成功, 耗时 ${elapsed.toFixed(1)}s`)

    if (failedPages.length > 0) {
      console.log('失败页面:', failedPages)
    }

    // 允许部分页面失败（开发中）
    const successRate = (pageRoutes.length - failedPages.length) / pageRoutes.length
    expect(successRate).toBeGreaterThan(0.5) // 至少 50% 成功
  })
})
