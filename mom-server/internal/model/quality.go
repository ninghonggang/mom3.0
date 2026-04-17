package model

import (
	"time"
)

// ========== 质量管理模块 ==========

// IQC IQC检验(来料检验)
type IQC struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	IQCNo        string     `json:"iqc_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_iqc"`
	SupplierID   int64      `json:"supplier_id"`
	SupplierName *string    `json:"supplier_name" gorm:"size:100"`
	MaterialID   int64      `json:"material_id"`
	MaterialCode string     `json:"material_code" gorm:"size:50"`
	MaterialName string     `json:"material_name" gorm:"size:100"`
	Quantity     float64    `json:"quantity" gorm:"type:decimal(18,4)"`
	Unit         string     `json:"unit" gorm:"size:20"`
	CheckUserID  int64      `json:"check_user_id"`
	CheckUserName *string   `json:"check_user_name" gorm:"size:50"`
	CheckDate    *time.Time `json:"check_date"`
	Result       int        `json:"result" gorm:"default:1"` // 1待检验/2合格/3不合格
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (IQC) TableName() string {
	return "qc_iqc"
}

// IQCItem IQC检验明细
type IQCItem struct {
	BaseModel
	IQCID       int64   `json:"iqc_id" gorm:"index"`
	CheckItemID int64   `json:"check_item_id"`
	CheckItem   string  `json:"check_item" gorm:"size:100"`
	CheckStandard string `json:"check_standard" gorm:"size:200"` // 检查标准
	CheckMethod *string `json:"check_method" gorm:"size:100"` // 检查方法
	Result      int      `json:"result" gorm:"default:1"` // 1合格/2不合格
	Remark     *string `json:"remark" gorm:"size:200"`
}

func (IQCItem) TableName() string {
	return "qc_iqc_item"
}

// IPQC 过程检验
type IPQC struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	IPQCNo       string     `json:"ipqc_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_ipqc"`
	OrderID      int64      `json:"order_id"` // 生产工单ID
	OrderNo      string     `json:"order_no" gorm:"size:50"`
	ProcessID    int64      `json:"process_id"`
	ProcessName  *string    `json:"process_name" gorm:"size:100"`
	Quantity     float64    `json:"quantity" gorm:"type:decimal(18,4)"`
	SampleSize   int        `json:"sample_size"` // 抽样数
	CheckUserID  int64      `json:"check_user_id"`
	CheckUserName *string   `json:"check_user_name" gorm:"size:50"`
	CheckDate    *time.Time `json:"check_date"`
	Result       int        `json:"result" gorm:"default:1"` // 1待检验/2合格/3不合格
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (IPQC) TableName() string {
	return "qc_ipqc"
}

// FQC 最终检验
type FQC struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	FQCNo        string     `json:"fqc_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_fqc"`
	OrderID      int64      `json:"order_id"`
	OrderNo      string     `json:"order_no" gorm:"size:50"`
	Quantity     float64    `json:"quantity" gorm:"type:decimal(18,4)"`
	SampleSize   int        `json:"sample_size"`
	QualifiedQty float64    `json:"qualified_qty" gorm:"type:decimal(18,4)"`
	RejectedQty  float64    `json:"rejected_qty" gorm:"type:decimal(18,4);default:0"`
	CheckUserID  int64      `json:"check_user_id"`
	CheckUserName *string   `json:"check_user_name" gorm:"size:50"`
	CheckDate    *time.Time `json:"check_date"`
	Result       int        `json:"result" gorm:"default:1"`
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (FQC) TableName() string {
	return "qc_fqc"
}

// OQC 出货检验
type OQC struct {
	BaseModel
	TenantID     int64      `json:"tenant_id" gorm:"index;not null"`
	OQCNo        string     `json:"oqc_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_oqc"`
	ShippingNo   string     `json:"shipping_no" gorm:"size:50"` // 发货单号
	CustomerID   int64      `json:"customer_id"`
	CustomerName *string    `json:"customer_name" gorm:"size:100"`
	Quantity     float64    `json:"quantity" gorm:"type:decimal(18,4)"`
	CheckUserID  int64      `json:"check_user_id"`
	CheckUserName *string   `json:"check_user_name" gorm:"size:50"`
	CheckDate    *time.Time `json:"check_date"`
	Result       int        `json:"result" gorm:"default:1"`
	Remark       *string    `json:"remark" gorm:"size:500"`
}

func (OQC) TableName() string {
	return "qc_oqc"
}

// DefectCode 不良品代码
type DefectCode struct {
	BaseModel
	TenantID    int64   `json:"tenant_id" gorm:"index;not null"`
	DefectCode  string  `json:"defect_code" gorm:"size:20;not null;uniqueIndex:idx_tenant_defect"`
	DefectName  string  `json:"defect_name" gorm:"size:100;not null"`
	DefectType  string  `json:"defect_type" gorm:"size:20"` // 尺寸/外观/功能/其他
	Severity    int     `json:"severity" gorm:"default:1"` // 1轻微/2一般/3严重
	Status      int     `json:"status" gorm:"default:1"`
}

func (DefectCode) TableName() string {
	return "qc_defect_code"
}

// DefectRecord 不良品记录
type DefectRecord struct {
	BaseModel
	TenantID      int64      `json:"tenant_id" gorm:"index;not null"`
	RecordNo      string     `json:"record_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_defect_rec"`
	OrderID       int64      `json:"order_id"`
	OrderNo       string     `json:"order_no" gorm:"size:50"`
	ProcessID     int64      `json:"process_id"`
	ProcessName   *string    `json:"process_name" gorm:"size:100"`
	DefectCodeID  int64      `json:"defect_code_id"`
	DefectCode    string     `json:"defect_code" gorm:"size:20"`
	DefectName    string     `json:"defect_name" gorm:"size:100"`
	Quantity      float64    `json:"quantity" gorm:"type:decimal(18,4)"`
	HandleMethod  int        `json:"handle_method" gorm:"default:1"` // 1返工/2返修/3报废/4特采
	HandleUserID  *int64     `json:"handle_user_id"`
	HandleDate    *time.Time `json:"handle_date"`
	Status        int        `json:"status" gorm:"default:1"` // 1待处理/2处理中/3已处理
}

func (DefectRecord) TableName() string {
	return "qc_defect_record"
}

// NCR 不良品处理单
type NCR struct {
	BaseModel
	TenantID    int64      `json:"tenant_id" gorm:"index;not null"`
	NCRNo       string     `json:"ncr_no" gorm:"size:50;not null;uniqueIndex:idx_tenant_ncr"`
	DefectID    int64      `json:"defect_id"` // 不良品记录ID
	SourceType  string     `json:"source_type" gorm:"size:20"` // IQC/IPQC/FQC/OQC
	IssueDesc   string     `json:"issue_desc" gorm:"type:text"` // 问题描述
	RootCause   *string    `json:"root_cause" gorm:"type:text"` // 根本原因
	CorrectiveAction *string `json:"corrective_action" gorm:"type:text"` // 纠正措施
	PreventiveAction *string `json:"preventive_action" gorm:"type:text"` // 预防措施
	VerifyResult *string    `json:"verify_result" gorm:"size:200"` // 验证结果
	VerifyUserID *int64    `json:"verify_user_id"`
	VerifyDate   *time.Time `json:"verify_date"`
	Status       int        `json:"status" gorm:"default:1"` // 1待处理/2处理中/3已完成/4已关闭
}

func (NCR) TableName() string {
	return "qc_ncr"
}

// SPCData SPC数据
type SPCData struct {
	BaseModel
	TenantID    int64      `json:"tenant_id" gorm:"index;not null"`
	EquipmentID int64      `json:"equipment_id"`
	StationID   int64      `json:"station_id"`
	ProcessID   int64      `json:"process_id"`
	ProcessName *string    `json:"process_name" gorm:"size:100"`
	CheckItem   string     `json:"check_item" gorm:"size:100"`
	CheckValue  float64    `json:"check_value" gorm:"type:decimal(18,4)"`
	USL         *float64   `json:"usl"` // 规格上限
	LSL         *float64   `json:"lsl"` // 规格下限
	CL          *float64   `json:"cl"`   // 中心值
	UCL         *float64   `json:"ucl"` // 控制上限
	LCL         *float64   `json:"lcl"` // 控制下限
	CheckTime   time.Time  `json:"check_time"`
}

func (SPCData) TableName() string {
	return "qc_spc_data"
}

// InspectionCharacteristic 检验特性
type InspectionCharacteristic struct {
	ID               uint    `json:"id" gorm:"primarykey"`
	Code             string  `json:"code" gorm:"size:50;uniqueIndex;not null"`
	Name             string  `json:"name" gorm:"size:100;not null"`
	Type             string  `json:"type" gorm:"size:20"`                      // QUANTITATIVE/QUALITATIVE
	SpecLower        float64 `json:"spec_lower" gorm:"type:decimal(18,4)"`
	SpecUpper        float64 `json:"spec_upper" gorm:"type:decimal(18,4)"`
	USL              float64 `json:"usl" gorm:"type:decimal(18,4)"`
	LSL              float64 `json:"lsl" gorm:"type:decimal(18,4)"`
	Target           float64 `json:"target" gorm:"type:decimal(18,4)"`
	Unit             string  `json:"unit" gorm:"size:20"`
	AQL              float64 `json:"aql" gorm:"type:decimal(5,2)"`
	InspectionMethod string  `json:"inspection_method" gorm:"size:100"`
	TenantID         int64   `json:"tenant_id" gorm:"index"`
	Status           int     `json:"status" gorm:"default:1"`
}

func (InspectionCharacteristic) TableName() string {
	return "qc_inspection_characteristic"
}

// AQLLevel AQL级别定义
type AQLLevel struct {
	BaseModel
	TenantID int64  `json:"tenant_id" gorm:"index;not null"`
	Level   string `json:"level" gorm:"size:10;not null"`
	Name    string `json:"name" gorm:"size:50;not null"`
	Type    string `json:"type" gorm:"size:20"`
	Order   int    `json:"order" gorm:"default:0"`
	Status  int    `json:"status" gorm:"default:1"`
	Remark  *string `json:"remark" gorm:"size:500"`
}

func (AQLLevel) TableName() string {
	return "qc_aql_level"
}

// AQLTableRow AQL标准表行
type AQLTableRow struct {
	BaseModel
	TenantID   int64  `json:"tenant_id" gorm:"index;not null"`
	AQLLevelID int64  `json:"aql_level_id" gorm:"index;not null"`
	AQLValue   string `json:"aql_value" gorm:"size:10"`
	BatchMin   int    `json:"batch_min"`
	BatchMax   int    `json:"batch_max"`
	SampleSize int    `json:"sample_size"`
	Ac         int    `json:"ac"`
	Re         int    `json:"re"`
}

func (AQLTableRow) TableName() string {
	return "qc_aql_table_row"
}

// SamplingPlan 抽样方案
type SamplingPlan struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	Code           string  `json:"code" gorm:"size:50;not null;uniqueIndex:idx_tenant_plan_code"`
	Name           string  `json:"name" gorm:"size:100;not null"`
	InspectionType string  `json:"inspection_type" gorm:"size:20"`
	AQLLevelID    int64   `json:"aql_level_id"`
	DefaultAQL    float64 `json:"default_aql" gorm:"type:decimal(5,2)"`
	MinBatchSize  int     `json:"min_batch_size" gorm:"default:0"`
	MaxBatchSize  int     `json:"max_batch_size" gorm:"default:0"`
	Status        int     `json:"status" gorm:"default:1"`
	Remark        *string `json:"remark" gorm:"size:500"`
}

func (SamplingPlan) TableName() string {
	return "qc_sampling_plan"
}
