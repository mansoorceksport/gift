package product

import "github.com/google/uuid"

type Repository interface {
	GetById(id uuid.UUID) (Product, error)
	Add(product Product) error
	Update(product Product) error
}
