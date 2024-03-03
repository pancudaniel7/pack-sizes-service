package service

import (
	"errors"
	"github.com/pancudaniel7/pack-sizes-service/api/pack"
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockPackDao is a mock type for the PackDao interface
type MockPackDao struct {
	mock.Mock
}

// AddPackSize mocks the AddPackSize method
func (m *MockPackDao) AddPackSize(pack model.Pack) error {
	args := m.Called(pack)
	return args.Error(0)
}

// GetPackSize mocks the GetPackSize method
func (m *MockPackDao) GetPackSize() (model.Pack, error) {
	args := m.Called()
	return args.Get(0).(model.Pack), args.Error(1)
}

func TestSetPackSize(t *testing.T) {
	mockDao := new(MockPackDao)
	service := NewPackService(mockDao)
	packDTO := pack.DTO{Sizes: []int{5, 10, 20}}

	mockDao.On("AddPackSize", mock.Anything).Return(nil)

	err := service.SetPackSize(packDTO)

	assert.NoError(t, err)
	mockDao.AssertExpectations(t)
}

func TestCalculatePacks_Success(t *testing.T) {
	mockDao := new(MockPackDao)
	service := NewPackService(mockDao)
	orderQty := 50
	expectedPacks := []pack.SizeQuantityDTO{
		{Size: 20, Quantity: 2},
		{Size: 10, Quantity: 1},
	}

	mockDao.On("GetPackSize").Return(model.Pack{Sizes: []int{5, 10, 20}}, nil)

	packs, err := service.CalculatePacks(orderQty)

	assert.NoError(t, err)
	assert.Equal(t, expectedPacks, packs)
	mockDao.AssertExpectations(t)
}

func TestCalculatePacks_DaoError(t *testing.T) {
	mockDao := new(MockPackDao)
	service := NewPackService(mockDao)

	mockDao.On("GetPackSize").Return(model.Pack{}, errors.New("fail to access pack sizes"))

	_, err := service.CalculatePacks(50)

	assert.Error(t, err)
	mockDao.AssertExpectations(t)
}
