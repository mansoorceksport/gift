package inventory

import "github.com/mansoorceksport/gift/internal/domain/product"

// Service Inventory handles internal products
type Service struct {
	products product.Repository
}
