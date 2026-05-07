const { Client } = require('pg');
const client = new Client({
  host: 'localhost',
  user: 'momadmin',
  password: 'mom123456',
  database: 'mom3.0'
});

async function run() {
  await client.connect();

  // Add IDOC management
  await client.query(`
    INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
    VALUES (1, 'IDOC管理', 'C', '/integration/idoc', 'integration/IdocList.vue', 'integration:idoc:list', 'Connection', 30, NULL, 1)
    ON CONFLICT (tenant_id, path) DO NOTHING
  `);

  // Add Mobile Job Report
  await client.query(`
    INSERT INTO sys_menu (tenant_id, menu_name, menu_type, path, component, perms, icon, sort, parent_id, status)
    VALUES (1, '移动报工', 'C', '/mes/mobile-job-report', 'mes/MobileJobReportList.vue', 'mes:mobile-job-report:list', 'Cellphone', 30, NULL, 1)
    ON CONFLICT (tenant_id, path) DO NOTHING
  `);

  // Verify
  const res = await client.query(`SELECT menu_name, path FROM sys_menu WHERE path IN ('/integration/idoc', '/mes/mobile-job-report') AND tenant_id = 1`);
  console.log('Added menus:', JSON.stringify(res.rows, null, 2));

  await client.end();
}

run().catch(console.error);
