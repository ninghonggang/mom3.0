import { test, expect } from '@playwright/test';

test('comprehensive menu and page exploration', async ({ page }) => {
  // Login using relative URL
  await page.goto('login');
  await page.waitForLoadState('networkidle');

  await page.fill('input[type="text"]', 'admin');
  await page.fill('input[type="password"]', 'admin123');
  await page.click('button:has-text("登 录")');

  // Wait for navigation after login
  await page.waitForURL(/\/dashboard/, { timeout: 15000 }).catch(() => {
    console.log('Current URL after login:', page.url());
  });
  await page.waitForTimeout(3000);

  // Screenshot dashboard
  await page.screenshot({ path: 'D:/tmp/mom3_dashboard.png', fullPage: false });
  console.log('Dashboard截图: D:/tmp/mom3_dashboard.png');

  // Get sidebar menu
  const sidebarItems = await page.evaluate(() => {
    const items = document.querySelectorAll('.sidebar .el-menu-item');
    return Array.from(items).map(item => item.textContent?.trim() || '');
  });

  console.log('\n=== 侧边栏菜单 (共' + sidebarItems.length + '项) ===');
  sidebarItems.forEach((item, i) => console.log(`${i + 1}. ${item}`));

  // Test comprehensive routes
  console.log('\n=== 路由测试 ===');

  const allRoutes = [
    '/dashboard',
    '/system/user', '/system/role', '/system/menu', '/system/dept',
    '/mdm/material', '/mdm/material-category',
    '/production/order', '/production/sales-order',
    '/equipment', '/equipment/check',
    '/wms/warehouse', '/wms/location', '/wms/inventory',
    '/quality/iqc', '/quality/ipqc',
    '/bpm/process', '/bpm/task',
    '/mes/sop-pdf', '/mes/team',
    '/scp/purchase', '/scp/rfq',
    '/fin/payment-request',
    '/report/production-daily'
  ];

  const results: Record<string, any> = {};

  for (const route of allRoutes) {
    try {
      await page.goto(route);
      await page.waitForLoadState('networkidle');
      await page.waitForTimeout(1500);

      // Check if redirected to login
      if (page.url().includes('/login')) {
        results[route] = { status: 'login-redirect' };
        console.log(`⟲ ${route}: 重定向到登录`);
        continue;
      }

      // Get page info
      const buttons = await page.locator('button').allTextContents();
      const tables = await page.locator('.el-table').count();
      const forms = await page.locator('.el-form').count();

      // Find main content title
      const title = await page.locator('h1, h2, h3, .page-title, [class*="title"], .content-header').first().textContent().catch(() => 'N/A');

      results[route] = { status: 'OK', title: title?.trim(), buttons: buttons.length, tables };

      const btnSummary = buttons.slice(0, 5).map(b => b.trim()).filter(Boolean).join(', ');
      console.log(`✓ ${route}: "${title?.trim()}" | 按钮:${buttons.length} | 表格:${tables}`);

      // Take screenshot for key pages
      if (['/system/user', '/mdm/material', '/wms/warehouse', '/bpm/process'].includes(route)) {
        await page.screenshot({ path: `D:/tmp/mom3${route.replace(/\//g, '_')}.png`, fullPage: false });
      }
    } catch (e) {
      results[route] = { status: 'error', message: e.message };
      console.log(`✗ ${route}: ${e.message}`);
    }
  }

  console.log('\n=== 汇总 ===');
  const okCount = Object.values(results).filter(r => r.status === 'OK').length;
  const failCount = Object.values(results).filter(r => r.status !== 'OK').length;
  console.log(`成功: ${okCount}, 失败/重定向: ${failCount}`);

  await page.screenshot({ path: 'D:/tmp/mom3_final.png', fullPage: false });
});
