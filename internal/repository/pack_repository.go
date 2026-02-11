package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/qndaa/pack-calculator/internal/model/domain"
)

type PackRepository struct {
	packs domain.Packs
}

// NewPackRepository loads packs from JSON file and stores them in memory
func NewPackRepository(config *Config) (*PackRepository, error) {
	// Read the packs from the JSON file, because we want to load them once and keep in memory for fast access
	data, err := os.ReadFile(config.PacksFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read packs file: %w", err)
	}

	var packs domain.Packs
	if err := json.Unmarshal(data, &packs); err != nil {
		return nil, fmt.Errorf("failed to parse packs file: %w", err)
	}

	if len(packs) == 0 {
		return nil, fmt.Errorf("no packs defined in config")
	}

	return &PackRepository{
		packs: packs,
	}, nil
}

func (r *PackRepository) GetPacks() []int {
	return r.packs
}
