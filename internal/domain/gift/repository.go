package gift

import "github.com/google/uuid"

type Repository interface {
	Create(gift Gift) error
	GetById(id uuid.UUID) (Gift, error)
}
