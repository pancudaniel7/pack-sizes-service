// Package service provides business logic for handling pack sizes
// and calculating quantities for orders based on given pack sizes.
package service

import (
	"github.com/pancudaniel7/pack-sizes-service/api/dto"
	"github.com/pancudaniel7/pack-sizes-service/internal/dao"
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"sort"
)

// DefaultPackService provides methods for managing pack sizes
// and calculating pack quantities for orders.
type DefaultPackService struct {
	dao dao.PackDao
}

// NewDefaultPackService creates a new instance of DefaultPackService
// with the given dao.PackDao.
func NewDefaultPackService(dao dao.PackDao) *DefaultPackService {
	return &DefaultPackService{
		dao: dao,
	}
}

// SetPackSize persists a new set of pack sizes.
func (s *DefaultPackService) SetPackSize(packDTO dto.PackSizesDTO) error {
	packObj := model.Pack{Sizes: packDTO.Sizes}
	return s.dao.AddPackSize(packObj)
}

// CalculatePacks calculates the optimal quantities of packs needed
// for the given order quantity using available pack sizes.
func (s *DefaultPackService) CalculatePacks(orderQty int) ([]dto.PackQuantitiesDTO, error) {
	packSizesObj, err := s.dao.GetPackSize()
	if err != nil {
		return nil, err
	}

	var packSizes []int
	for _, size := range packSizesObj.Sizes {
		packSizes = append(packSizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	var sqpDTO []dto.PackQuantitiesDTO
	calculatePacks(&sqpDTO, packSizes, orderQty, 0)

	return sqpDTO, nil
}

// calculatePacks is a recursive function that determines the combination of pack sizes
// that can be used to fulfill an order quantity. It takes into account the sizes of packs
// available in descending order and calculates the number of packs required for each size.
func calculatePacks(sizeQuantityPacks *[]dto.PackQuantitiesDTO, descOrderedPackSizes []int, orderQuantity int, cursor int) {

	for _, packSize := range descOrderedPackSizes[cursor:] {
		lastItem := descOrderedPackSizes[len(descOrderedPackSizes)-1]
		if packSize == lastItem || orderQuantity > packSize {
			div := orderQuantity / packSize
			rem := orderQuantity % packSize

			sameQuantityIndex := findSameQuantityIndex(sizeQuantityPacks, packSize)
			if sameQuantityIndex != -1 {
				changeSameQuantity(sizeQuantityPacks, &descOrderedPackSizes, sameQuantityIndex)
				if div == 0 {
					break
				}
			} else {
				if div == 0 {
					*sizeQuantityPacks = append(*sizeQuantityPacks, dto.PackQuantitiesDTO{Quantity: 1, Size: packSize})
					break
				}
				*sizeQuantityPacks = append(*sizeQuantityPacks, dto.PackQuantitiesDTO{Quantity: div, Size: packSize})
			}

			if rem != 0 {
				calculatePacks(sizeQuantityPacks, descOrderedPackSizes, rem, cursor+1)
				break
			}
		}
	}
}

func changeSameQuantity(sizeQuantityPacks *[]dto.PackQuantitiesDTO, descOrderedPackSizes *[]int, index int) {
	if index != -1 {
		samePackSize := (*sizeQuantityPacks)[index].Size
		packSizeIndex := findPackSizeIndex(descOrderedPackSizes, samePackSize)

		(*sizeQuantityPacks)[index].Size = (*descOrderedPackSizes)[packSizeIndex-1]
		(*sizeQuantityPacks)[index].Quantity = (*sizeQuantityPacks)[index].Quantity + 1
	}
}

// findPackSizeIndex returns the index of a specific pack size within the slice
// of pack sizes ordered in descending order.
func findPackSizeIndex(descOrderedPackSizes *[]int, packSize int) int {
	for index, size := range *descOrderedPackSizes {
		if size == packSize {
			return index
		}
	}
	return -1
}

// findSameQuantityIndex searches for a pack size in the slice of pack sizes and quantities
// and returns the index of the matching pack size if found.
func findSameQuantityIndex(sizeQuantityPacks *[]dto.PackQuantitiesDTO, packSize int) int {
	for index, sizeQuantity := range *sizeQuantityPacks {
		if sizeQuantity.Size == packSize {
			return index
		}
	}
	return -1
}
