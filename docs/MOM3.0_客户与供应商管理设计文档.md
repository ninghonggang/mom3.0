# MOM3.0 客户与供应商管理设计文档

**版本**: V1.0 | **所属模块**: M02主数据管理扩展 | **项目**: 闻荫科技MOM3.0

---

## 1. 概述

### 1.1 定位

客户与供应商管理是MOM3.0供应链闭环的基础。主数据管理模块(M02)需要同时管理：
- **供应商主数据**：采购合作伙伴、物料供应商
- **客户主数据**：销售订单客户、经销商、终端客户

### 1.2 核心实体关系

```
                    ┌─────────────────┐
                    │   供应商主数据   │
                    │ mdm_supplier    │
                    └────────┬────────┘
                             │
              ┌──────────────┼──────────────┐
              │              │              │
              ▼              ▼              ▼
    ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
    │ 采购信息扩展 │  │ 绩效评估    │  │ ASN管理     │
    │ SCP模块     │  │ SCP模块     │  │ WMS模块     │
    └─────────────┘  └─────────────┘  └─────────────┘

                    ┌─────────────────┐
                    │   客户主数据     │
                    │ mdm_customer    │
                    └────────┬────────┘
                             │
              ┌──────────────┴──────────────┐
              │              │              │
              ▼              ▼              ▼
    ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
    │ 销售订单    │  │ 信用管理    │  │ 发货地址    │
    │ SCP模块     │  │ SCP模块     │  │ WMS模块     │
    └─────────────┘  └─────────────┘  └─────────────┘
```

---

## 2. 数据模型设计

### 2.1 供应商主数据

```sql
-- 供应商主数据
CREATE TABLE mdm_supplier (
    id BIGSERIAL PRIMARY KEY,
    supplier_code VARCHAR(50) UNIQUE NOT NULL,
    supplier_name VARCHAR(200) NOT NULL,
    supplier_type VARCHAR(20) NOT NULL,     -- MANUFACTURER/TRADER/SERVICE
    category VARCHAR(50),                   -- 供应商类别：原材料/辅料/包材/设备/服务
    legal_person VARCHAR(50),               -- 法定代表人
    registered_capital DECIMAL(18,2),       -- 注册资本
    established_date DATE,                 -- 成立日期
    
    -- 地址信息
    country VARCHAR(50) DEFAULT '中国',
    province VARCHAR(50),
    city VARCHAR(50),
    district VARCHAR(50),
    address_detail TEXT,
    longitude DECIMAL(10,6),
    latitude DECIMAL(10,6),
    
    -- 联系信息
    contact_person VARCHAR(50),
    contact_phone VARCHAR(50),
    contact_mobile VARCHAR(50),
    contact_email VARCHAR(100),
    contact_fax VARCHAR(50),
    website VARCHAR(200),
    
    -- 银行信息
    bank_name VARCHAR(100),
    bank_account VARCHAR(100),
    bank_account_name VARCHAR(100),
    tax_no VARCHAR(50),                    -- 税号
    
    -- 经营状态
    business_status VARCHAR(20) DEFAULT 'ACTIVE', -- ACTIVE/SUSPENDED/REVOKED
    cooperation_status VARCHAR(20) DEFAULT 'POTENTIAL', -- POTENTIAL/ACTIVE/PROBATION/BLACKLIST
    first_cooperation_date DATE,
    
    -- 分类评级
    supplier_grade VARCHAR(10),             -- A/B/C/D
    is_preferred SMALLINT DEFAULT 0,       -- 是否首选供应商
    is_blacklist SMALLINT DEFAULT 0,        -- 是否黑名单
    
    -- 审核状态
    audit_status VARCHAR(20) DEFAULT 'PENDING', -- PENDING/APPROVED/REJECTED
    approved_by BIGINT,
    approved_time TIMESTAMP,
    audit_remark TEXT,
    
    remark TEXT,
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 供应商联系人（多个）
CREATE TABLE mdm_supplier_contact (
    id BIGSERIAL PRIMARY KEY,
    supplier_id BIGINT NOT NULL,
    contact_name VARCHAR(50) NOT NULL,
    department VARCHAR(100),
    position VARCHAR(50),
    phone VARCHAR(50),
    mobile VARCHAR(50),
    email VARCHAR(100),
    is_primary SMALLINT DEFAULT 0,         -- 是否主联系人
    remark VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 供应商银行账户（多个）
CREATE TABLE mdm_supplier_bank_account (
    id BIGSERIAL PRIMARY KEY,
    supplier_id BIGINT NOT NULL,
    bank_name VARCHAR(100) NOT NULL,
    bank_branch VARCHAR(200),
    bank_account VARCHAR(100) NOT NULL,
    account_name VARCHAR(100) NOT NULL,
    account_type VARCHAR(20),              -- 基本户/一般户
    is_primary SMALLINT DEFAULT 0,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 供应商附件（营业执照等）
CREATE TABLE mdm_supplier_attachment (
    id BIGSERIAL PRIMARY KEY,
    supplier_id BIGINT NOT NULL,
    attachment_type VARCHAR(50) NOT NULL,  -- BUSINESS_LICENSE/Tax_REG/CERTIFICATE/OTHER
    attachment_name VARCHAR(200),
    file_url VARCHAR(500) NOT NULL,
    file_size BIGINT,
    expiry_date DATE,
    is_expired SMALLINT DEFAULT 0,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 2.2 客户主数据

```sql
-- 客户主数据
CREATE TABLE mdm_customer (
    id BIGSERIAL PRIMARY KEY,
    customer_code VARCHAR(50) UNIQUE NOT NULL,
    customer_name VARCHAR(200) NOT NULL,
    customer_type VARCHAR(20) NOT NULL,     -- DEALER/DIRECT/END_USER/INTERNAL
    category VARCHAR(50),                   -- 客户类别：汽车整车/零部件/经销商
    customer_level VARCHAR(10),            -- A/B/C级客户
    
    -- 基本信息
    legal_person VARCHAR(50),
    registered_capital DECIMAL(18,2),
    
    -- 地址信息
    country VARCHAR(50) DEFAULT '中国',
    province VARCHAR(50),
    city VARCHAR(50),
    district VARCHAR(50),
    address_detail TEXT,
    
    -- 联系信息
    contact_person VARCHAR(50),
    contact_phone VARCHAR(50),
    contact_mobile VARCHAR(50),
    contact_email VARCHAR(100),
    
    -- 银行信息
    bank_name VARCHAR(100),
    bank_account VARCHAR(100),
    bank_account_name VARCHAR(100),
    tax_no VARCHAR(50),
    
    -- 信用管理
    credit_limit DECIMAL(18,2) DEFAULT 0,   -- 信用额度
    credit_used DECIMAL(18,2) DEFAULT 0,   -- 已使用额度
    payment_terms VARCHAR(50),             -- 付款条款
    payment_days INTEGER DEFAULT 0,         -- 账期天数
    settlement_cycle VARCHAR(20),          -- MONTHLY/QUARTERLY
    
    -- 客户分类
    industry VARCHAR(50),                   -- 所属行业
    customer_source VARCHAR(50),            -- 客户来源
    is_key_customer SMALLINT DEFAULT 0,   -- 是否重点客户
    is_active SMALLINT DEFAULT 1,
    
    -- 审核状态
    audit_status VARCHAR(20) DEFAULT 'PENDING',
    approved_by BIGINT,
    approved_time TIMESTAMP,
    
    remark TEXT,
    tenant_id BIGINT NOT NULL,
    created_by VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 客户联系人
CREATE TABLE mdm_customer_contact (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    contact_name VARCHAR(50) NOT NULL,
    department VARCHAR(100),
    position VARCHAR(50),
    phone VARCHAR(50),
    mobile VARCHAR(50),
    email VARCHAR(100),
    is_primary SMALLINT DEFAULT 0,
    remark VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 客户收货地址（多个）
CREATE TABLE mdm_customer_delivery_address (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    address_name VARCHAR(100),             -- 地址别名
    contact_person VARCHAR(50),
    contact_phone VARCHAR(50),
    province VARCHAR(50),
    city VARCHAR(50),
    district VARCHAR(50),
    address_detail TEXT,
    is_default SMALLINT DEFAULT 0,        -- 是否默认地址
    is_active SMALLINT DEFAULT 1,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- 客户银行账户
CREATE TABLE mdm_customer_bank_account (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    bank_name VARCHAR(100) NOT NULL,
    bank_branch VARCHAR(200),
    bank_account VARCHAR(100) NOT NULL,
    account_name VARCHAR(100 NOT NULL,
    is_primary SMALLINT DEFAULT 0,
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 2.3 供应商采购信息扩展（SCP模块使用）

```sql
-- 供应商采购信息
CREATE TABLE scp_supplier_purchase_info (
    id BIGSERIAL PRIMARY KEY,
    supplier_id BIGINT NOT NULL UNIQUE,
    payment_terms VARCHAR(50),                -- 付款条款：T/T/L/C/OA
    credit_limit DECIMAL(18,2),             -- 信用额度
    tax_rate DECIMAL(5,2) DEFAULT 13.00,  -- 税率
    min_order_amount DECIMAL(18,2),         -- 最小订单量
    min_order_qty DECIMAL(18,3),           -- 最小订单数量
    lead_time_days INTEGER DEFAULT 0,       -- 标准交货周期
    emergency_lead_time_days INTEGER DEFAULT 0, -- 紧急交货周期
    
    -- 交货条款
    delivery_terms VARCHAR(100),           -- EXW/FOB/CIF/DDU
    
    -- 质量要求
    quality_standard VARCHAR(100),         -- 质量标准
    quality_certification VARCHAR(200),    -- 质量认证
    is_first_article_required SMALLINT DEFAULT 1, -- 是否需要首件确认
    
    -- 合作状态
    supplier_grade VARCHAR(10),           -- 供应商等级：A/B/C/D
    is_preferred SMALLINT DEFAULT 0,      -- 是否首选供应商
    is_blacklist SMALLINT DEFAULT 0,       -- 是否黑名单
    blacklist_reason VARCHAR(200),
    blacklist_date DATE,
    
    -- 合作统计
    cooperation_start_date DATE,
    total_orders INTEGER DEFAULT 0,
    total_cooperation_amount DECIMAL(18,2) DEFAULT 0, -- 累计合作金额
    last_order_date DATE,
    
    tenant_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- 供应商物料关系（供应商能供应的物料）
CREATE TABLE scp_supplier_material (
    id BIGSERIAL PRIMARY KEY,
    supplier_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    material_code VARCHAR(50),
    material_name VARCHAR(100),
    is_approved_source SMALLINT DEFAULT 0,  -- 是否已批准供应商
    is_alternative_source SMALLINT DEFAULT 0, -- 是否可替代供应商
    is_preferred SMALLINT DEFAULT 0,       -- 是否首选物料
    min_supply_qty DECIMAL(18,3),         -- 最小供应量
    standard_lead_time INTEGER,           -- 标准交货周期
    standard_price DECIMAL(18,4),         -- 标准价格
    price_effective_date DATE,
    quality_agreement_url VARCHAR(500),
    remark VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    UNIQUE(supplier_id, material_id),
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

## 3. API设计

### 3.1 供应商管理API

```
# 供应商主数据
GET    /api/v1/mdm/suppliers                    # 供应商列表
POST   /api/v1/mdm/suppliers                    # 创建供应商
GET    /api/v1/mdm/suppliers/:id               # 供应商详情
PUT    /api/v1/mdm/suppliers/:id               # 更新供应商
DELETE /api/v1/mdm/suppliers/:id               # 删除供应商(仅PENDING)
POST   /api/v1/mdm/suppliers/:id/approve        # 审批通过
POST   /api/v1/mdm/suppliers/:id/reject         # 审批拒绝
POST   /api/v1/mdm/suppliers/:id/blacklist      # 加入黑名单
POST   /api/v1/mdm/suppliers/:id/remove-blacklist # 移出黑名单

# 供应商联系人
GET    /api/v1/mdm/suppliers/:id/contacts       # 联系人列表
POST   /api/v1/mdm/suppliers/:id/contacts       # 添加联系人
PUT    /api/v1/mdm/suppliers/:id/contacts/:contactId # 更新联系人
DELETE /api/v1/mdm/suppliers/:id/contacts/:contactId # 删除联系人

# 供应商银行账户
GET    /api/v1/mdm/suppliers/:id/bank-accounts # 银行账户列表
POST   /api/v1/mdm/suppliers/:id/bank-accounts  # 添加银行账户
PUT    /api/v1/mdm/suppliers/:id/bank-accounts/:accountId # 更新账户
DELETE /api/v1/mdm/suppliers/:id/bank-accounts/:accountId # 删除账户

# 供应商附件
GET    /api/v1/mdm/suppliers/:id/attachments    # 附件列表
POST   /api/v1/mdm/suppliers/:id/attachments    # 上传附件
DELETE /api/v1/mdm/suppliers/:id/attachments/:attachmentId # 删除附件

# 供应商物料关系
GET    /api/v1/mdm/suppliers/:id/materials      # 供应物料列表
POST   /api/v1/mdm/suppliers/:id/materials      # 添加供应物料
PUT    /api/v1/mdm/suppliers/:id/materials/:materialId # 更新供应物料
DELETE /api/v1/mdm/suppliers/:id/materials/:materialId # 删除供应物料

# 供应商查询
GET    /api/v1/mdm/suppliers/by-grade/:grade   # 按等级查询
GET    /api/v1/mdm/suppliers/blacklist         # 黑名单供应商
GET    /api/v1/mdm/suppliers/search            # 搜索供应商
```

### 3.2 客户管理API

```
# 客户主数据
GET    /api/v1/mdm/customers                   # 客户列表
POST   /api/v1/mdm/customers                    # 创建客户
GET    /api/v1/mdm/customers/:id               # 客户详情
PUT    /api/v1/mdm/customers/:id               # 更新客户
DELETE /api/v1/mdm/customers/:id               # 删除客户(仅PENDING)
POST   /api/v1/mdm/customers/:id/approve        # 审批通过
POST   /api/v1/mdm/customers/:id/reject        # 审批拒绝

# 客户联系人
GET    /api/v1/mdm/customers/:id/contacts       # 联系人列表
POST   /api/v1/mdm/customers/:id/contacts       # 添加联系人
PUT    /api/v1/mdm/customers/:id/contacts/:contactId # 更新联系人
DELETE /api/v1/mdm/customers/:id/contacts/:contactId # 删除联系人

# 收货地址
GET    /api/v1/mdm/customers/:id/delivery-addresses # 收货地址列表
POST   /api/v1/mdm/customers/:id/delivery-addresses # 添加收货地址
PUT    /api/v1/mdm/customers/:id/delivery-addresses/:addressId # 更新收货地址
DELETE /api/v1/mdm/customers/:id/delivery-addresses/:addressId # 删除收货地址
POST   /api/v1/mdm/customers/:id/delivery-addresses/:addressId/set-default # 设为默认

# 客户银行账户
GET    /api/v1/mdm/customers/:id/bank-accounts  # 银行账户列表
POST   /api/v1/mdm/customers/:id/bank-accounts  # 添加银行账户
PUT    /api/v1/mdm/customers/:id/bank-accounts/:accountId # 更新账户
DELETE /api/v1/mdm/customers/:id/bank-accounts/:accountId # 删除账户

# 客户查询
GET    /api/v1/mdm/customers/by-level/:level   # 按等级查询
GET    /api/v1/mdm/customers/key-customers     # 重点客户
GET    /api/v1/mdm/customers/search            # 搜索客户
```

### 3.3 SCP供应商采购信息API

```
# 采购信息
GET    /api/v1/scp/suppliers/:supplierId/purchase-info  # 获取采购信息
PUT    /api/v1/scp/suppliers/:supplierId/purchase-info  # 更新采购信息
POST   /api/v1/scp/suppliers/:supplierId/materials       # 添加供应物料关系
GET    /api/v1/scp/suppliers/:supplierId/materials      # 查询供应物料列表

# 供应商评级
GET    /api/v1/scp/suppliers/grades               # 供应商等级标准
PUT    /api/v1/scp/suppliers/grades/:id          # 更新等级标准

# 黑名单管理
GET    /api/v1/scp/suppliers/blacklist           # 黑名单列表
POST   /api/v1/scp/suppliers/:id/add-blacklist   # 加入黑名单
POST   /api/v1/scp/suppliers/:id/remove-blacklist # 移出黑名单
```

---

## 4. 业务规则

### 4.1 供应商等级标准

| 等级 | 综合评分 | 付款条款 | 优先级 |
|------|---------|---------|--------|
| A级 | 90-100分 | 月结30天 | 首选供应商 |
| B级 | 75-89分 | 月结45天 | 正常采购 |
| C级 | 60-74分 | 月结60天/现金 | 观察供应商 |
| D级 | <60分 | 现金/预付 | 整改或淘汰 |

### 4.2 客户信用控制

```go
// 信用检查
func (s *CustomerService) CheckCreditLimit(customerID int64, orderAmount float64) error {
    customer := s.GetCustomer(customerID)
    if customer.CreditUsed + orderAmount > customer.CreditLimit {
        return fmt.Errorf("超过信用额度 %.2f，可用额度 %.2f", 
            orderAmount, customer.CreditLimit - customer.CreditUsed)
    }
    return nil
}
```

### 4.3 供应商物料关系

1. 一个物料可以有多个供应商
2. 必须指定一个首选供应商
3. 采购订单默认从首选供应商采购
4. 替代供应商需要额外审批

---

## 5. 前端页面

| 页面 | 路由 | 说明 |
|------|------|------|
| 供应商列表 | /mdm/supplier-list | 供应商主数据管理 |
| 供应商详情 | /mdm/supplier-view/:id | 查看供应商详情 |
| 供应商编辑 | /mdm/supplier-edit/:id | 新建/编辑供应商 |
| 客户列表 | /mdm/customer-list | 客户主数据管理 |
| 客户详情 | /mdm/customer-view/:id | 查看客户详情 |
| 客户编辑 | /mdm/customer-edit/:id | 新建/编辑客户 |
| 黑名单管理 | /mdm/blacklist | 供应商/客户黑名单 |

---

*文档版本: V1.0 | 创建日期: 2026-04-09*
