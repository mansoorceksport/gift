package product

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/gift/internal/domain/common/entity"
)

type Product struct {
	item *entity.Item
}

func NewProduct(name, description string) (Product, error) {
	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
			Quantity:    0,
		},
	}, nil
}

func (p Product) GetId() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) SetItem(item *entity.Item) {
	p.item = item
}

func (p Product) GetQuantity() int {
	return p.item.Quantity
}

func (p Product) SetQuantity(q int) {
	p.item.Quantity = q
}
