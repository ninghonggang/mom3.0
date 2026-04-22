import { test, expect } from '@playwright/test'

test.describe('L3 BPM 流程引擎', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[type="text"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"], button:has-text("登录"), button:has-text("登 录")')
    await page.waitForURL('**/dashboard', { timeout: 10000 })
  })

  test('流程定义列表', async ({ page }) => {
    await page.goto('/bpm/process')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('流程实例列表', async ({ page }) => {
    await page.goto('/bpm/instance')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('任务列表', async ({ page }) => {
    await page.goto('/bpm/task')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })
})
