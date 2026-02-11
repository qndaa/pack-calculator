package usecase

import (
	"context"
	"slices"

	"github.com/qndaa/pack-calculator/internal/model/domain"
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
	packs := c.packRepo.FindAll()

	itemsToSend, numberOfItemsWithSteps := c.findMinimumItemsToSend(input, packs)
	minimumPacks := c.findMinimumPacks(itemsToSend, numberOfItemsWithSteps)

	return &dto.CalculateResponse{
		Packs: minimumPacks,
	}, nil
}

// findMinimumItemsToSend calculates the minimum number of items to send that is greater than or equal to the requested items, 
// and also returns a map of how we can reach that number using the available packs.
func (c *Calculator) findMinimumItemsToSend(input *dto.CalculateRequest, packs domain.Packs) (int, map[int]int) {
	output := input.Items

	slices.SortFunc(packs, func(a, b int) int { return b - a })
	minPackSize := packs[len(packs)-1]
	maxPossibleItems := output + minPackSize

	// dp[i] = minimum number of packs to reach exactly i items, -1 means unreachable
	dp := make([]int, maxPossibleItems+1)
	parent := make([]int, maxPossibleItems+1) // stores which pack size was used
	for i := range dp {
		dp[i] = -1
		parent[i] = -1
	}
	dp[0] = 0

	for i := 1; i <= maxPossibleItems; i++ {
		for _, size := range packs {
			if i >= size && dp[i-size] != -1 {
				newCount := dp[i-size] + 1
				if dp[i] == -1 || newCount < dp[i] {
					dp[i] = newCount
					parent[i] = size
				}
			}
		}
	}

	// Find smallest reachable value >= input.Items
	for i := output; i <= maxPossibleItems; i++ {
		if dp[i] != -1 {
			output = i
			break
		}
	}

	// Convert parent array to map for backward compatibility
	numberOfItemsWithSteps := make(map[int]int)
	for i, size := range parent {
		if size != -1 {
			numberOfItemsWithSteps[i] = size
		}
	}

	return output, numberOfItemsWithSteps
}


// findMinimumPacks takes the minimum number of items to send and the map of how we can reach that number,
//  and converts it into a list of packs with their quantities.
func (c *Calculator) findMinimumPacks(itemsToSend int, numberOfItemsWithSteps map[int]int) []dto.PackResponse {
	counts := make(map[int]int)
	curr := itemsToSend
	for curr > 0 {
		size := numberOfItemsWithSteps[curr]
		counts[size]++
		curr -= size
	}

	// Convert map to sorted DTO (descending by pack size)
	output := make([]dto.PackResponse, 0, len(counts))
	for size, quantity := range counts {
		output = append(output, dto.PackResponse{
			Value:    size,
			Quantity: quantity,
		})
	}

	// Sort by Value descending to ensure consistent ordering
	slices.SortFunc(output, func(a, b dto.PackResponse) int {
		return b.Value - a.Value
	})

	return output
}
