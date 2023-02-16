package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/gift/internal/domain/common/database/collection"
	"github.com/mansoorceksport/gift/internal/domain/common/entity"
	"github.com/mansoorceksport/gift/internal/domain/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var productColl = "product"

type ProductMongoRepository struct {
	db      *mongo.Database
	product *mongo.Collection
}

type mongoProduct struct {
	Item collection.MongoItem `bson:"item"`
}

func NewProductMongoRepository(db *mongo.Database) ProductMongoRepository {
	productCollection := db.Collection(productColl)
	return ProductMongoRepository{
		db:      db,
		product: productCollection,
	}
}

func newProductMongo(product product.Product) mongoProduct {
	return mongoProduct{
		Item: collection.MongoItem{
			ID:          product.GetItem().ID,
			Name:        product.GetItem().Name,
			Description: product.GetItem().Description,
			Quantity:    product.GetItem().Quantity,
		},
	}
}

func (p mongoProduct) ToAggregate() product.Product {
	pp := product.Product{}
	pp.SetItem(&entity.Item{
		ID:          p.Item.ID,
		Name:        p.Item.Name,
		Description: p.Item.Description,
		Quantity:    p.Item.Quantity,
	})
	return pp
}

func (p ProductMongoRepository) GetById(id uuid.UUID) (product.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := p.product.FindOne(ctx, bson.M{"id": id})
	var mp mongoProduct
	err := result.Decode(&mp)
	if err != nil {
		return product.Product{}, err
	}

	return mp.ToAggregate(), nil
}

func (p ProductMongoRepository) Add(product product.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := newProductMongo(product)
	_, err := p.product.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductMongoRepository) Update(product product.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := newProductMongo(product)
	_, err := p.product.UpdateByID(ctx, product.GetId(), internal)
	if err != nil {
		return err
	}
	return nil
}
