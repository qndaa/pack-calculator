package usecase

import (
	"context"

	"github.com/qndaa/pack-calculator/internal/repository/interfaces"
)

type PackCreator struct {
	packRepository interfaces.PackRepository
}

func NewPackCreator(packRepository interfaces.PackRepository) *PackCreator {
	return &PackCreator{
		packRepository: packRepository,
	}
}

func (p *PackCreator) CreatePack(ctx context.Context, size int) error {
	return p.packRepository.Create(size)
}
