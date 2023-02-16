package mongo

import (
	"context"
	"github.com/google/uuid"
	customer2 "github.com/mansoorceksport/gift/internal/domain/customer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var customerColl = "customer"

type CustomerRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// MongoCustomer is internal type that is used to store a CustomerAggregate
// inside this repository.
type MongoCustomer struct {
	ID          uuid.UUID `bson:"id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
}

func NewCustomerRepository(db *mongo.Database) customer2.Repository {
	customerCollection := db.Collection(customerColl)
	return CustomerRepository{
		db:       db,
		customer: customerCollection,
	}
}

func NewFromCustomer(c customer2.Customer) MongoCustomer {
	return MongoCustomer{
		ID:          c.GetID(),
		Name:        c.GetName(),
		Description: c.GetDescription(),
	}
}

func (m MongoCustomer) ToAggregate() customer2.Customer {
	c := customer2.Customer{}
	c.SetName(m.Name)
	c.SetID(m.ID)
	c.SetDescription(m.Description)
	return c
}

func (c CustomerRepository) GetById(id uuid.UUID) (customer2.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := c.customer.FindOne(ctx, bson.M{"id": id})
	var mc MongoCustomer
	err := result.Decode(&mc)
	if err != nil {
		return customer2.Customer{}, err
	}

	return mc.ToAggregate(), nil
}

func (c CustomerRepository) Add(customer customer2.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := NewFromCustomer(customer)
	_, err := c.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (c CustomerRepository) Remove(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.customer.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (c CustomerRepository) Update(customer customer2.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := NewFromCustomer(customer)
	_, err := c.customer.UpdateByID(ctx, customer.GetID(), internal)
	if err != nil {
		return err
	}
	return nil
}
