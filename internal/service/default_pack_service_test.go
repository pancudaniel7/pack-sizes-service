package service

import (
	"testing"

	"github.com/pancudaniel7/pack-sizes-service/api/dto"
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Assuming the dao.PackDao interface includes methods AddPackSize and GetPackSize
type MockPackDao struct {
	mock.Mock
}

func (m *MockPackDao) AddPackSize(pack model.Pack) error {
	args := m.Called(pack)
	return args.Error(0)
}

func (m *MockPackDao) GetPackSize() (model.Pack, error) {
	args := m.Called()
	return args.Get(0).(model.Pack), args.Error(1)
}

func TestSetPackSize(t *testing.T) {
	mockDao := new(MockPackDao)
	service := NewDefaultPackService(mockDao)

	packSizes := []int{250, 500, 1000, 2000, 5000}
	packDTO := dto.PackSizesDTO{Sizes: packSizes}
	packObj := model.Pack{Sizes: packSizes}

	// Setup expectation
	mockDao.On("AddPackSize", packObj).Return(nil)

	err := service.SetPackSize(packDTO)

	assert.NoError(t, err)
	mockDao.AssertExpectations(t)
}

func TestCalculatePacks(t *testing.T) {
	mockDao := new(MockPackDao)
	service := NewDefaultPackService(mockDao)

	// Mock response from GetPackSize for available pack sizes
	packSizesObj := model.Pack{Sizes: []int{250, 500, 1000, 2000, 5000}}
	mockDao.On("GetPackSize").Return(packSizesObj, nil)

	orderQty := 12001
	expectedResp := []dto.PackQuantitiesDTO{
		{Size: 5000, Quantity: 2},
		{Size: 2000, Quantity: 1},
		{Size: 250, Quantity: 1},
	}

	result, err := service.CalculatePacks(orderQty)

	assert.NoError(t, err)
	assert.Equal(t, expectedResp, result)
	mockDao.AssertExpectations(t)
}
