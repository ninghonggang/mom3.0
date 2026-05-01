const { chromium } = require('playwright');
const fs = require('fs');

const routes = JSON.parse(fs.readFileSync('/tmp/routes.json', 'utf8'));
const BASE = 'http://localhost:5175';
const results = [];
let browser;

async function init() {
  browser = await chromium.launch({ headless: true });
  const context = await browser.newContext();
  const page = await context.newPage();
  
  // Login first
  await page.goto(BASE + '/login', { waitUntil: 'networkidle' });
  await page.fill('input[placeholder="请输入用户名"]', 'admin');
  await page.fill('input[placeholder="请输入密码"]', 'admin123');
  await page.click('button:has-text("登 录")');
  await page.waitForURL('**/*', { timeout: 15000 });
  await page.waitForTimeout(3000); // Wait for post-login redirect
  console.log('Login successful');
  
  return { browser, context, page };
}

async function testRoute(page, route) {
  const errors = [];
  const warnings = [];
  
  page.on('console', msg => {
    if (msg.type() === 'error') errors.push(msg.text());
  });
  page.on('pageerror', err => errors.push('PAGE ERROR: ' + err.message));
  
  try {
    const response = await page.goto(BASE + route, { 
      waitUntil: 'networkidle', 
      timeout: 20000 
    });
    const status = response ? response.status() : 0;
    
    // Wait a bit for async errors
    await page.waitForTimeout(2000);
    
    const title = await page.title();
    const url = page.url();
    
    results.push({
      route,
      status,
      url,
      title,
      errors: errors.filter(e => !e.includes('401') && !e.includes('NProgress')),
      warnings: warnings.length
    });
  } catch (e) {
    results.push({
      route,
      status: 0,
      url: '',
      title: '',
      errors: ['NAVIGATION ERROR: ' + e.message],
      warnings: 0
    });
  }
}

async function main() {
  const { browser, page } = await init();
  
  for (let i = 0; i < routes.length; i++) {
    const route = routes[i];
    process.stdout.write(`[${i+1}/${routes.length}] ${route}... `);
    await testRoute(page, route);
    const last = results[results.length - 1];
    if (last.errors.length > 0) {
      process.stdout.write(`ERRORS: ${last.errors.length}\n`);
    } else {
      process.stdout.write(`OK\n`);
    }
  }
  
  await browser.close();
  
  // Write results
  const failed = results.filter(r => r.errors.length > 0 || r.status >= 400);
  const passed = results.filter(r => r.errors.length === 0 && r.status < 400);
  
  fs.writeFileSync('/tmp/mom-test-results.json', JSON.stringify(results, null, 2));
  
  console.log('\n=== SUMMARY ===');
  console.log(`Total: ${results.length}`);
  console.log(`Passed: ${passed.length}`);
  console.log(`Failed/Errors: ${failed.length}`);
  console.log('\n=== ERRORS ===');
  failed.forEach(r => {
    console.log(`\n${r.route} (status:${r.status})`);
    r.errors.forEach(e => console.log(`  - ${e.substring(0, 200)}`));
  });
}

main().catch(console.error);
