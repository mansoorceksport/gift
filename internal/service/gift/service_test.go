package gift

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/gift/internal/domain/common/database/mongo"
	"github.com/mansoorceksport/gift/internal/domain/common/entity"
	"github.com/mansoorceksport/gift/internal/domain/product"
	mongo2 "github.com/mansoorceksport/gift/internal/domain/product/mongo"
	"log"
	"testing"
)

var customerId uuid.UUID
var service *Service
var ctx context.Context
var products map[uuid.UUID]product.Product
var giftItems []*entity.GiftItem

func TestMain(m *testing.M) {
	ctx = context.Background()
	customerId = uuid.MustParse("71a193ce-106b-40ba-bb60-c2f2d448cddf")
	mongoDb, err := mongo.NewMongo(ctx, "mongodb://localhost:55575")
	if err != nil {
		log.Fatal(err)
	}
	service = NewGiftService(
		WithMongoGiftRepository(mongoDb),
		WithMongoCustomerRepository(mongoDb),
	)

	products = map[uuid.UUID]product.Product{}
	apple, err := product.NewProduct("apple", "just apple")
	if err != nil {
		log.Fatal(err)
	}
	apple.SetQuantity(100)
	products[apple.GetId()] = apple

	orange, err := product.NewProduct("orange", "just orange")
	if err != nil {
		log.Fatal(err)
	}
	orange.SetQuantity(100)
	products[orange.GetId()] = orange

	productMongoRepository := mongo2.NewProductMongoRepository(mongoDb)
	for _, p := range products {
		err = productMongoRepository.Add(p)
		if err != nil {
			return
		}
	}

	giftItems = append(giftItems, &entity.GiftItem{
		Id:                apple.GetId(),
		Name:              apple.GetItem().Name,
		Description:       apple.GetItem().Description,
		QuantityRequested: apple.GetQuantity(),
		QuantityLocked:    0,
		QuantityReleased:  0,
	}, &entity.GiftItem{
		Id:                orange.GetId(),
		Name:              orange.GetItem().Name,
		Description:       orange.GetItem().Description,
		QuantityRequested: orange.GetQuantity(),
		QuantityLocked:    0,
		QuantityReleased:  0,
	})

	m.Run()
}

func TestService_Create(t *testing.T) {
	err := service.Create(customerId, "sultan birth day", giftItems)
	if err != nil {
		t.Fatal(err)
	}
}
