import { test, expect } from '@playwright/test'

test.describe('L1 登录与基础功能', () => {
  test.beforeEach(async ({ page }) => {
    // 基础设置
  })

  // 1. 登录页面正确加载
  test('登录页面正确加载', async ({ page }) => {
    await page.goto('/login')
    await expect(page.locator('h1, h2, .login-title, [class*="title"]').first()).toBeVisible()
    await expect(page.locator('input[type="text"], input[placeholder*="用户名"], input[placeholder*="账号"]').first()).toBeVisible()
    await expect(page.locator('input[type="password"]').first()).toBeVisible()
    await expect(page.locator('button[type="submit"], button:has-text("登录"), button:has-text("登 录")').first()).toBeVisible()
  })

  // 2. 输入用户名和密码
  test('输入用户名和密码', async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[type="text"], input[placeholder*="用户名"], input[placeholder*="账号"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await expect(page.locator('input[type="text"], input[placeholder*="用户名"]')).toHaveValue('admin')
    await expect(page.locator('input[type="password"]')).toHaveValue('admin123')
  })

  // 3. 登录成功并跳转仪表盘
  test('登录成功并跳转仪表盘', async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[type="text"], input[placeholder*="用户名"], input[placeholder*="账号"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"], button:has-text("登录"), button:has-text("登 录")')
    await page.waitForURL('**/dashboard', { timeout: 10000 })
    await expect(page.url()).toContain('/dashboard')
  })

  // 4. 未登录访问受保护路由重定向
  test('未登录访问受保护路由重定向', async ({ page }) => {
    await page.goto('/dashboard')
    await page.waitForURL('**/login**', { timeout: 5000 })
    await expect(page.url()).toContain('/login')
  })

  // 5. 控制台无关键错误
  test('控制台无关键错误', async ({ page }) => {
    const errors: string[] = []
    page.on('console', msg => {
      if (msg.type() === 'error') {
        errors.push(msg.text())
      }
    })
    await page.goto('/login')
    await page.waitForTimeout(2000)
    const criticalErrors = errors.filter(e =>
      !e.includes('favicon') &&
      !e.includes('WebSocket') &&
      !e.includes('ws://') &&
      !e.includes('net::ERR')
    )
    expect(criticalErrors).toHaveLength(0)
  })
})
