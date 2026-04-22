import { test, expect } from '@playwright/test'

test.describe('L2/L3 APS 计划排程', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.waitForLoadState('networkidle')
    await page.fill('input[type="text"], input[placeholder*="用户名"], input[placeholder*="账号"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"], button:has-text("登录"), button:has-text("登 录")')
    await page.waitForURL('**/dashboard', { timeout: 15000 })
  })

  test('MPS 主生产计划', async ({ page }) => {
    await page.goto('/aps/mps')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('MRP 物料需求计划', async ({ page }) => {
    await page.goto('/aps/mrp')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('排程甘特图', async ({ page }) => {
    await page.goto('/aps/schedule')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('缺料分析', async ({ page }) => {
    await page.goto('/aps/material-shortage')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('滚动排程', async ({ page }) => {
    await page.goto('/aps/rolling-config')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })
})
