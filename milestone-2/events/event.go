package events

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	ID() uuid.UUID
	Name() string
	Timestamp() time.Time
	Body() interface{}
}

const (
	OrdersReceivedTopic = "OrdersReceived"
)
