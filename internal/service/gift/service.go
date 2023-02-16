package gift

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/gift/internal/domain/common/entity"
	"github.com/mansoorceksport/gift/internal/domain/customer"
	customerRepository "github.com/mansoorceksport/gift/internal/domain/customer/mongo"
	"github.com/mansoorceksport/gift/internal/domain/gift"
	giftRepository "github.com/mansoorceksport/gift/internal/domain/gift/mongo"
	"github.com/mansoorceksport/gift/internal/domain/product"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceGiftConfiguration func(service *Service)

// Service Gift handles customer requests for gift
type Service struct {
	gift     gift.Repository
	customer customer.Repository
	products product.Repository
}

func NewGiftService(config ...ServiceGiftConfiguration) *Service {
	service := &Service{}
	for _, configuration := range config {
		configuration(service)
	}
	return service
}

func WithMongoGiftRepository(db *mongo.Database) ServiceGiftConfiguration {
	return func(service *Service) {
		service.gift = giftRepository.NewGiftMongo(db)
	}
}

func WithMongoCustomerRepository(db *mongo.Database) ServiceGiftConfiguration {
	return func(service *Service) {
		service.customer = customerRepository.NewCustomerRepository(db)
	}
}

func (s Service) Create(customerId uuid.UUID, name string, giftItems []*entity.GiftItem) error {
	c, err := s.customer.GetById(customerId)
	if err != nil {
		return err
	}
	newGift, err := gift.NewGift(name, c.GetEntity())
	if err != nil {
		return err
	}
	for _, item := range giftItems {
		newGift.AddGiftItem(item)
	}

	err = s.gift.Create(*newGift)
	if err != nil {
		return err
	}

	return nil
}
