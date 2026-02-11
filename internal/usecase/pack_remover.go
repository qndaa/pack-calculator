package usecase

import (
	"context"

	"github.com/qndaa/pack-calculator/internal/repository/interfaces"
)

type PackRemover struct {
	packRepository interfaces.PackRepository
}

func NewPackRemover(packRepository interfaces.PackRepository) *PackRemover {
	return &PackRemover{
		packRepository: packRepository,
	}
}

func (p *PackRemover) RemovePack(ctx context.Context, size int) error {
	return p.packRepository.Delete(size)
}
