package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/qndaa/pack-calculator/internal/model/dto"
	"github.com/qndaa/pack-calculator/internal/repository"
	"github.com/qndaa/pack-calculator/internal/usecase"
	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	// Initialize the PackRepository and Calculator
	// In a real test, you might want to use a mock repository instead of the actual implementation
	// But we are using hardcoded packs in the repository, so it's fine for this test
	packRepository, err := repository.NewPackRepository()
	require.NoError(t, err, "Failed to create PackRepository")

	calculator := usecase.NewCalculator(packRepository)

	testCases := []struct {
		name     string
		input    *dto.CalculateRequest
		expected *dto.CalculateResponse
	}{
		{
			name: "1 item should return 1x250",
			input: &dto.CalculateRequest{
				Items: 1,
			},
			expected: &dto.CalculateResponse{
				Packs: []dto.PackResponse{
					{Value: 250, Quantity: 1},
				},
			},
		},
		{
			name: "250 items should return 1x250",
			input: &dto.CalculateRequest{
				Items: 250,
			},
			expected: &dto.CalculateResponse{
				Packs: []dto.PackResponse{
					{Value: 250, Quantity: 1},
				},
			},
		},
		{
			name: "251 items should return 1x500",
			input: &dto.CalculateRequest{
				Items: 251,
			},
			expected: &dto.CalculateResponse{
				Packs: []dto.PackResponse{
					{Value: 500, Quantity: 1},
				},
			},
		},
		{
			name: "501 items should return 1x500 + 1x250",
			input: &dto.CalculateRequest{
				Items: 501,
			},
			expected: &dto.CalculateResponse{
				Packs: []dto.PackResponse{
					{Value: 500, Quantity: 1},
					{Value: 250, Quantity: 1},
				},
			},
		},
		{
			name: "12001 items should return 2x5000 + 1x2000 + 1x250",
			input: &dto.CalculateRequest{
				Items: 12001,
			},
			expected: &dto.CalculateResponse{
				Packs: []dto.PackResponse{
					{Value: 5000, Quantity: 2},
					{Value: 2000, Quantity: 1},
					{Value: 250, Quantity: 1},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := calculator.Calculate(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("Calculate() error = %v", err)
			}
			if !reflect.DeepEqual(response, tc.expected) {
				t.Errorf("Calculate() = %v, want %v", response, tc.expected)
			}
		})
	}
}
