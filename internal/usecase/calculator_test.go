package usecase_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/qndaa/pack-calculator/internal/model/dto"
	"github.com/qndaa/pack-calculator/internal/usecase"
)


func TestCalculate(t *testing.T) {

	testCases := []struct {
		name     string
		input    *dto.CalculateRequest
		expected *dto.CalculateResponse
	}{
		{
			name: "Calculate 750 items",
			input: &dto.CalculateRequest{
				Items: 750,
			},
			expected: &dto.CalculateResponse{
				Packs: map[int]int{
					500: 1,
					250: 1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calculator := usecase.NewCalculator()
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
