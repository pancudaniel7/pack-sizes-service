package dao

import "github.com/pancudaniel7/pack-sizes-service/internal/model"

type PackDao interface {
	GetPackSize() (model.Pack, error)
	AddPackSize(pack model.Pack) error
}
