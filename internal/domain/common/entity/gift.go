package entity

import "github.com/google/uuid"

type Gift struct {
	ID   uuid.UUID
	Name string
}

type GiftItem struct {
	Id                uuid.UUID
	Name              string
	Description       string
	QuantityRequested int
	QuantityLocked    int
	QuantityReleased  int
}
