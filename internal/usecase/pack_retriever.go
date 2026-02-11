package usecase

import (
	"context"

	"github.com/qndaa/pack-calculator/internal/model/dto"
	"github.com/qndaa/pack-calculator/internal/repository/interfaces"
)

type PackRetriever struct {
	packRepo interfaces.PackRepository
}

func NewPackRetriever(packRepo interfaces.PackRepository) *PackRetriever {
	return &PackRetriever{
		packRepo: packRepo,
	}
}

func (p *PackRetriever) GetPacks(ctx context.Context) *dto.GetPacksResponse {
	packs := p.packRepo.FindAll()
	return &dto.GetPacksResponse{
		Packs: packs,
	}
}
