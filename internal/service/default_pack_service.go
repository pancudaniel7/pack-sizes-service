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

func (s *DefaultPackService) SetPackSize(packDTO dto.PackDTO) error {
	packObj := model.Pack{Sizes: packDTO.Sizes}
	return s.dao.AddPackSize(packObj)
}

func (s *DefaultPackService) CalculatePacks(orderQty int) ([]dto.SizeQuantityPackDTO, error) {
	packObj, err := s.dao.GetPackSize()
	if err != nil {
		return nil, err
	}

	packs := calculatePacks(orderQty, packObj.Sizes)
	return packs, nil
}

func calculatePacks(orderQty int, packSizes []int) []dto.SizeQuantityPackDTO {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	var packs []dto.SizeQuantityPackDTO
	remaining := orderQty

	for _, size := range packSizes {
		if remaining >= size {
			quantity := remaining / size
			packs = append(packs, dto.SizeQuantityPackDTO{Size: size, Quantity: quantity})
			remaining %= size
		}
	}

	if remaining > 0 {
		for _, size := range packSizes {
			if size >= remaining {
				packs = append(packs, dto.SizeQuantityPackDTO{Size: size, Quantity: 1})
				break
			}
		}
	}

	return packs
}
