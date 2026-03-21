import { test, expect } from '@playwright/test'

test.describe('登录功能', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
  })

  test('登录页面正确加载', async ({ page }) => {
    await expect(page.locator('h1, h2, .login-title, [class*="title"]').first()).toBeVisible()
    await expect(page.locator('input[type="text"], input[placeholder*="用户名"], input[placeholder*="账号"]').first()).toBeVisible()
    await expect(page.locator('input[type="password"]').first()).toBeVisible()
    await expect(page.locator('button[type="submit"], button:has-text("登录"), button:has-text("登 录")').first()).toBeVisible()
  })

  test('可以输入用户名和密码', async ({ page }) => {
    const usernameInput = page.locator('input[type="text"], input[placeholder*="用户名"], input[placeholder*="账号"]').first()
    const passwordInput = page.locator('input[type="password"]').first()

    await usernameInput.fill('admin')
    await passwordInput.fill('admin123')

    await expect(usernameInput).toHaveValue('admin')
    await expect(passwordInput).toHaveValue('admin123')
  })

  test('提交按钮可用', async ({ page }) => {
    const submitBtn = page.locator('button[type="submit"], button:has-text("登录"), button:has-text("登 录")').first()
    await expect(submitBtn).toBeEnabled()
  })
})
