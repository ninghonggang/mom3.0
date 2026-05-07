const { Client } = require('pg');
const client = new Client({
  host: 'localhost',
  user: 'momadmin',
  password: 'mom123456',
  database: 'mom3.0'
});

async function run() {
  await client.connect();

  // Get menu IDs
  const res = await client.query(`
    SELECT id FROM sys_menu
    WHERE tenant_id = 1 AND path IN (
      '/wms/vmi', '/wms/vmi/vendor', '/wms/vmi/material', '/wms/vmi/transaction',
      '/quality/certificate',
      '/scp/customer-credit',
      '/integration/idoc',
      '/mes/mobile-job-report'
    )
  `);

  console.log('Found menus:', res.rows);

  // Assign to admin role (role_id = 1)
  for (const row of res.rows) {
    try {
      await client.query(`
        INSERT INTO sys_role_menu (role_id, menu_id)
        VALUES (1, $1)
      `, [row.id]);
      console.log('Assigned menu', row.id, 'to role 1');
    } catch (e) {
      if (e.code === '23505') console.log('Menu', row.id, 'already assigned to role 1');
      else throw e;
    }
  }

  await client.end();
  console.log('Done');
}

run().catch(console.error);
