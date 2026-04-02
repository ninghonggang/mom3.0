const { Client } = require('pg');
const fs = require('fs');

async function runSQL() {
  const client = new Client({
    host: 'localhost',
    port: 5432,
    user: 'mes',
    password: 'mes123',
    database: 'mom3_db'
  });

  try {
    await client.connect();
    const sql = fs.readFileSync('./reset-menus.sql', 'utf8');
    await client.query(sql);
    console.log('SQL executed successfully');
  } catch (err) {
    console.error('Error:', err.message);
  } finally {
    await client.end();
  }
}

runSQL();
