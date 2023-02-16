package gift

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/gift/internal/domain/common/entity"
)

var (
	ErrGiftCustomerRequired = errors.New("customer required")
)

// Gift Aggregate
type Gift struct {
	gift      *entity.Gift
	customer  *entity.Person
	giftItems []*entity.GiftItem
}

func NewGift(name string, customer *entity.Person) (*Gift, error) {
	if customer == nil {
		return nil, ErrGiftCustomerRequired
	}
	return &Gift{
		gift: &entity.Gift{
			ID:   uuid.New(),
			Name: name,
		},
		customer:  customer,
		giftItems: make([]*entity.GiftItem, 0),
	}, nil
}

func (g *Gift) ConvertItemToGiftItem(item *entity.Item) error {
	giftItem := &entity.GiftItem{
		Id:                item.ID,
		Name:              item.Name,
		Description:       item.Description,
		QuantityRequested: item.Quantity,
		QuantityLocked:    0,
		QuantityReleased:  0,
	}
	g.giftItems = append(g.giftItems, giftItem)
	return nil
}

func (g *Gift) AddGiftItem(item *entity.GiftItem) {
	g.giftItems = append(g.giftItems, item)
}

func (g *Gift) GetName() string {
	return g.gift.Name
}

func (g *Gift) SetName(name string) {
	g.gift.Name = name
}

func (g *Gift) GetID() uuid.UUID {
	return g.gift.ID
}

func (g *Gift) SetID(id uuid.UUID) {
	g.gift.ID = id
}

func (g *Gift) SetCustomer(person *entity.Person) {
	g.customer = person
}

func (g *Gift) GetCustomer() *entity.Person {
	return g.customer
}

func (g *Gift) GetProducts() []*entity.GiftItem {
	return g.giftItems
}

func (g *Gift) GetMapProducts() map[uuid.UUID]*entity.Item {
	var items map[uuid.UUID]*entity.Item
	for _, gi := range g.giftItems {
		items[gi.Id] = &entity.Item{
			ID:          gi.Id,
			Name:        gi.Name,
			Description: gi.Name,
			Quantity:    gi.QuantityRequested,
		}
	}
	return items
}
