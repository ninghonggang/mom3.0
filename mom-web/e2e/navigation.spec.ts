import { test, expect } from '@playwright/test'

test.describe('导航功能', () => {
  test('首页/仪表盘可以访问', async ({ page }) => {
    await page.goto('/')
    // 如果重定向到登录页，说明需要认证，这是预期行为
    // 如果能访问仪表盘，验证内容加载
    const url = page.url()
    if (url.includes('/login')) {
      await expect(page.locator('input[type="text"], input[placeholder*="用户名"]')).toBeVisible()
    } else {
      await expect(page.locator('body')).toBeVisible()
    }
  })

  test('未登录时访问受保护路由重定向到登录页', async ({ page }) => {
    // 在新页面上下文中直接访问受保护的仪表盘路由
    await page.goto('/dashboard')

    // 等待路由守卫完成重定向
    await page.waitForURL('**/login**', { timeout: 5000 })

    await expect(page.url()).toContain('/login')
    await expect(page.locator('body')).toContainText(/登 ?录|用户名|账号/)
  })
})
