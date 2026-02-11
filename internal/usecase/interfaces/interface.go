package interfaces

import (
	"context"

	"github.com/qndaa/pack-calculator/internal/model/dto"
)

type Calculator interface {
	Calculate(ctx context.Context, input *dto.CalculateRequest) (*dto.CalculateResponse, error)
}

type PackRetriever interface {
	GetPacks(ctx context.Context) *dto.GetPacksResponse
}

type PackRemover interface {
	RemovePack(ctx context.Context, size int) error
}

type PackCreator interface {
	CreatePack(ctx context.Context, size int) error
}
