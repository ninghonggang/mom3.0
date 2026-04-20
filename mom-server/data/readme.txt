
● 已完成。数据库初始化脚本已生成并提交：

  文件位置: mom-server/data/init.sql

  包含内容:
  - 40+ 表结构 (系统管理、主数据、设备、生产执行、仓储、APS、质量、追溯、能源)
  - 所有表索引
  - 初始化数据 (租户、用户、角色、菜单、字典数据)
  - 示例数据 (车间、生产线、工位、仓库、设备)

  使用方法:
  # 创建数据库
  psql -U postgres -c "CREATE DATABASE mom3;"

  # 执行初始化脚本
  psql -U postgres -d mom3 -f mom-server/data/init.sql

  默认账号: admin / admin123
