package dto

import "fmt"

// CalculateRequest represents the request payload for calculating the optimal pack combination.
type CalculateRequest struct {
	Items int `json:"items"`
}

// Validate checks if the CalculateRequest is valid.
// It returns an error if the number of items is less than or equal to zero.
// In a real application, body request should be validated with go-playground/validator or similar library, but for simplicity, we implement a basic validation method here.
func (r *CalculateRequest) Validate() error {
	if r.Items <= 0 {
		return fmt.Errorf("items must be greater than 0")
	}
	return nil
}

type Pack struct {
	Value    int `json:"value"`
	Quantity int `json:"quantity"`
}

type CalculateResponse struct {
	Packs []Pack `json:"packs"`
}

type GetPacksResponse struct {
	Packs []int `json:"packs"`
}
