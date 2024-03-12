// Package service provides business logic for handling pack sizes
// and calculating quantities for orders based on given pack sizes.
package service

import (
	"github.com/pancudaniel7/pack-sizes-service/api/dto"
	"github.com/pancudaniel7/pack-sizes-service/internal/dao"
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"math"
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
func (s *DefaultPackService) CalculatePacks(orderQty int) (*[]dto.PackQuantitiesDTO, error) {
	packSizesObj, err := s.dao.GetPackSize()
	if err != nil {
		return nil, err
	}

	var packSizes []int
	for _, size := range packSizesObj.Sizes {
		packSizes = append(packSizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	sqpDTOMap := calculatePacks(&packSizes, orderQty)
	sqpDTOs := convertSQPMap(sqpDTOMap)

	return sqpDTOs, nil
}

func convertSQPMap(sqpDTOMap map[int]int) *[]dto.PackQuantitiesDTO {
	packQuantities := make([]dto.PackQuantitiesDTO, 0, len(sqpDTOMap))
	for quantity, size := range sqpDTOMap {
		packQuantities = append(packQuantities, dto.PackQuantitiesDTO{
			Size:     size,
			Quantity: quantity,
		})
	}
	return &packQuantities
}

func calculatePacks(orderPackSizes *[]int, orderQuantity int) map[int]int {

	seSlots := math.MaxInt32        // selected empty slots
	snPacks := math.MaxInt32        // selected number of packs
	spQuantity := make(map[int]int) // selected pack quantity

	for _, packSize := range *orderPackSizes {

		es, np, pq := recursiveCalculatePacks(orderPackSizes, orderQuantity, packSize, 0)
		if es < seSlots {
			snPacks = np
			seSlots = es
			spQuantity = pq
		} else if es == seSlots && np < snPacks {
			snPacks = np
			seSlots = es
			spQuantity = pq
		}
	}

	return spQuantity
}

func recursiveCalculatePacks(orderPackSizes *[]int, remainingOrderQuantity int, selectedPack int, numberOfPacks int) (int, int, map[int]int) {
	numberOfPacks++

	seSlots := math.MaxInt32        // selected empty slots
	snPacks := numberOfPacks        // selected number of packs
	cPack := selectedPack           // current pack
	spQuantity := make(map[int]int) // selected pack quantity

	if remainingOrderQuantity <= selectedPack {
		seSlots = selectedPack - remainingOrderQuantity
	} else {
		remainingOrderQuantity = remainingOrderQuantity - selectedPack

		for _, packSize := range *orderPackSizes {
			es, np, pq := recursiveCalculatePacks(orderPackSizes, remainingOrderQuantity, packSize, numberOfPacks)
			if es < seSlots {
				seSlots = es
				snPacks = np
				cPack = packSize
				spQuantity = pq
			} else if es == seSlots && np < snPacks {
				snPacks = np
				cPack = packSize
				spQuantity = pq
			}
		}
	}

	if val, exists := spQuantity[cPack]; exists {
		spQuantity[cPack] = val + 1
	} else {
		spQuantity[cPack] = 1
	}

	return seSlots, snPacks, spQuantity
}
