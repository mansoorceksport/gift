package order

import (
	"github.com/mansoorceksport/gift/internal/domain/common/idempotent"
	"github.com/mansoorceksport/gift/internal/domain/customer"
	"github.com/mansoorceksport/gift/internal/domain/product"
)

// Service Order handles customer order
type Service struct {
	idempotent idempotent.Idempotent
	customer   customer.Repository
	products   product.Repository
}
