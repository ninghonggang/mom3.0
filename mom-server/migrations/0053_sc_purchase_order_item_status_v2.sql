-- ============================================================
-- 0053 / 2026-07-09 / batch 3-1
-- 给 scp_purchase_order_item 加 status_v2 字段 + seed 4 字典码
-- 配套: internal/pkg/status/status.go 加 PurchaseOrderItem 字典
--       internal/model/scp.go 加 PurchaseOrderItem.StatusCode()/SetStatus()
--       internal/service/scp.go ReceivePurchaseOrderItem 切字典码
-- ============================================================

-- 1. 加 status_v2 字段(双轨,不动旧 status)
ALTER TABLE scp_purchase_order_item
    ADD COLUMN IF NOT EXISTS status_v2 VARCHAR(30);

CREATE INDEX IF NOT EXISTS idx_scp_poi_status_v2
    ON scp_purchase_order_item (status_v2);

-- 2. seed mdm_status_dict 字典码(对应 status.go 的 PurchaseOrderItemPending/Partial/Completed/Cancelled)
--    列实际叫 label / element_type / is_terminal,不是 label_zh/label_en/is_active
--    mdm_status_dict 没有 (entity, code) UNIQUE 约束,改用 WHERE NOT EXISTS 幂等
INSERT INTO mdm_status_dict (domain, entity, code, label, element_type, is_terminal, sort_order, description)
SELECT v.domain, v.entity, v.code, v.label, v.element_type, v.is_terminal, v.sort_order, v.description
FROM (VALUES
    ('scp',    'scp_purchase_order_item', 'PENDING',   '待收货',   'info',    FALSE, 1, '采购订单行项目待收货'),
    ('scp',    'scp_purchase_order_item', 'PARTIAL',   '部分收货', 'warning', FALSE, 2, '采购订单行项目部分收货'),
    ('scp',    'scp_purchase_order_item', 'COMPLETED', '已收货',   'success', TRUE,  3, '采购订单行项目全部收货'),
    ('scp',    'scp_purchase_order_item', 'CANCELLED', '已取消',   'info',    TRUE,  4, '采购订单行项目已取消')
) AS v(domain, entity, code, label, element_type, is_terminal, sort_order, description)
WHERE NOT EXISTS (
    SELECT 1 FROM mdm_status_dict d
    WHERE d.entity = v.entity AND d.code = v.code
);

-- 3. 数据回填:把现有 legacy status → status_v2(空表或 0 行时 NOTICE 跳过)
DO $$
BEGIN
    UPDATE scp_purchase_order_item
       SET status_v2 = CASE status
           WHEN 'PENDING'   THEN 'PENDING'
           WHEN 'PARTIAL'   THEN 'PARTIAL'
           WHEN 'COMPLETED' THEN 'COMPLETED'
           WHEN 'CANCELLED' THEN 'CANCELLED'
           ELSE 'PENDING'
       END
     WHERE status_v2 IS NULL OR status_v2 = '';
    RAISE NOTICE 'scp_purchase_order_item status_v2 回填完成';
EXCEPTION WHEN OTHERS THEN
    RAISE NOTICE 'scp_purchase_order_item status_v2 回填跳过: %', SQLERRM;
END
$$;

-- 4. 回滚(用 DO 块逐表容错)
/*
ALTER TABLE scp_purchase_order_item DROP COLUMN IF EXISTS status_v2;
DROP INDEX IF EXISTS idx_scp_poi_status_v2;
DELETE FROM mdm_status_dict WHERE entity = 'scp_purchase_order_item';
*/
