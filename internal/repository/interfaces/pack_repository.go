package interfaces

import "github.com/qndaa/pack-calculator/internal/model/domain"

type PackRepository interface {
	FindAll() domain.Packs
	Delete(size int) error
	Create(size int) error
}
