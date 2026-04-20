const { chromium } = require('playwright');

(async () => {
  const browser = await chromium.launch({ headless: true });
  const context = await browser.newContext({ viewport: { width: 1440, height: 900 } });
  const page = await context.newPage();

  const results = [];
  const consoleErrors = [];

  page.on('console', msg => {
    if (msg.type() === 'error') consoleErrors.push(msg.text());
  });

  async function test(name, fn) {
    try {
      await fn();
      results.push(`✅ ${name}`);
    } catch (e) {
      results.push(`❌ ${name}: ${e.message}`);
    }
  }

  // 1. Login
  await test('登录页面加载', async () => {
    await page.goto('http://localhost:5175/', { waitUntil: 'networkidle' });
    await page.waitForTimeout(2000);
  });

  await test('登录功能', async () => {
    const usernameInput = page.locator('input').first();
    if (await usernameInput.count() > 0) {
      await usernameInput.fill('admin');
      const passwordInput = page.locator('input[type="password"]').first();
      if (await passwordInput.count() > 0) {
        await passwordInput.fill('admin123');
        await page.waitForTimeout(500);
        const submitBtn = page.locator('button[type="submit"]').first();
        if (await submitBtn.count() > 0) {
          await submitBtn.click();
          await page.waitForTimeout(3000);
        }
      }
    }
  });

  // 2. Navigate through main menu items
  const menuItems = [
    { name: '首页', path: '/dashboard' },
    { name: '生产执行', path: '/production' },
    { name: '质量管理', path: '/quality' },
    { name: '仓库管理', path: '/warehouse' },
    { name: '设备管理', path: '/equipment' },
    { name: 'APS计划', path: '/aps' },
    { name: '系统管理', path: '/system' },
  ];

  for (const item of menuItems) {
    await test(`菜单 - ${item.name}`, async () => {
      const menuLink = page.locator(`a[href*="${item.path}"]`).first();
      if (await menuLink.count() > 0) {
        await menuLink.click();
        await page.waitForTimeout(1500);
      }
    });
  }

  // 3. Test new pages (AGV, ASN, MES)
  const newPages = [
    { name: 'AGV调度', path: '/agv/task' },
    { name: '供应商ASN', path: '/supplier-asn/asn' },
    { name: 'MES班组', path: '/mes/team' },
    { name: 'MES工艺路线', path: '/mes/process-routes' },
    { name: 'MES产品离线', path: '/mes/offline' },
    { name: '系统集成', path: '/integration/interface-config' },
    { name: 'ERP同步', path: '/integration/erp/sync-log' },
  ];

  for (const p of newPages) {
    await test(`新页面 - ${p.name}`, async () => {
      await page.goto(`http://localhost:5175/#${p.path}`, { waitUntil: 'networkidle' });
      await page.waitForTimeout(2000);
    });
  }

  // API tests with correct paths
  const apiTests = [
    { name: 'API - 菜单列表', path: '/api/v1/system/menu/list' },
    { name: 'API - AGV任务', path: '/api/v1/agv/task/list' },
    { name: 'API - 供应商ASN', path: '/api/v1/supplier/asn/list' },
    { name: 'API - MES班组', path: '/api/v1/mes/team/list' },
    { name: 'API - 工艺路线', path: '/api/v1/mes/process-routes/list' },
    { name: 'API - 产品离线', path: '/api/v1/mes/offline/list' },
    { name: 'API - 接口配置', path: '/api/v1/integration/interface-config/list' },
  ];

  for (const api of apiTests) {
    await test(api.name, async () => {
      const resp = await page.request.get(`http://localhost:9081${api.path}`);
      if (!resp.ok()) throw new Error(`HTTP ${resp.status()}`);
      const json = await resp.json();
      if (json.code && json.code !== 200) throw new Error(`API error: ${json.message}`);
    });
  }

  // Console errors
  const criticalErrors = consoleErrors.filter(e =>
    !e.includes('warning') &&
    !e.includes('devtools') &&
    !e.includes('favicon') &&
    !e.includes('Failed to load resource')
  );

  console.log('\n========== QA 测试结果 ==========\n');
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
