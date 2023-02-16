package collection

import "github.com/google/uuid"

type MongoItem struct {
	ID          uuid.UUID `bson:"id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Quantity    int       `bson:"quantity"`
}

// bulk and
// customer reliability (with clipboard
// form mirroring
// akan di damping advocate untuk minimalisir urusan post audit
