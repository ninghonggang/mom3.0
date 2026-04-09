package service

import (
	"context"
	"fmt"
	"math"
	"mom-server/internal/model"
	"mom-server/internal/repository"
)

type SPCDataService struct {
	repo *repository.SPCDataRepository
}

func NewSPCDataService(repo *repository.SPCDataRepository) *SPCDataService {
	return &SPCDataService{repo: repo}
}

func (s *SPCDataService) List(ctx context.Context) ([]model.SPCData, int64, error) {
	return s.repo.List(ctx, 1)
}

func (s *SPCDataService) GetByID(ctx context.Context, id string) (*model.SPCData, error) {
	var spcID uint
	_, err := fmt.Sscanf(id, "%d", &spcID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, spcID)
}

func (s *SPCDataService) Create(ctx context.Context, spcData *model.SPCData) error {
	spcData.TenantID = 1
	return s.repo.Create(ctx, spcData)
}

func (s *SPCDataService) Update(ctx context.Context, id string, spcData *model.SPCData) error {
	var spcID uint
	_, err := fmt.Sscanf(id, "%d", &spcID)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"check_value": spcData.CheckValue,
		"usl":         spcData.USL,
		"lsl":         spcData.LSL,
		"cl":          spcData.CL,
		"ucl":         spcData.UCL,
		"lcl":         spcData.LCL,
	}
	return s.repo.Update(ctx, spcID, updates)
}

func (s *SPCDataService) Delete(ctx context.Context, id string) error {
	var spcID uint
	_, err := fmt.Sscanf(id, "%d", &spcID)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, spcID)
}

type SPCChartQuery struct {
	EquipmentID int64
	ProcessID   int64
	StationID   int64
	CheckItem   string
	Limit       int
}

func (s *SPCDataService) GetChartData(ctx context.Context, query SPCChartQuery) ([]model.SPCData, error) {
	if query.Limit <= 0 {
		query.Limit = 100
	}
	return s.repo.GetChartData(ctx, 1, query.EquipmentID, query.ProcessID, query.StationID, query.CheckItem, query.Limit)
}

// SPCStats SPC统计结果
type SPCStats struct {
	Mean   float64 `json:"mean"`
	StdDev float64 `json:"std_dev"`
	Cp     float64 `json:"cp"`
	Cpk    float64 `json:"cpk"`
	Ca     float64 `json:"ca"`
	Count  int     `json:"count"`
}

// CalculateCPK 计算Cp/Cpk
func CalculateCPK(values []float64, usl, lsl float64) SPCStats {
	n := len(values)
	if n == 0 {
		return SPCStats{}
	}

	// 计算均值
	var sum, sumSq float64
	for _, v := range values {
		sum += v
		sumSq += v * v
	}
	mean := sum / float64(n)

	// 计算标准差 (总体标准差)
	variance := (sumSq / float64(n)) - (mean * mean)
	if variance < 0 {
		variance = 0
	}
	stdDev := math.Sqrt(variance)

	// Cp = (USL - LSL) / (6 * sigma)
	tolerance := usl - lsl
	var cp float64
	if stdDev > 0 {
		cp = tolerance / (6 * stdDev)
	}

	// Ca = (Mean - CL) / (Tolerane / 2)
	cl := (usl + lsl) / 2
	var ca float64
	if tolerance > 0 {
		ca = (mean - cl) / (tolerance / 2)
	}

	// Cpu = (USL - Mean) / (3 * sigma)
	// Cpl = (Mean - LSL) / (3 * sigma)
	cpu := 0.0
	cpl := 0.0
	if stdDev > 0 {
		cpu = (usl - mean) / (3 * stdDev)
		cpl = (mean - lsl) / (3 * stdDev)
	}
	cpk := math.Min(cpu, cpl)

	return SPCStats{
		Mean:   mean,
		StdDev: stdDev,
		Cp:     cp,
		Cpk:    cpk,
		Ca:     ca,
		Count:  n,
	}
}
