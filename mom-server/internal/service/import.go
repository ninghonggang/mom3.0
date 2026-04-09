package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"mom-server/internal/model"
	"mom-server/internal/repository"

	"github.com/xuri/excelize/v2"
)

type ImportService struct {
	importTaskRepo  *repository.ImportTaskRepository
	materialRepo    *repository.MaterialRepository
	bomRepo         *repository.BOMRepository
	bomItemRepo     *repository.BOMItemRepository
}

func NewImportService(importTaskRepo *repository.ImportTaskRepository, materialRepo *repository.MaterialRepository, bomRepo *repository.BOMRepository, bomItemRepo *repository.BOMItemRepository) *ImportService {
	return &ImportService{
		importTaskRepo:  importTaskRepo,
		materialRepo:    materialRepo,
		bomRepo:         bomRepo,
		bomItemRepo:     bomItemRepo,
	}
}

// MaterialImportRow 物料导入行数据
type MaterialImportRow struct {
	MaterialCode  string  `json:"material_code"`
	MaterialName  string  `json:"material_name"`
	MaterialType  string  `json:"material_type"`
	Spec          string  `json:"spec"`
	Unit          string  `json:"unit"`
	Status        int     `json:"status"`
	RowNum        int     `json:"row_num"`
	Error         string  `json:"error,omitempty"`
}

// BOMItemImportRow BOM明细行
type BOMItemImportRow struct {
	MaterialCode string  `json:"material_code"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
	ScrapRate    float64 `json:"scrap_rate"`
}

// BOMImportRow BOM导入行
type BOMImportRow struct {
	BOMCode      string             `json:"bom_code"`
	BOMName      string             `json:"bom_name"`
	MaterialCode string             `json:"material_code"`
	Version      string             `json:"version"`
	EffDate      string             `json:"eff_date"`
	Status       string             `json:"status"`
	Items        []BOMItemImportRow `json:"items"`
}

// CreateImportTask 创建导入任务
func (s *ImportService) CreateImportTask(ctx context.Context, tenantID int64, importType, fileName, filePath, createdBy string) (*model.ImportTask, error) {
	taskNo := fmt.Sprintf("IMP%s%s", time.Now().Format("20060102150405"), uuid.New().String()[:8])

	task := &model.ImportTask{
		TenantID:   tenantID,
		TaskNo:     taskNo,
		ImportType: importType,
		FileName:  fileName,
		FilePath:  filePath,
		Status:    model.ImportStatusPending,
		CreatedBy: createdBy,
	}

	if err := s.importTaskRepo.Create(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to create import task: %w", err)
	}

	return task, nil
}

// GetImportTask 获取导入任务
func (s *ImportService) GetImportTask(ctx context.Context, id string) (*model.ImportTask, error) {
	taskID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errors.New("invalid task id")
	}

	task, err := s.importTaskRepo.GetByID(ctx, uint(taskID))
	if err != nil {
		return nil, fmt.Errorf("task not found: %w", err)
	}

	return task, nil
}

// ParseExcelFromPath 从文件路径解析Excel
func (s *ImportService) ParseExcelFromPath(filePath string) ([]MaterialImportRow, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}
	defer f.Close()

	// 获取第一个sheet
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read sheet: %w", err)
	}

	return s.parseExcelRows(rows)
}

// ParseExcel 解析Excel文件
func (s *ImportService) ParseExcel(file multipart.File) ([]MaterialImportRow, error) {
	// 打开上传的文件
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}
	defer f.Close()

	// 获取第一个sheet
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read sheet: %w", err)
	}

	return s.parseExcelRows(rows)
}

// parseExcelRows 解析Excel行数据的公共方法
func (s *ImportService) parseExcelRows(rows [][]string) ([]MaterialImportRow, error) {
	// 跳过表头，从第二行开始
	if len(rows) < 2 {
		return nil, errors.New("excel file is empty or has no data rows")
	}

	var results []MaterialImportRow
	// 表头映射
	headerMap := make(map[string]int)
	for colIdx, cell := range rows[0] {
		headerMap[cell] = colIdx
	}

	// 验证必要列
	requiredCols := []string{"物料编码", "物料名称", "物料类型", "单位"}
	for _, col := range requiredCols {
		if _, ok := headerMap[col]; !ok {
			return nil, fmt.Errorf("missing required column: %s", col)
		}
	}

	// 解析数据行
	for rowIdx, row := range rows[1:] {
		rowNum := rowIdx + 2
		item := MaterialImportRow{RowNum: rowNum}

		getCellValue := func(colName string) string {
			colIdx, ok := headerMap[colName]
			if !ok || colIdx >= len(row) {
				return ""
			}
			val := strings.TrimSpace(row[colIdx])
			return val
		}

		item.MaterialCode = getCellValue("物料编码")
		item.MaterialName = getCellValue("物料名称")
		item.MaterialType = getCellValue("物料类型")
		item.Spec = getCellValue("规格")
		item.Unit = getCellValue("单位")

		// 状态默认启用
		statusStr := getCellValue("状态")
		if statusStr == "禁用" || statusStr == "0" {
			item.Status = 0
		} else {
			item.Status = 1
		}

		results = append(results, item)
	}

	return results, nil
}

// ValidateMaterial 验证物料数据
func (s *ImportService) ValidateMaterial(ctx context.Context, row *MaterialImportRow) error {
	if row.MaterialCode == "" {
		return errors.New("物料编码不能为空")
	}
	if row.MaterialName == "" {
		return errors.New("物料名称不能为空")
	}
	if row.MaterialType == "" {
		return errors.New("物料类型不能为空")
	}
	if row.Unit == "" {
		return errors.New("单位不能为空")
	}

	// 验证物料类型
	validTypes := map[string]bool{"原材料": true, "半成品": true, "成品": true}
	if !validTypes[row.MaterialType] {
		return errors.New("物料类型必须是: 原材料、半成品、成品")
	}

	// 转换物料类型为代码
	typeMap := map[string]string{"原材料": "raw", "半成品": "semi", "成品": "finished"}
	row.MaterialType = typeMap[row.MaterialType]

	return nil
}

// ImportMaterials 批量导入物料
func (s *ImportService) ImportMaterials(ctx context.Context, taskID uint, rows []MaterialImportRow, tenantID int64) error {
	// 更新任务状态为处理中
	if err := s.importTaskRepo.UpdateStatus(ctx, taskID, model.ImportStatusProcessing, len(rows), 0, 0); err != nil {
		return fmt.Errorf("failed to update task status: %w", err)
	}

	var successCount, failCount int
	var failData []MaterialImportRow

	for i := range rows {
		row := &rows[i]

		// 验证数据
		if err := s.ValidateMaterial(ctx, row); err != nil {
			row.Error = err.Error()
			failData = append(failData, *row)
			failCount++
			continue
		}

		// 创建物料
		material := &model.Material{
			TenantID:     tenantID,
			MaterialCode: row.MaterialCode,
			MaterialName: row.MaterialName,
			MaterialType: row.MaterialType,
			Unit:         row.Unit,
			Status:       row.Status,
		}
		if row.Spec != "" {
			material.Spec = &row.Spec
		}

		if err := s.materialRepo.Create(ctx, material); err != nil {
			row.Error = fmt.Sprintf("创建失败: %v", err)
			failData = append(failData, *row)
			failCount++
			continue
		}

		successCount++
	}

	// 更新最终状态
	status := model.ImportStatusSuccess
	if failCount > 0 {
		status = model.ImportStatusFail
	}

	failDataJSON, _ := json.Marshal(failData)
	updates := map[string]interface{}{
		"status":        status,
		"total_rows":    len(rows),
		"success_rows":  successCount,
		"fail_rows":     failCount,
		"fail_data_json": failDataJSON,
		"completed_at":  time.Now(),
	}

	if err := s.importTaskRepo.Update(ctx, taskID, updates); err != nil {
		return fmt.Errorf("failed to update task result: %w", err)
	}

	return nil
}

// SaveUploadedFile 保存上传的文件
func (s *ImportService) SaveUploadedFile(file multipart.File, header *multipart.FileHeader, uploadDir string) (string, error) {
	// 创建上传目录
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// 生成文件名
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filePath := filepath.Join(uploadDir, fileName)

	// 创建目标文件
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return filePath, nil
}

// GenerateMaterialTemplate 生成物料导入模板
func (s *ImportService) GenerateMaterialTemplate() (*excelize.File, error) {
	f := excelize.NewFile()
	sheet := "物料导入模板"
	index, err := f.NewSheet(sheet)
	if err != nil {
		return nil, err
	}
	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")

	// 设置表头
	headers := []string{"物料编码", "物料名称", "物料类型", "规格", "单位", "状态"}
	for colIdx, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(colIdx+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// 设置表头样式
	style, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#CCCCCC"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	f.SetCellStyle(sheet, "A1", "F1", style)

	// 设置列宽
	f.SetColWidth(sheet, "A", "A", 15)
	f.SetColWidth(sheet, "B", "B", 25)
	f.SetColWidth(sheet, "C", "C", 12)
	f.SetColWidth(sheet, "D", "D", 15)
	f.SetColWidth(sheet, "E", "E", 10)
	f.SetColWidth(sheet, "F", "F", 10)

	// 添加示例数据
	exampleData := [][]interface{}{
		{"M001", "原材料示例", "原材料", "规格1", "KG", "启用"},
		{"M002", "半成品示例", "半成品", "规格2", "PCS", "启用"},
		{"M003", "成品示例", "成品", "规格3", "BOX", "禁用"},
	}
	for rowIdx, row := range exampleData {
		for colIdx, val := range row {
			cell, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+2)
			f.SetCellValue(sheet, cell, val)
		}
	}

	return f, nil
}

// ParseBOMExcel 解析BOM导入Excel
func (s *ImportService) ParseBOMExcel(filePath string) ([]BOMImportRow, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read sheet: %w", err)
	}

	return s.parseBOMExcelRows(rows)
}

// parseBOMExcelRows 解析BOM Excel行数据
func (s *ImportService) parseBOMExcelRows(rows [][]string) ([]BOMImportRow, error) {
	if len(rows) < 2 {
		return nil, errors.New("excel file is empty or has no data rows")
	}

	var results []BOMImportRow
	headerMap := make(map[string]int)
	for colIdx, cell := range rows[0] {
		headerMap[strings.TrimSpace(cell)] = colIdx
	}

	requiredCols := []string{"BOM编码", "BOM名称", "物料编码", "版本", "生效日期", "状态", "明细(JSON)"}
	for _, col := range requiredCols {
		if _, ok := headerMap[col]; !ok {
			return nil, fmt.Errorf("missing required column: %s", col)
		}
	}

	for rowIdx, row := range rows[1:] {
		rowNum := rowIdx + 2
		item := BOMImportRow{}

		getCellValue := func(colName string) string {
			colIdx, ok := headerMap[colName]
			if !ok || colIdx >= len(row) {
				return ""
			}
			return strings.TrimSpace(row[colIdx])
		}

		item.BOMCode = getCellValue("BOM编码")
		item.BOMName = getCellValue("BOM名称")
		item.MaterialCode = getCellValue("物料编码")
		item.Version = getCellValue("版本")
		item.EffDate = getCellValue("生效日期")
		item.Status = getCellValue("状态")
		if item.Status == "" {
			item.Status = "DRAFT"
		}

		// 解析明细JSON
		itemsJSON := getCellValue("明细(JSON)")
		if itemsJSON != "" {
			var items []BOMItemImportRow
			if err := json.Unmarshal([]byte(itemsJSON), &items); err == nil {
				item.Items = items
			}
		}

		if item.BOMCode == "" {
			return nil, fmt.Errorf("第%d行: BOM编码不能为空", rowNum)
		}
		if item.MaterialCode == "" {
			return nil, fmt.Errorf("第%d行: 物料编码不能为空", rowNum)
		}

		results = append(results, item)
	}

	return results, nil
}

// ImportBOMs 导入BOM数据
func (s *ImportService) ImportBOMs(ctx context.Context, taskID uint, rows []BOMImportRow, tenantID int64) error {
	// 更新任务状态为处理中
	if err := s.importTaskRepo.UpdateStatus(ctx, taskID, model.ImportStatusProcessing, len(rows), 0, 0); err != nil {
		return fmt.Errorf("failed to update task status: %w", err)
	}

	var successCount, failCount int
	var failData []BOMImportRow

	for i := range rows {
		row := &rows[i]

		// 验证BOM数据
		if err := s.ValidateBOM(ctx, row); err != nil {
			failData = append(failData, *row)
			failCount++
			continue
		}

		// 查找物料ID
		material, err := s.materialRepo.GetByCode(ctx, row.MaterialCode)
		if err != nil {
			failData = append(failData, *row)
			failCount++
			continue
		}

		// 解析日期
		var effDate *time.Time
		if row.EffDate != "" {
			t, err := time.Parse("2006-01-02", row.EffDate)
			if err == nil {
				effDate = &t
			}
		}

		// 创建BOM
		bom := &model.MdmBOM{
			TenantID:      tenantID,
			BOMCode:      row.BOMCode,
			BOMName:      row.BOMName,
			MaterialID:   material.ID,
			MaterialCode: material.MaterialCode,
			MaterialName: material.MaterialName,
			Version:      row.Version,
			Status:       row.Status,
			EffDate:      effDate,
			IsCurrent:    1,
		}
		if err := s.bomRepo.Create(ctx, bom); err != nil {
			failData = append(failData, *row)
			failCount++
			continue
		}

		// 创建BOM明细
		for lineNo, item := range row.Items {
			// 查找子物料ID
			subMat, err := s.materialRepo.GetByCode(ctx, item.MaterialCode)
			if err != nil {
				failData = append(failData, *row)
				failCount++
				continue
			}
			bomItem := &model.MdmBOMItem{
				TenantID:     tenantID,
				BOMID:        bom.ID,
				LineNo:       lineNo + 1,
				MaterialID:   subMat.ID,
				MaterialCode: subMat.MaterialCode,
				MaterialName: subMat.MaterialName,
				Quantity:     item.Quantity,
				Unit:         item.Unit,
				ScrapRate:    item.ScrapRate,
			}
			if err := s.bomItemRepo.Create(ctx, bomItem); err != nil {
				failData = append(failData, *row)
				failCount++
				continue
			}
		}

		successCount++
	}

	// 更新最终状态
	status := model.ImportStatusSuccess
	if failCount > 0 {
		status = model.ImportStatusFail
	}

	failDataJSON, _ := json.Marshal(failData)
	updates := map[string]interface{}{
		"status":         status,
		"total_rows":     len(rows),
		"success_rows":   successCount,
		"fail_rows":      failCount,
		"fail_data_json": failDataJSON,
		"completed_at":   time.Now(),
	}

	if err := s.importTaskRepo.Update(ctx, taskID, updates); err != nil {
		return fmt.Errorf("failed to update task result: %w", err)
	}

	return nil
}

// ValidateBOM 验证BOM数据
func (s *ImportService) ValidateBOM(ctx context.Context, row *BOMImportRow) error {
	if row.BOMCode == "" {
		return errors.New("BOM编码不能为空")
	}
	if row.BOMName == "" {
		return errors.New("BOM名称不能为空")
	}
	if row.MaterialCode == "" {
		return errors.New("物料编码不能为空")
	}
	return nil
}

// GenerateBOMTemplate 生成BOM导入模板
func (s *ImportService) GenerateBOMTemplate() (*excelize.File, error) {
	f := excelize.NewFile()
	sheet := "BOM导入模板"
	s1, _ := f.NewSheet(sheet)
	f.SetSheetName("Sheet1", sheet)
	// 设置表头
	headers := []string{"BOM编码", "BOM名称", "物料编码", "版本", "生效日期", "状态", "明细(JSON)"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	// 设置示例数据说明
	f.SetCellValue(sheet, "G2", `[{"material_code":"MAT-001","quantity":1,"unit":"件","scrap_rate":0}]`)
	f.SetActiveSheet(s1)
	return f, nil
}
