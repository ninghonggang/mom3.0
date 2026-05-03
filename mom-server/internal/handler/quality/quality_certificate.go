package quality

import (
	"strconv"

	"mom-server/internal/middleware"
	"mom-server/internal/model"
	"mom-server/internal/pkg/response"
	"mom-server/internal/service"

	"github.com/gin-gonic/gin"
)

type QualityCertificateHandler struct {
	svc *service.QualityCertificateService
}

func NewQualityCertificateHandler(svc *service.QualityCertificateService) *QualityCertificateHandler {
	return &QualityCertificateHandler{svc: svc}
}

// List 获取质量证书列表(分页)
// GET /quality/certificate/page
func (h *QualityCertificateHandler) List(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	orderCode := c.Query("order_code")
	productCode := c.Query("product_code")
	productName := c.Query("product_name")
	batchNo := c.Query("batch_no")
	certType := c.Query("cert_type")
	result, _ := strconv.Atoi(c.Query("result"))
	status, _ := strconv.Atoi(c.Query("status"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	req := &model.QualityCertificateQuery{
		OrderCode:   orderCode,
		ProductCode: productCode,
		ProductName: productName,
		BatchNo:     batchNo,
		CertType:    certType,
		Result:      result,
		Status:      status,
		StartDate:   startDate,
		EndDate:     endDate,
		Page:        page,
		PageSize:    pageSize,
	}

	list, total, err := h.svc.List(c.Request.Context(), tenantID, req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, gin.H{"list": list, "total": total, "page": page, "page_size": pageSize})
}

// Get 获取单个质量证书
// GET /quality/certificate/:id
func (h *QualityCertificateHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	cert, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, cert)
}

// Create 创建质量证书
// POST /quality/certificate
func (h *QualityCertificateHandler) Create(c *gin.Context) {
	tenantID := middleware.GetTenantID(c)
	if tenantID <= 0 {
		tenantID = 1
	}
	var req model.QualityCertificateCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	cert, err := h.svc.Create(c.Request.Context(), tenantID, &req)
	if err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, cert)
}

// Update 更新质量证书
// PUT /quality/certificate/:id
func (h *QualityCertificateHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	var req model.QualityCertificateCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}

// Delete 删除质量证书
// DELETE /quality/certificate/:id
func (h *QualityCertificateHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorMsg(c, err.Error())
		return
	}
	response.Success(c, nil)
}