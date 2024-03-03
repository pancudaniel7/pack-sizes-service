package service

import (
	"github.com/pancudaniel7/pack-sizes-service/api/pack"
)

type PackServiceInterface interface {
	SetPackSize(packDTO pack.DTO) error
	calculatePack()
}
