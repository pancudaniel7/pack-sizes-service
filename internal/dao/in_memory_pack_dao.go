package dao

import (
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"sync"
)

type InMemoryPackDao struct {
	pack model.Pack
	mu   sync.Mutex
}

func NewInMemoryPackDao(initialPack model.Pack) *InMemoryPackDao {
	return &InMemoryPackDao{
		pack: initialPack,
	}
}

func (dao *InMemoryPackDao) GetPackSize() (model.Pack, error) {
	dao.mu.Lock()
	defer dao.mu.Unlock()

	return dao.pack, nil
}

func (dao *InMemoryPackDao) AddPackSize(pack model.Pack) error {
	dao.mu.Lock()
	defer dao.mu.Unlock()

	dao.pack = pack

	return nil
}
