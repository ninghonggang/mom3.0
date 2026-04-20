package model

import "time"

// EquipmentDocument 设备文档
type EquipmentDocument struct {
	BaseModel
	TenantID       int64   `json:"tenant_id" gorm:"index;not null"`
	EquipmentID    int64   `json:"equipment_id" gorm:"index;not null"`
	EquipmentCode  string  `json:"equipment_code" gorm:"size:50"`
	EquipmentName  string  `json:"equipment_name" gorm:"size:100"`
	DocType        string  `json:"doc_type" gorm:"size:30"`  // 文档类型: MANUAL/SPEC/CERT/DRAWING/REPORT/OTHER
	DocName        string  `json:"doc_name" gorm:"size:200;not null"` // 文档名称
	DocCode        string  `json:"doc_code" gorm:"size:50"`  // 文档编号
	FileName       string  `json:"file_name" gorm:"size:200"` // 文件名
	FilePath       string  `json:"file_path" gorm:"size:500"` // 文件路径
	FileSize       int64   `json:"file_size"`                // 文件大小(字节)
	FileType       string  `json:"file_type" gorm:"size:50"` // 文件类型
	FileURL        *string `json:"file_url" gorm:"size:500"` // 文件URL
	Version        string  `json:"version" gorm:"size:20"`   // 版本
	EffectiveDate  *time.Time `json:"effective_date"`         // 生效日期
	ExpiryDate     *time.Time `json:"expiry_date"`           // 失效日期
	Description    *string `json:"description" gorm:"size:500"` // 描述
	UploadedBy     *int64  `json:"uploaded_by"`
	UploadedByName *string `json:"uploaded_by_name" gorm:"size:50"`
	UploadTime     *time.Time `json:"upload_time"`
	Status         int     `json:"status" gorm:"default:1"`  // 1有效/2失效
}

func (EquipmentDocument) TableName() string {
	return "equ_equipment_document"
}

// EquipmentDocumentCreate 设备文档创建
type EquipmentDocumentCreate struct {
	EquipmentID    int64  `json:"equipment_id" binding:"required"`
	DocType       string `json:"doc_type" binding:"required"`
	DocName       string `json:"doc_name" binding:"required"`
	DocCode       string `json:"doc_code"`
	FileName      string `json:"file_name"`
	FilePath      string `json:"file_path"`
	FileSize      int64  `json:"file_size"`
	FileType      string `json:"file_type"`
	FileURL       *string `json:"file_url"`
	Version       string `json:"version"`
	EffectiveDate *time.Time `json:"effective_date"`
	ExpiryDate    *time.Time `json:"expiry_date"`
	Description   *string `json:"description"`
}

type EquipmentDocumentUpdate struct {
	DocType       *string `json:"doc_type"`
	DocName       *string `json:"doc_name"`
	DocCode       *string `json:"doc_code"`
	FileName      *string `json:"file_name"`
	FilePath      *string `json:"file_path"`
	FileSize      *int64  `json:"file_size"`
	FileType      *string `json:"file_type"`
	FileURL       *string `json:"file_url"`
	Version       *string `json:"version"`
	EffectiveDate *time.Time `json:"effective_date"`
	ExpiryDate    *time.Time `json:"expiry_date"`
	Description   *string `json:"description"`
	Status        *int    `json:"status"`
}
