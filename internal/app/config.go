package app

import "github.com/qndaa/pack-calculator/internal/repository"

type Config struct {
	repositoryConfig *repository.Config
}

func NewConfig() *Config {
	// In a real application, you might want to load this from environment variables or a config file
	return &Config{
		repositoryConfig: &repository.Config{
			PacksFile: "packs.json",
		},
	}
}
