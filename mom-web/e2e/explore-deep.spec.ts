import { test, expect } from '@playwright/test';

test('deep page exploration', async ({ page }) => {
  await page.goto('http://localhost:5176/');
  await page.waitForLoadState('networkidle');

  await page.fill('input[type="text"]', 'admin');
  await page.fill('input[type="password"]', 'admin123');
  await page.click('button:has-text("登 录")');
  await page.waitForURL('**/dashboard', { timeout: 15000 });
  await page.waitForTimeout(3000);

  await page.screenshot({ path: 'D:/tmp/mom3_dashboard.png', fullPage: false });

  // Navigate to a specific page
  await page.goto('http://localhost:5176/mdm/material');
  await page.waitForLoadState('networkidle');
  await page.waitForTimeout(2000);

  // Get full page structure
  const pageInfo = await page.evaluate(() => {
    const app = document.querySelector('#app');
    return {
      appHTML: app?.innerHTML.substring(0, 3000),
      url: window.location.href,
      hash: window.location.hash
    };
  });

  console.log('URL:', pageInfo.url);
  console.log('Hash:', pageInfo.hash);
  console.log('App HTML (first 3000 chars):\n', pageInfo.appHTML);

  // Look for route-specific content
  const routeContent = await page.evaluate(() => {
    // Find any element that contains route-specific text
    const body = document.body;
    const allText = body.innerText;

    // Check for specific indicators
    return {
      hasMaterial: allText.includes('物料'),
      hasTable: !!document.querySelector('.el-table'),
      hasForm: !!document.querySelector('.el-form'),
      allClasses: Array.from(document.querySelectorAll('[class]')).map(e => e.className).join(', ').substring(0, 2000)
    };
  });

  console.log('Route content:', routeContent);

  await page.screenshot({ path: 'D:/tmp/mom3_material_page.png', fullPage: true });
});
