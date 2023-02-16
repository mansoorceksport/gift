package mongo

import (
	"context"
	"fmt"
	"github.com/mansoorceksport/gift/internal/domain/common/database/mongo"
	"github.com/mansoorceksport/gift/internal/domain/customer"
	"testing"
)

func TestNewCustomerRepository(t *testing.T) {
	ctx := context.Background()
	m, err := mongo.NewMongo(ctx, "mongodb://localhost:55575")
	if err != nil {
		return
	}
	customerRepository := NewCustomerRepository(m)
	cus, err := customer.NewCustomer("mansoor")
	cus.SetDescription("my name is mansoor")
	if err != nil {
		t.Fatal(err)
	}
	err = customerRepository.Add(cus)
	if err != nil {
		t.Fatal(err)
	}

	c, err := customerRepository.GetById(cus.GetID())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("customer ID : %s \n", cus.GetID())

	if c.GetID() != cus.GetID() {
		t.Fatal("failed to add customer")
	}

}
