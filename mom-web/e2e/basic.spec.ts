import { test, expect } from '@playwright/test'

test.describe('基础功能', () => {
  test('页面标题正确', async ({ page }) => {
    await page.goto('/login')
    const title = await page.title()
    expect(title.length).toBeGreaterThan(0)
  })

  test('没有控制台错误', async ({ page }) => {
    const errors: string[] = []
    page.on('console', msg => {
      if (msg.type() === 'error') {
        errors.push(msg.text())
      }
    })

    await page.goto('/login')
    await page.waitForTimeout(1000)

    // 过滤掉常见的非关键错误（favicon、WebSocket 后端未运行）
    const criticalErrors = errors.filter(e =>
      !e.includes('favicon') &&
      !e.includes('WebSocket') &&
      !e.includes('ws://')
    )
    expect(criticalErrors).toHaveLength(0)
  })

  test('响应式布局正常', async ({ page }) => {
    await page.goto('/login')

    // 测试桌面视图
    await page.setViewportSize({ width: 1920, height: 1080 })
    await expect(page.locator('body')).toBeVisible()

    // 测试平板视图
    await page.setViewportSize({ width: 768, height: 1024 })
    await expect(page.locator('body')).toBeVisible()

    // 测试手机视图
    await page.setViewportSize({ width: 375, height: 667 })
    await expect(page.locator('body')).toBeVisible()
  })
})
