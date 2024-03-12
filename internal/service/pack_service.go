package service

import (
	"github.com/pancudaniel7/pack-sizes-service/api/dto"
)

type PackService interface {
	SetPackSize(packDTO dto.PackSizesDTO) error
	CalculatePacks(orderQty int) (*[]dto.PackQuantitiesDTO, error)
}
