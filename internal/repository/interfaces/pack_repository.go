package interfaces

import "github.com/qndaa/pack-calculator/internal/model/domain"

type PackRepository interface {
	GetPacks() domain.Packs
}
