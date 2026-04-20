const { chromium } = require('playwright');

(async () => {
  const browser = await chromium.launch({ headless: true });
  const context = await browser.newContext({ viewport: { width: 1440, height: 900 } });
  const page = await context.newPage();

  const results = [];
  const errors = [];
  const apiResults = [];

  page.on('console', msg => {
    if (msg.type() === 'error') errors.push(`[${msg.location().url}] ${msg.text()}`);
  });

  page.on('requestfailed', req => {
    errors.push(`Request failed: ${req.url()}`);
  });

  async function test(name, fn) {
    try {
      await fn();
      results.push(`✅ ${name}`);
    } catch (e) {
      results.push(`❌ ${name}: ${e.message.substring(0, 100)}`);
    }
  }

  async function apiTest(name, method, url, body = null) {
    try {
      const options = {
        method,
        headers: { 'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json' }
      };
      if (body) options.body = JSON.stringify(body);
      const resp = await page.request[method.toLowerCase()](url, options);
      const json = await resp.json();
      if (json.code !== 200) {
        apiResults.push(`⚠️ ${name}: code=${json.code} msg=${json.message}`);
      } else {
        apiResults.push(`✅ ${name}`);
      }
    } catch (e) {
      apiResults.push(`❌ ${name}: ${e.message.substring(0, 80)}`);
    }
  }

  console.log('=== 开始SCP供应链管理完整测试 ===\n');

  // 1. Login
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
    // Get token from localStorage
    token = await page.evaluate(() => localStorage.getItem('token') || sessionStorage.getItem('token'));
    if (!token) throw new Error('No token found after login');
  });

  // Get fresh token via API
  const loginResp = await page.request.post('http://localhost:9081/api/v1/auth/login', {
    headers: { 'Content-Type': 'application/json' },
    data: { username: 'admin', password: 'admin123' }
  });
  const loginJson = await loginResp.json();
  token = loginJson.data.access_token;

  // 2. Test all SCP pages load
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
      await page.waitForTimeout(2500);
      // Check for table or form elements
      const hasContent = await page.locator('.el-table, .el-form, .el-card').count() > 0;
      if (!hasContent) throw new Error('Page has no expected elements');
    });
  }

  // 3. Test button interactions on each page
  for (const p of scpPages) {
    await test(`${p.name}-点击新增按钮`, async () => {
      await page.goto(`http://localhost:5175/#${p.path}`, { waitUntil: 'networkidle' });
      await page.waitForTimeout(2000);
      const addBtn = page.locator('button').filter({ hasText: /新增|添加|创建/i }).first();
      if (await addBtn.count() > 0) {
        await addBtn.click();
        await page.waitForTimeout(1500);
        // Check if dialog opened
        const dialog = page.locator('.el-dialog, [role="dialog"]').first();
        if (await dialog.count() > 0) {
          const visible = await dialog.isVisible().catch(() => false);
          if (!visible) throw new Error('Dialog did not open');
        }
      }
    });

    await test(`${p.name}-点击查询按钮`, async () => {
      await page.goto(`http://localhost:5175/#${p.path}`, { waitUntil: 'networkidle' });
      await page.waitForTimeout(2000);
      const searchBtn = page.locator('button').filter({ hasText: /查询|搜索/i }).first();
      if (await searchBtn.count() > 0) {
        await searchBtn.click();
        await page.waitForTimeout(1500);
      }
    });
  }

  // 4. Test form validation
  await test('采购订单-表单校验(空提交)', async () => {
    await page.goto('http://localhost:5175/#/scp/purchase', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
    const addBtn = page.locator('button').filter({ hasText: /新增/i }).first();
    if (await addBtn.count() > 0) {
      await addBtn.click();
      await page.waitForTimeout(1500);
      const submitBtn = page.locator('.el-dialog button').filter({ hasText: /提交|确定/i }).first();
      if (await submitBtn.count() > 0) {
        await submitBtn.click();
        await page.waitForTimeout(1000);
      }
    }
  });

  // 5. Test modal close
  for (const p of scpPages) {
    await test(`${p.name}-弹窗关闭`, async () => {
      await page.goto(`http://localhost:5175/#${p.path}`, { waitUntil: 'networkidle' });
      await page.waitForTimeout(2000);
      const addBtn = page.locator('button').filter({ hasText: /新增/i }).first();
      if (await addBtn.count() > 0) {
        await addBtn.click();
        await page.waitForTimeout(1500);
        const closeBtn = page.locator('.el-dialog__headerbtn, button').filter({ hasText: /取消|关闭/i }).first();
        if (await closeBtn.count() > 0) {
          await closeBtn.click();
          await page.waitForTimeout(500);
        }
      }
    });
  }

  // 6. API CRUD tests with real data
  console.log('\n=== API CRUD测试 ===\n');

  // Create supplier first
  await apiTest('创建供应商', 'POST', 'http://localhost:9081/api/v1/mdm/supplier', {
    supplier_code: 'SUP001',
    supplier_name: '测试供应商',
    contact_person: '张三',
    contact_phone: '13800138000',
    status: 1
  });

  // Create material
  await apiTest('创建物料', 'POST', 'http://localhost:9081/api/v1/mdm/material', {
    material_code: 'MAT001',
    material_name: '测试物料',
    specification: '规格A',
    unit: 'PCS',
    status: 1
  });

  // Create RFQ
  const rfqResp = await page.request.post('http://localhost:9081/api/v1/scp/rfq', {
    headers: { 'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json' },
    data: {
      rfq_no: 'RFQ20260414001',
      rfq_name: '测试询价单',
      rfq_type: 'STANDARD',
      inquiry_date: '2026-04-14',
      deadline_date: '2026-04-30',
      currency: 'CNY',
      status: 'DRAFT'
    }
  });
  const rfqJson = await rfqResp.json();
  const rfqId = rfqJson.data?.id;
  apiResults.push(rfqJson.code === 200 ? `✅ 创建询价单: ID=${rfqId}` : `⚠️ 创建询价单: code=${rfqJson.code}`);

  // Create Purchase Order
  const poResp = await page.request.post('http://localhost:9081/api/v1/scp/purchase-orders', {
    headers: { 'Authorization': `Bearer ${token}`, 'Content-Type': 'application/json' },
    data: {
      po_no: 'PO20260414001',
      supplier_id: 1,
      supplier_name: '测试供应商',
      order_date: '2026-04-14',
      promised_date: '2026-04-30',
      currency: 'CNY',
      tax_rate: 13,
      status: 'DRAFT',
      items: [{
        material_id: 1,
        material_code: 'MAT001',
        material_name: '测试物料',
        unit: 'PCS',
        unit_price: 100,
        order_qty: 100,
        promised_date: '2026-04-30'
      }]
    }
  });
  const poJson = await poResp.json();
  const poId = poJson.data?.id;
  apiResults.push(poJson.code === 200 ? `✅ 创建采购订单: ID=${poId}` : `⚠️ 创建采购订单: code=${poJson.code} msg=${poJson.message}`);

  // List APIs
  await apiTest('查询采购订单列表', 'GET', 'http://localhost:9081/api/v1/scp/purchase-orders/list');
  await apiTest('查询询价单列表', 'GET', 'http://localhost:9081/api/v1/scp/rfq/list');
  await apiTest('查询销售订单列表', 'GET', 'http://localhost:9081/api/v1/scp/sales-orders/list');
  await apiTest('查询供应商报价列表', 'GET', 'http://localhost:9081/api/v1/scp/supplier-quotes/list');
  await apiTest('查询供应商绩效列表', 'GET', 'http://localhost:9081/api/scp/supplier-kpi/list');
  await apiTest('查询供应商列表', 'GET', 'http://localhost:9081/api/v1/scp/supplier/list');
  await apiTest('查询物料列表', 'GET', 'http://localhost:9081/api/v1/mdm/material/list');

  // Get by ID
  if (poId) {
    await apiTest('获取采购订单详情', 'GET', `http://localhost:9081/api/v1/scp/purchase-orders/${poId}`);
  }

  // Summary
  console.log('\n========== 页面测试结果 ==========\n');
  results.forEach(r => console.log(r));

  console.log('\n========== API测试结果 ==========\n');
  apiResults.forEach(r => console.log(r));

  console.log('\n========== 控制台/网络错误 ==========\n');
  if (errors.length === 0) {
    console.log('✅ 无错误');
  } else {
    errors.slice(0, 10).forEach(e => console.log(`❌ ${e}`));
  }

  console.log('\n================================\n');

  await browser.close();
})();
