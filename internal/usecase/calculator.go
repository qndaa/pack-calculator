package usecase

import (
	"context"

	"github.com/qndaa/pack-calculator/internal/model/dto"
)

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Calculate(ctx context.Context, input *dto.CalculateRequest) (*dto.CalculateResponse, error) {
	return nil, nil
}
