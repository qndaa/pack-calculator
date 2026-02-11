package usecase

import (
	"context"

	"github.com/qndaa/pack-calculator/internal/model/dto"
	"github.com/qndaa/pack-calculator/internal/repository/interfaces"
)

type Calculator struct {
	packRepo interfaces.PackRepository
}

func NewCalculator(packRepo interfaces.PackRepository) *Calculator {
	return &Calculator{
		packRepo: packRepo,
	}
}

func (c *Calculator) Calculate(ctx context.Context, input *dto.CalculateRequest) (*dto.CalculateResponse, error) {
	return nil, nil
}

func (c *Calculator) GetPacks(ctx context.Context) (*dto.GetPacksResponse, error) {
	packs := c.packRepo.GetPacks()
	return &dto.GetPacksResponse{
		Packs: packs,
	}, nil
}
