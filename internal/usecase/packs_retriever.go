package usecase

import (
	"context"

	"github.com/qndaa/pack-calculator/internal/model/dto"
	"github.com/qndaa/pack-calculator/internal/repository/interfaces"
)

type PacksRetriever struct {
	packRepo interfaces.PackRepository
}

func NewPacksRetriever(packRepo interfaces.PackRepository) *PacksRetriever {
	return &PacksRetriever{
		packRepo: packRepo,
	}
}

func (p *PacksRetriever) GetPacks(ctx context.Context) *dto.GetPacksResponse {
	packs := p.packRepo.GetPacks()
	return &dto.GetPacksResponse{
		Packs: packs,
	}
}
