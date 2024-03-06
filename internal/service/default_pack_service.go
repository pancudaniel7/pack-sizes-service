package service

import (
	"github.com/pancudaniel7/pack-sizes-service/api/dto"
	"github.com/pancudaniel7/pack-sizes-service/internal/dao"
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"sort"
)

type DefaultPackService struct {
	dao dao.PackDao
}

func NewDefaultPackService(dao dao.PackDao) *DefaultPackService {
	return &DefaultPackService{
		dao: dao,
	}
}

func (s *DefaultPackService) SetPackSize(packDTO dto.PackSizesDTO) error {
	packObj := model.Pack{Sizes: packDTO.Sizes}
	return s.dao.AddPackSize(packObj)
}

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

func findPackSizeIndex(descOrderedPackSizes *[]int, packSize int) int {
	for index, size := range *descOrderedPackSizes {
		if size == packSize {
			return index
		}
	}
	return -1
}

func findSameQuantityIndex(sizeQuantityPacks *[]dto.PackQuantitiesDTO, packSize int) int {
	for index, sizeQuantity := range *sizeQuantityPacks {
		if sizeQuantity.Size == packSize {
			return index
		}
	}
	return -1
}
