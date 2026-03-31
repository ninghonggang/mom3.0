package model

// Supplier 供应商
type Supplier struct {
	BaseModel
	TenantID    int64   `json:"tenant_id" gorm:"index;not null"`
	Code        string  `json:"code" gorm:"size:50;not null;uniqueIndex:idx_tenant_supplier_code"` // 供应商编码
	Name        string  `json:"name" gorm:"size:100;not null"`                                      // 供应商名称
	Type        string  `json:"type" gorm:"size:20"`                                               // 供应商类型: 原材料/辅料/设备/服务
	Contact     *string `json:"contact" gorm:"size:100"`                                           // 联系人
	Phone       *string `json:"phone" gorm:"size:20"`                                              // 联系电话
	Email       *string `json:"email" gorm:"size:100"`                                            // 邮箱
	Address     *string `json:"address" gorm:"size:200"`                                           // 地址
	Category    *string `json:"category" gorm:"size:50"`                                          // 物料类别(关联的物料类型)
	Level       int     `json:"level" gorm:"default:1"`                                            // 供应商等级: 1=A级/2=B级/3=C级
	Status      int     `json:"status" gorm:"default:1"`                                          // 状态: 1=启用/2=禁用
	Remark      *string `json:"remark" gorm:"size:500"`                                            // 备注
}

func (Supplier) TableName() string {
	return "mdm_supplier"
}
