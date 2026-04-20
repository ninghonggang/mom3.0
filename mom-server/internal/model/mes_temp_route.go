package model

// TempRoute MES日计划临时工艺路线
type TempRoute struct {
	BaseModel
	TenantID     int64  `json:"tenant_id" gorm:"index;not null"`
	OrderDayID   int64  `json:"order_day_id" gorm:"index;not null"`
	TempRouteName string `json:"temp_route_name" gorm:"size:200;not null"`
	RouteContent string `json:"route_content" gorm:"type:text"`
	Reason       string `json:"reason" gorm:"size:500"`
	Status       int    `json:"status" gorm:"default:0"` // 0=pending待审核, 1=approved已批准, 2=rejected已拒绝
	Creator      string `json:"creator" gorm:"size:64"`
}

func (TempRoute) TableName() string {
	return "plan_mes_order_day_temp_route"
}

// TempRouteCreate 创建临时工艺路线请求
type TempRouteCreate struct {
	OrderDayID    int64  `json:"order_day_id" binding:"required"`
	TempRouteName string `json:"temp_route_name" binding:"required"`
	RouteContent  string `json:"route_content"`
	Reason        string `json:"reason"`
}

// TempRouteUpdate 更新临时工艺路线请求
type TempRouteUpdate struct {
	TempRouteName string `json:"temp_route_name"`
	RouteContent  string `json:"route_content"`
	Reason        string `json:"reason"`
}

// TempRouteApprove 审核临时工艺路线请求
type TempRouteApprove struct {
	ID      int64  `json:"id" binding:"required"`
	Status  int    `json:"status" binding:"required"` // 1=approved, 2=rejected
	Comment string `json:"comment"`
}
