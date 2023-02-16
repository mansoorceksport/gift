package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/gift/internal/domain/common/entity"
	mongoCustomer "github.com/mansoorceksport/gift/internal/domain/customer/mongo"
	"github.com/mansoorceksport/gift/internal/domain/gift"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var giftColl = "gift"

type GiftRepository struct {
	db   *mongo.Database
	gift *mongo.Collection
}

func NewGiftMongo(db *mongo.Database) gift.Repository {
	giftCollection := db.Collection(giftColl)
	return &GiftRepository{
		db:   db,
		gift: giftCollection,
	}
}

type mongoGiftItem struct {
	ID                uuid.UUID `bson:"id"`
	Name              string    `bson:"name"`
	Description       string    `bson:"description"`
	QuantityRequested int       `bson:"quantityRequested"`
	QuantityLocked    int       `bson:"quantityLocked"`
	QuantityReleased  int       `bson:"quantityReleased"`
}

type mongoGift struct {
	ID       uuid.UUID                   `bson:"id"`
	Name     string                      `bson:"name"`
	Customer mongoCustomer.MongoCustomer `bson:"customer"`
	Items    []mongoGiftItem             `bson:"items"`
}

func newGiftMongo(g gift.Gift) mongoGift {
	var items []mongoGiftItem
	for _, p := range g.GetProducts() {
		items = append(items, mongoGiftItem{
			ID:                p.Id,
			Name:              p.Name,
			Description:       p.Description,
			QuantityRequested: p.QuantityRequested,
			QuantityLocked:    p.QuantityLocked,
			QuantityReleased:  p.QuantityReleased,
		})
	}

	return mongoGift{
		ID:    g.GetID(),
		Name:  g.GetName(),
		Items: items,
		Customer: mongoCustomer.MongoCustomer{
			ID:          g.GetCustomer().ID,
			Name:        g.GetCustomer().Name,
			Description: g.GetCustomer().Description,
		},
	}
}

func (mg mongoGift) ToAggregate() gift.Gift {
	g := gift.Gift{}
	g.SetID(mg.ID)
	g.SetName(mg.Name)
	g.SetCustomer(&entity.Person{
		ID:          mg.Customer.ID,
		Name:        mg.Customer.Name,
		Description: mg.Customer.Description,
	})
	for _, item := range mg.Items {
		g.AddGiftItem(&entity.GiftItem{
			Id:                item.ID,
			Name:              item.Name,
			Description:       item.Description,
			QuantityRequested: item.QuantityRequested,
			QuantityLocked:    item.QuantityLocked,
			QuantityReleased:  item.QuantityReleased,
		})
	}
	return g
}

func (g GiftRepository) Create(gift gift.Gift) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := newGiftMongo(gift)
	_, err := g.gift.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (g GiftRepository) GetById(id uuid.UUID) (gift.Gift, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := g.gift.FindOne(ctx, bson.M{"id": id})
	var mg mongoGift
	err := result.Decode(&mg)
	if err != nil {
		return gift.Gift{}, err
	}

	return mg.ToAggregate(), nil
}
