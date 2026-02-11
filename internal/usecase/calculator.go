package usecase

import (
	"context"
	"slices"

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

// TODO: fix the algorithm to minimize the number of packs used, not just a greedy approach
func (c *Calculator) Calculate(ctx context.Context, input *dto.CalculateRequest) (*dto.CalculateResponse, error) {
	packs := c.packRepo.GetPacks()

	// pack size -> count
	// easy to convert to response format later
	// access to count by pack size is O(1)
	result := map[int]int{}

	// Sort packs in asc order to try larger packs first
	slices.SortFunc(packs, func(a, b int) int {
		return b - a
	})

	remaining := input.Items
	//for i := 0; i < len(packs)-1 && remaining > 0; i++ {
	//	if
	//}

	for _, pack := range packs {
		if remaining >= pack {
			count := remaining / pack
			result[pack] = count
			remaining -= count * pack
		}
	}

	// If there are remaining items that cannot be packed, we can add one more smallest pack
	if remaining > 0 {
		smallestPack := packs[len(packs)-1]
		result[smallestPack]++
	}

	// Convert result map to response format
	response := make([]dto.PackResponse, 0, len(result))
	for packSize, count := range result {
		response = append(response, dto.PackResponse{
			Value:    packSize,
			Quantity: count,
		})
	}

	return &dto.CalculateResponse{
		Packs: response,
	}, nil
}
