package dto

import "fmt"

type CalculateRequest struct {
	Items int `json:"items"`
}

func (r *CalculateRequest) Validate() error {
	if r.Items <= 0 {
		return fmt.Errorf("items must be greater than 0")
	}
	return nil
}

type CalculateResponse struct {
	Packs map[int]int `json:"packs"`
}
