import { test, expect } from '@playwright/test';

test('explore full menu structure', async ({ page }) => {
  // Navigate to app
  await page.goto('http://localhost:5176/');
  await page.waitForLoadState('networkidle');

  // Login
  await page.fill('input[type="text"]', 'admin');
  await page.fill('input[type="password"]', 'admin123');
  await page.click('button:has-text("登 录")');

  // Wait for dashboard
  await page.waitForURL('**/dashboard', { timeout: 15000 });
  await page.waitForTimeout(2000);

  // Screenshot dashboard
  await page.screenshot({ path: 'D:/tmp/mom3_dashboard.png', fullPage: false });

  console.log('=== 顶部菜单 ===');

  // Get ALL menu items including nested - use more specific selectors
  const allMenuItems = page.locator('.el-menu-item, .el-sub-menu__title');
  const count = await allMenuItems.count();

  // Group menus by top-level
  const menuStructure: Record<string, string[]> = {};
  let currentTopMenu = '';

  for (let i = 0; i < count; i++) {
    const item = allMenuItems.nth(i);
    const classAttr = await item.getAttribute('class') || '';
    const text = (await item.textContent())?.trim() || '';

    if (!text) continue;

    if (classAttr.includes('el-sub-menu__title')) {
      // This is a sub-menu header (category)
      currentTopMenu = text;
      if (!menuStructure[currentTopMenu]) {
        menuStructure[currentTopMenu] = [];
      }
    } else {
      // Check if it's a top-level menu item (direct child of el-menu)
      const parentClass = await item.locator('..').getAttribute('class') || '';
      if (parentClass.includes('el-menu') && !parentClass.includes('el-menu--inline')) {
        currentTopMenu = text;
        menuStructure[currentTopMenu] = [];
      } else if (currentTopMenu && menuStructure[currentTopMenu]) {
        // It's a sub-item
        menuStructure[currentTopMenu].push(text);
      }
    }
  }

  // Print structure
  for (const [topMenu, subMenus] of Object.entries(menuStructure)) {
    console.log(`- ${topMenu}:`);
    if (subMenus.length > 0) {
      for (const sub of subMenus) {
        console.log(`    - ${sub}`);
      }
    } else {
      console.log(`    (无子菜单)`);
    }
  }

  // Try direct navigation to known routes from all-pages.spec.ts
  console.log('\n=== 各页面元素 ===');

  const keyRoutes = [
    '/dashboard',
    '/mes/sop-pdf',
    '/wms/warehouse-area',
    '/bpm/task-delegate',
    '/system/user',
    '/production/order'
  ];

  for (const route of keyRoutes) {
    try {
      await page.goto(`http://localhost:5176/#${route}`);
      await page.waitForLoadState('networkidle');
      await page.waitForTimeout(1500);

      const buttons = await page.locator('button').allTextContents();
      const tables = await page.locator('.el-table').count();
      const forms = await page.locator('.el-form').count();
      const inputs = await page.locator('.el-input').count();
      const selects = await page.locator('.el-select').count();

      const btnTexts = buttons.slice(0, 8).map(b => b.trim()).filter(Boolean).join(', ');

      console.log(`\n路由: ${route}`);
      console.log(`- 按钮: ${btnTexts}${buttons.length > 8 ? '...' : ''}`);
      console.log(`- 表格: ${tables}, 表单: ${forms}, 输入框: ${inputs}, 下拉框: ${selects}`);

      await page.screenshot({ path: `D:/tmp/mom3_${route.replace(/\//g, '_')}.png`, fullPage: false });
    } catch (e) {
      console.log(`\n路由: ${route} - 错误: ${e.message}`);
    }
  }

  await page.screenshot({ path: 'D:/tmp/mom3_final.png', fullPage: false });
});
