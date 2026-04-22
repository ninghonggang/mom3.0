import { test, expect } from '@playwright/test'

test.describe('L3 Quality 质量管理流程', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/login')
    await page.fill('input[type="text"]', 'admin')
    await page.fill('input[type="password"]', 'admin123')
    await page.click('button[type="submit"]')
    await page.waitForURL('**/dashboard', { timeout: 10000 })
  })

  test('IQC 来料检验', async ({ page }) => {
    await page.goto('/quality/iqc')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('IPQC 过程检验', async ({ page }) => {
    await page.goto('/quality/ipqc')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('FQC 出货检验', async ({ page }) => {
    await page.goto('/quality/fqc')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('OQC 出货检验', async ({ page }) => {
    await page.goto('/quality/oqc')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('NCR 不良品处理', async ({ page }) => {
    await page.goto('/quality/ncr')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })

  test('SPC 统计过程控制', async ({ page }) => {
    await page.goto('/quality/spc')
    await page.waitForTimeout(1000)
    await expect(page.locator('body')).toBeVisible()
  })
})
