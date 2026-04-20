const { chromium } = require('playwright');

(async () => {
  const browser = await chromium.launch({ headless: true });
  const context = await browser.newContext({ viewport: { width: 1440, height: 900 } });
  const page = await context.newPage();

  const results = [];
  const errors = [];

  page.on('console', msg => {
    if (msg.type() === 'error') errors.push(msg.text());
  });

  async function test(name, fn) {
    try {
      await fn();
      results.push(`✅ ${name}`);
    } catch (e) {
      results.push(`❌ ${name}: ${e.message.substring(0, 80)}`);
    }
  }

  // 1. Login
  console.log('=== 开始SCP供应链管理L6级测试 ===\n');

  await test('登录', async () => {
    await page.goto('http://localhost:5175/', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    const usernameInput = page.locator('input').first();
    await usernameInput.fill('admin');
    const passwordInput = page.locator('input[type="password"]').first();
    await passwordInput.fill('admin123');
    await page.waitForTimeout(500);
    const submitBtn = page.locator('button[type="submit"]').first();
    await submitBtn.click();
    await page.waitForTimeout(3000);
  });

  // 2. Navigate to SCP module
  await test('打开供应链管理菜单', async () => {
    const scpMenu = page.locator('a[href*="/scp"]').first();
    if (await scpMenu.count() > 0) {
      await scpMenu.click();
      await page.waitForTimeout(1500);
    }
  });

  // 3. Test each SCP page
  const scpPages = [
    { name: '采购订单', path: '/scp/purchase' },
    { name: '询价单', path: '/scp/rfq' },
    { name: '供应商报价', path: '/scp/supplier-quote' },
    { name: '销售订单', path: '/scp/sales-order' },
    { name: '供应商绩效', path: '/scp/supplier-kpi' },
  ];

  for (const p of scpPages) {
    await test(`${p.name}页面加载`, async () => {
      await page.goto(`http://localhost:5175/#${p.path}`, { waitUntil: 'networkidle' });
      await page.waitForTimeout(2000);
      // Check page has content
      const content = await page.textContent('body');
      if (!content || content.length < 100) throw new Error('Page appears empty');
    });
  }

  // 4. Test buttons on Purchase Order page
  await test('采购订单-新增按钮', async () => {
    await page.goto('http://localhost:5175/#/scp/purchase', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    const addBtn = page.locator('button').filter({ hasText: /新增|添加/i }).first();
    if (await addBtn.count() > 0) {
      await addBtn.click();
      await page.waitForTimeout(1500);
      // Check if modal/dialog opened
      const dialog = page.locator('.el-dialog, .ant-modal, [role="dialog"]').first();
      if (await dialog.count() > 0) {
        // Close it
        const closeBtn = page.locator('button').filter({ hasText: /取消|关闭/i }).first();
        if (await closeBtn.count() > 0) await closeBtn.click();
      }
    }
  });

  await test('采购订单-查询按钮', async () => {
    await page.goto('http://localhost:5175/#/scp/purchase', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    const searchBtn = page.locator('button').filter({ hasText: /查询|搜索/i }).first();
    if (await searchBtn.count() > 0) {
      await searchBtn.click();
      await page.waitForTimeout(1500);
    }
  });

  // 5. Test RFQ page
  await test('询价单-新增按钮', async () => {
    await page.goto('http://localhost:5175/#/scp/rfq', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    const addBtn = page.locator('button').filter({ hasText: /新增|添加/i }).first();
    if (await addBtn.count() > 0) {
      await addBtn.click();
      await page.waitForTimeout(1500);
      const closeBtn = page.locator('button').filter({ hasText: /取消|关闭/i }).first();
      if (await closeBtn.count() > 0) await closeBtn.click();
    }
  });

  // 6. Test Supplier Quote page
  await test('供应商报价-新增按钮', async () => {
    await page.goto('http://localhost:5175/#/scp/supplier-quote', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    const addBtn = page.locator('button').filter({ hasText: /新增|添加/i }).first();
    if (await addBtn.count() > 0) {
      await addBtn.click();
      await page.waitForTimeout(1500);
      const closeBtn = page.locator('button').filter({ hasText: /取消|关闭/i }).first();
      if (await closeBtn.count() > 0) await closeBtn.click();
    }
  });

  // 7. Test Sales Order page
  await test('销售订单-新增按钮', async () => {
    await page.goto('http://localhost:5175/#/scp/sales-order', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    const addBtn = page.locator('button').filter({ hasText: /新增|添加/i }).first();
    if (await addBtn.count() > 0) {
      await addBtn.click();
      await page.waitForTimeout(1500);
      const closeBtn = page.locator('button').filter({ hasText: /取消|关闭/i }).first();
      if (await closeBtn.count() > 0) await closeBtn.click();
    }
  });

  // 8. Test Supplier KPI page
  await test('供应商绩效-查询按钮', async () => {
    await page.goto('http://localhost:5175/#/scp/supplier-kpi', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    const searchBtn = page.locator('button').filter({ hasText: /查询|搜索/i }).first();
    if (await searchBtn.count() > 0) {
      await searchBtn.click();
      await page.waitForTimeout(1500);
    }
  });

  // 9. API tests
  const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ0ZW5hbnRfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4iLCJyb2xlcyI6WyJhZG1pbiJdLCJpc3MiOiJtb20tc2VydmVyIiwiZXhwIjoxNzc2MTM1OTY0LCJuYmYiOjE3NzYxMjg3NjQsImlhdCI6MTc3NjEyODc2NH0.5SXi2-Fc2Q1ulo7N-MfQS74azsFC8Hg8w8KNuN4nsbY';

  const apiTests = [
    { name: 'API-采购订单列表', path: '/api/v1/scp/purchase-order/list' },
    { name: 'API-询价单列表', path: '/api/v1/scp/rfq/list' },
    { name: 'API-供应商报价列表', path: '/api/v1/scp/supplier-quote/list' },
    { name: 'API-销售订单列表', path: '/api/v1/scp/sales-order/list' },
    { name: 'API-供应商绩效列表', path: '/api/v1/scp/supplier-kpi/list' },
  ];

  for (const api of apiTests) {
    await test(api.name, async () => {
      const resp = await page.request.get(`http://localhost:9081${api.path}`, {
        headers: { 'Authorization': `Bearer ${token}` }
      });
      if (!resp.ok()) throw new Error(`HTTP ${resp.status()}`);
      const json = await resp.json();
      if (json.code && json.code !== 200) throw new Error(`API error: ${json.message}`);
    });
  }

  // Critical errors
  const criticalErrors = errors.filter(e =>
    !e.includes('warning') &&
    !e.includes('devtools') &&
    !e.includes('favicon') &&
    !e.includes('Failed to load resource')
  );

  console.log('\n========== SCP L6测试结果 ==========\n');
  results.forEach(r => console.log(r));
  console.log('\n========== 控制台错误 ==========\n');
  if (criticalErrors.length === 0) {
    console.log('✅ 无关键控制台错误');
  } else {
    console.log(`⚠️ ${criticalErrors.length} 个错误:`);
    criticalErrors.slice(0, 5).forEach(e => console.log(`  - ${e.substring(0, 100)}`));
  }
  console.log('\n================================\n');

  await browser.close();
})();
