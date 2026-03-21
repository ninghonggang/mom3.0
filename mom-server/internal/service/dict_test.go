package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"mom-server/internal/dto"
	"gorm.io/gorm"
)

// MockDictRepository is a mock implementation for testing
type MockDictRepository struct {
	dicts     map[int64]*DictType
	nextID    int64
	dataDict  map[string][]DictData
}

func NewMockDictRepository() *MockDictRepository {
	return &MockDictRepository{
		dicts:    make(map[int64]*DictType),
		nextID:   1,
		dataDict: make(map[string][]DictData),
	}
}

type DictType struct {
	ID        int64
	DictName  string
	DictType  string
	Status    int
}

type DictData struct {
	ID        int64
	DictLabel string
	DictValue string
	DictType  string
}

func (m *MockDictRepository) Create(ctx context.Context, dict *DictType) error {
	dict.ID = m.nextID
	m.nextID++
	m.dicts[dict.ID] = dict
	return nil
}

func (m *MockDictRepository) Update(ctx context.Context, id int64, updates map[string]interface{}) error {
	if dict, ok := m.dicts[id]; ok {
		if name, ok := updates["dict_name"].(string); ok {
			dict.DictName = name
		}
		if dtype, ok := updates["dict_type"].(string); ok {
			dict.DictType = dtype
		}
		if status, ok := updates["status"].(int); ok {
			dict.Status = status
		}
	}
	return nil
}

func (m *MockDictRepository) Delete(ctx context.Context, id int64) error {
	delete(m.dicts, id)
	return nil
}

func (m *MockDictRepository) FindByID(ctx context.Context, id int64) (*DictType, error) {
	if dict, ok := m.dicts[id]; ok {
		return dict, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockDictRepository) FindByType(ctx context.Context, dictType string) (*DictType, error) {
	for _, dict := range m.dicts {
		if dict.DictType == dictType {
			return dict, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockDictRepository) FindByPage(ctx context.Context, req dto.PageRequest) ([]DictType, int64, error) {
	var result []DictType
	for _, dict := range m.dicts {
		result = append(result, *dict)
	}
	return result, int64(len(result)), nil
}

func (m *MockDictRepository) CreateData(ctx context.Context, data *DictData) error {
	data.ID = int64(len(m.dataDict) + 1)
	m.dataDict[data.DictType] = append(m.dataDict[data.DictType], *data)
	return nil
}

func (m *MockDictRepository) FindDataByType(ctx context.Context, dictType string) ([]DictData, error) {
	if data, ok := m.dataDict[dictType]; ok {
		return data, nil
	}
	return []DictData{}, nil
}

// Tests
func TestMockDictRepository_Create(t *testing.T) {
	repo := NewMockDictRepository()

	dict := &DictType{
		DictName: "Status",
		DictType: "sys_status",
		Status:   1,
	}

	err := repo.Create(context.Background(), dict)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), dict.ID)
}

func TestMockDictRepository_FindByID(t *testing.T) {
	repo := NewMockDictRepository()

	dict := &DictType{
		DictName: "Status",
		DictType: "sys_status",
		Status:   1,
	}
	repo.Create(context.Background(), dict)

	found, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "Status", found.DictName)
	assert.Equal(t, "sys_status", found.DictType)
}

func TestMockDictRepository_FindByID_NotFound(t *testing.T) {
	repo := NewMockDictRepository()

	_, err := repo.FindByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestMockDictRepository_Update(t *testing.T) {
	repo := NewMockDictRepository()

	dict := &DictType{
		DictName: "Status",
		DictType: "sys_status",
		Status:   1,
	}
	repo.Create(context.Background(), dict)

	updates := map[string]interface{}{
		"dict_name": "Updated Status",
		"status":    0,
	}
	err := repo.Update(context.Background(), 1, updates)
	assert.NoError(t, err)

	updated, _ := repo.FindByID(context.Background(), 1)
	assert.Equal(t, "Updated Status", updated.DictName)
	assert.Equal(t, 0, updated.Status)
}

func TestMockDictRepository_Delete(t *testing.T) {
	repo := NewMockDictRepository()

	dict := &DictType{
		DictName: "Status",
		DictType: "sys_status",
		Status:   1,
	}
	repo.Create(context.Background(), dict)

	err := repo.Delete(context.Background(), 1)
	assert.NoError(t, err)

	_, err = repo.FindByID(context.Background(), 1)
	assert.Error(t, err)
}

func TestMockDictRepository_FindByType(t *testing.T) {
	repo := NewMockDictRepository()

	dict := &DictType{
		DictName: "Status",
		DictType: "sys_status",
		Status:   1,
	}
	repo.Create(context.Background(), dict)

	found, err := repo.FindByType(context.Background(), "sys_status")
	assert.NoError(t, err)
	assert.Equal(t, "Status", found.DictName)
}

func TestMockDictRepository_FindByType_NotFound(t *testing.T) {
	repo := NewMockDictRepository()

	_, err := repo.FindByType(context.Background(), "non_existent")
	assert.Error(t, err)
}

func TestMockDictRepository_FindByPage(t *testing.T) {
	repo := NewMockDictRepository()

	repo.Create(context.Background(), &DictType{DictName: "Dict1", DictType: "type1", Status: 1})
	repo.Create(context.Background(), &DictType{DictName: "Dict2", DictType: "type2", Status: 1})

	dicts, total, err := repo.FindByPage(context.Background(), dto.PageRequest{Page: 1, PageSize: 10})
	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, dicts, 2)
}

func TestMockDictRepository_DictData(t *testing.T) {
	repo := NewMockDictRepository()

	// Create dict type first
	dict := &DictType{
		DictName: "Status",
		DictType: "sys_status",
		Status:   1,
	}
	repo.Create(context.Background(), dict)

	// Create dict data
	data := &DictData{
		DictLabel: "Enabled",
		DictValue: "1",
		DictType:  "sys_status",
	}
	err := repo.CreateData(context.Background(), data)
	assert.NoError(t, err)

	// Find data by type
	items, err := repo.FindDataByType(context.Background(), "sys_status")
	assert.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, "Enabled", items[0].DictLabel)
	assert.Equal(t, "1", items[0].DictValue)
}
