package customer

import "github.com/google/uuid"

type Repository interface {
	GetById(id uuid.UUID) (Customer, error)
	Add(customer Customer) error
	Remove(id uuid.UUID) error
	Update(customer Customer) error
}
