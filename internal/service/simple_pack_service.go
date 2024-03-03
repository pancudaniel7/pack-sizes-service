package service

import (
	"github.com/pancudaniel7/pack-sizes-service/api/pack"
	"github.com/pancudaniel7/pack-sizes-service/internal/dao"
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"sort"
)

type PackService struct {
	dao dao.PackDao
}

func NewPackService(dao dao.PackDao) *PackService {
	return &PackService{
		dao: dao,
	}
}

func (s *PackService) SetPackSize(packDTO pack.DTO) error {
	packObj := model.Pack{Sizes: packDTO.Sizes}
	return s.dao.AddPackSize(packObj)
}

func (s *PackService) CalculatePacks(orderQty int) ([]pack.SizeQuantityDTO, error) {
	packObj, err := s.dao.GetPackSize()
	if err != nil {
		return nil, err
	}

	packs := calculatePacks(orderQty, packObj.Sizes)
	return packs, nil
}

func calculatePacks(orderQty int, packSizes []int) []pack.SizeQuantityDTO {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	var packs []pack.SizeQuantityDTO
	remaining := orderQty

	for _, size := range packSizes {
		if remaining >= size {
			quantity := remaining / size
			packs = append(packs, pack.SizeQuantityDTO{Size: size, Quantity: quantity})
			remaining %= size
		}
	}

	if remaining > 0 {
		for _, size := range packSizes {
			if size >= remaining {
				packs = append(packs, pack.SizeQuantityDTO{Size: size, Quantity: 1})
				break
			}
		}
	}

	return packs
}
