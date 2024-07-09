package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/martoio/manning-async-event-handling/models"
)

type OrderReceivedEvent struct {
	EventId        uuid.UUID `json:"id"`
	EventTimestamp time.Time `json:"timestamp"`
	EventBody      models.Order
}

func (or OrderReceivedEvent) ID() uuid.UUID {
	return or.EventId
}

func (or OrderReceivedEvent) Name() string {
	return "OrderReceived"
}

func (or OrderReceivedEvent) Timestamp() time.Time {
	return or.EventTimestamp
}

func (or OrderReceivedEvent) Body() interface{} {
	return or.EventBody
}
