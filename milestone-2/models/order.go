package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID
	Name       string
	ProductIds []int
	CustomerId int
	CreatedAt  time.Time
}

type OrderStore struct {
	Orders []Order
}

type NewOrderParams struct {
	Name       string
	ProductIds []int
	CustomerId int
}

func (store *OrderStore) NewOrder(params NewOrderParams) *Order {
	id := uuid.New()

	order := Order{
		ID:         id,
		Name:       params.Name,
		ProductIds: params.ProductIds,
		CustomerId: params.CustomerId,
		CreatedAt:  time.Now(),
	}

	store.Orders = append(store.Orders, order)

	return &order
}
