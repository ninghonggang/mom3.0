package model

import "time"

type Material struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID     string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	MaterialCode string    `gorm:"column:material_code;type:varchar(64);uniqueIndex;not null" json:"material_code"`
	MaterialName string    `gorm:"column:material_name;type:varchar(128);not null" json:"material_name"`
	MaterialType string    `gorm:"column:material_type;type:varchar(32);not null" json:"material_type"` // RAW/SEMIFINISHED/FINISHED/PACKAGING
	Spec         string    `gorm:"column:spec;type:varchar(256)" json:"spec"`
	Unit         string    `gorm:"column:unit;type:varchar(64)" json:"unit"`
	UnitWeight   float64   `gorm:"column:unit_weight;type:decimal(18,4)" json:"unit_weight"`
	Description  string    `gorm:"column:description;type:varchar(512)" json:"description"`
	Status       string    `gorm:"column:status;type:varchar(64);default:ACTIVE" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Material) TableName() string { return "material" }

type Bom struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID      string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	BomCode       string    `gorm:"column:bom_code;type:varchar(64);uniqueIndex;not null" json:"bom_code"`
	MaterialID    uint64    `gorm:"column:material_id;index;not null" json:"material_id"`
	Version       string    `gorm:"column:version;type:varchar(64);not null" json:"version"`
	EffectiveDate time.Time `gorm:"column:effective_date" json:"effective_date"`
	ExpiryDate    time.Time `gorm:"column:expiry_date" json:"expiry_date"`
	Status        string    `gorm:"column:status;type:varchar(64);default:DRAFT" json:"status"` // DRAFT/ACTIVE/OBSOLETE
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Items []BomItem `gorm:"foreignKey:BomID" json:"items,omitempty"`
}

func (Bom) TableName() string { return "bom" }

type BomItem struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	BomID           uint64    `gorm:"column:bom_id;index;not null" json:"bom_id"`
	ChildMaterialID uint64    `gorm:"column:child_material_id;not null" json:"child_material_id"`
	Quantity        float64   `gorm:"column:quantity;type:decimal(18,4);not null" json:"quantity"`
	Unit            string    `gorm:"column:unit;type:varchar(64)" json:"unit"`
	Position        string    `gorm:"column:position;type:varchar(32)" json:"position"`
	ScrapRate       float64   `gorm:"column:scrap_rate;type:decimal(5,2);default:0" json:"scrap_rate"`
	IsKeyPart       bool      `gorm:"column:is_key_part;default:false" json:"is_key_part"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (BomItem) TableName() string { return "bom_item" }

type Workshop struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID      string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	WorkshopCode  string    `gorm:"column:workshop_code;type:varchar(64);uniqueIndex;not null" json:"workshop_code"`
	WorkshopName  string    `gorm:"column:workshop_name;type:varchar(128);not null" json:"workshop_name"`
	FactoryID     uint64    `gorm:"column:factory_id" json:"factory_id"`
	Area          float64   `gorm:"column:area;type:decimal(12,2)" json:"area"`
	ManagerID     string    `gorm:"column:manager_id;type:varchar(64)" json:"manager_id"`
	Status        string    `gorm:"column:status;type:varchar(64);default:ACTIVE" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Workshop) TableName() string { return "workshop" }

type ProductionLine struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID  string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	LineCode  string    `gorm:"column:line_code;type:varchar(64);uniqueIndex;not null" json:"line_code"`
	LineName  string    `gorm:"column:line_name;type:varchar(128);not null" json:"line_name"`
	WorkshopID uint64   `gorm:"column:workshop_id;index;not null" json:"workshop_id"`
	LineType  string    `gorm:"column:line_type;type:varchar(32)" json:"line_type"`
	Capacity  float64   `gorm:"column:capacity;type:decimal(12,2)" json:"capacity"`
	Status    string    `gorm:"column:status;type:varchar(64);default:ACTIVE" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (ProductionLine) TableName() string { return "production_line" }

type Workstation struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID        string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	WorkstationCode string    `gorm:"column:workstation_code;type:varchar(64);uniqueIndex;not null" json:"workstation_code"`
	WorkstationName string    `gorm:"column:workstation_name;type:varchar(128);not null" json:"workstation_name"`
	LineID          uint64    `gorm:"column:line_id;index" json:"line_id"`
	WorkshopID      uint64    `gorm:"column:workshop_id;index" json:"workshop_id"`
	WorkstationType string    `gorm:"column:workstation_type;type:varchar(32)" json:"workstation_type"`
	Status          string    `gorm:"column:status;type:varchar(64);default:ACTIVE" json:"status"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Workstation) TableName() string { return "workstation" }

type Customer struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID      string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	CustomerCode  string    `gorm:"column:customer_code;type:varchar(64);uniqueIndex;not null" json:"customer_code"`
	CustomerName  string    `gorm:"column:customer_name;type:varchar(128);not null" json:"customer_name"`
	CustomerType  string    `gorm:"column:customer_type;type:varchar(32)" json:"customer_type"`
	ContactPerson string    `gorm:"column:contact_person;type:varchar(64)" json:"contact_person"`
	ContactPhone  string    `gorm:"column:contact_phone;type:varchar(32)" json:"contact_phone"`
	Address       string    `gorm:"column:address;type:varchar(256)" json:"address"`
	Status        string    `gorm:"column:status;type:varchar(64);default:ACTIVE" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Customer) TableName() string { return "customer" }

type Supplier struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	TenantID      string    `gorm:"column:tenant_id;type:varchar(64);index;not null" json:"tenant_id"`
	SupplierCode  string    `gorm:"column:supplier_code;type:varchar(64);uniqueIndex;not null" json:"supplier_code"`
	SupplierName  string    `gorm:"column:supplier_name;type:varchar(128);not null" json:"supplier_name"`
	SupplierType  string    `gorm:"column:supplier_type;type:varchar(32)" json:"supplier_type"`
	ContactPerson string    `gorm:"column:contact_person;type:varchar(64)" json:"contact_person"`
	ContactPhone  string    `gorm:"column:contact_phone;type:varchar(32)" json:"contact_phone"`
	Address       string    `gorm:"column:address;type:varchar(256)" json:"address"`
	Status        string    `gorm:"column:status;type:varchar(64);default:ACTIVE" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Supplier) TableName() string { return "supplier" }
