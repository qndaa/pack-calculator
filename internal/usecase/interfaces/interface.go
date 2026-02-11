package interfaces

import (
	"context"

	"github.com/qndaa/pack-calculator/internal/model/dto"
)

type Calculator interface {
	Calculate(ctx context.Context, input *dto.CalculateRequest) (*dto.CalculateResponse, error)
}
