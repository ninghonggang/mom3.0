import { test, expect } from '@playwright/test';

test('debug login flow', async ({ page }) => {
  const consoleMessages: string[] = [];
  const networkRequests: string[] = [];

  page.on('console', msg => {
    consoleMessages.push(`[${msg.type()}] ${msg.text()}`);
  });

  page.on('request', req => {
    if (req.url().includes('/api/')) {
      networkRequests.push(`→ ${req.method()} ${req.url()}`);
    }
  });

  page.on('response', res => {
    if (res.url().includes('/api/')) {
      networkRequests.push(`← ${res.status()} ${res.url()}`);
    }
  });

  await page.goto('login');
  await page.waitForLoadState('networkidle');

  console.log('Page loaded, filling credentials...');

  await page.fill('input[type="text"]', 'admin');
  await page.fill('input[type="password"]', 'admin123');

  console.log('Credentials filled, clicking login...');

  // Click and wait for response
  await page.click('button:has-text("登 录")');

  // Wait a bit for any responses
  await page.waitForTimeout(3000);

  console.log('\n=== Network Requests ===');
  networkRequests.forEach(r => console.log(r));

  console.log('\n=== Console Messages ===');
  consoleMessages.filter(m => m.includes('error') || m.includes('Error')).forEach(m => console.log(m));

  console.log('\nCurrent URL:', page.url());

  await page.screenshot({ path: 'D:/tmp/mom3_login_debug.png' });
});
