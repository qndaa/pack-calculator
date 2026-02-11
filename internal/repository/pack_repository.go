package repository

import (
	"github.com/qndaa/pack-calculator/internal/model/domain"
)

type PackRepository struct {
	packs domain.Packs
}

// NewPackRepository loads packs from JSON file and stores them in memory
func NewPackRepository() (*PackRepository, error) {
	// For simplicity, we hardcode the packs here. In a real application, you would load this from a Database or a JSON file.
	return &PackRepository{
		packs: []int{250, 500, 1000, 2000, 5000},
	}, nil
}

func (r *PackRepository) GetPacks() []int {
	return r.packs
}
