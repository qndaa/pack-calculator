package repository

import (
	"fmt"

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

func (r *PackRepository) FindAll() domain.Packs {
	return r.packs
}

func (r *PackRepository) Delete(size int) error {
	for i, pack := range r.packs {
		if pack == size {
			r.packs = append(r.packs[:i], r.packs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("pack size %d not found", size)
}

func (r *PackRepository) Create(size int) error {
	// Check if the pack already exists
	for _, pack := range r.packs {
		if pack == size {
			return fmt.Errorf("pack size %d already exists", size)
		}
	}

	r.packs = append(r.packs, size)
	return nil
}
