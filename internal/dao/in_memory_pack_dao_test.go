package dao

import (
	"github.com/pancudaniel7/pack-sizes-service/internal/model"
	"testing"
)

func TestInMemoryPackDao_GetPackSize(t *testing.T) {
	initialPack := model.Pack{Sizes: []int{250, 500, 1000, 2000, 5000}}
	dao := NewInMemoryPackDao(initialPack)

	pack, err := dao.GetPackSize()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(pack.Sizes) != len(initialPack.Sizes) {
		t.Fatalf("Expected %v, got %v", initialPack.Sizes, pack.Sizes)
	}
}

func TestInMemoryPackDao_AddPackSize(t *testing.T) {
	initialPack := model.Pack{Sizes: []int{250, 500, 1000, 2000, 5000}}
	dao := NewInMemoryPackDao(initialPack)

	newPack := model.Pack{Sizes: []int{100, 200, 300, 400, 500}}
	err := dao.AddPackSize(newPack)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	pack, err := dao.GetPackSize()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(pack.Sizes) != len(newPack.Sizes) {
		t.Fatalf("Expected %v, got %v", newPack.Sizes, pack.Sizes)
	}
}
