package customer

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/gift/internal/domain/common/entity"
)

// Customer Aggregate is the combination of entities and value objects.
// business logic of customer needs to be inside the aggregate
type Customer struct {
	person *entity.Person
}

func NewCustomer(name string) (Customer, error) {
	return Customer{
		person: &entity.Person{
			ID:          uuid.New(),
			Name:        name,
			Description: "",
		},
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetDescription() string {
	return c.person.Description
}

func (c *Customer) SetDescription(description string) {
	c.person.Description = description
}

func (c *Customer) GetEntity() *entity.Person {
	return c.person
}
