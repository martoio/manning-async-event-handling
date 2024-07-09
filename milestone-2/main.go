package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/martoio/manning-async-event-handling/events"
	"github.com/martoio/manning-async-event-handling/models"
	"github.com/martoio/manning-async-event-handling/publisher"
)

type OrderReceivedMessageBody struct {
}

type OrderReceivedMessage struct {
	Id        int                      `json:"id"`
	Name      string                   `json:"name"`
	Timestamp time.Time                `json:"timestamp"`
	Body      OrderReceivedMessageBody `json:"body"`
}

func main() {

	orderStore := models.OrderStore{
		Orders: make([]models.Order, 0),
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	r.Post("/orders", func(w http.ResponseWriter, r *http.Request) {
		type Body struct {
			Name       string `json:"name"`
			ProductIds []int  `json:"productIds"`
			CustomerId int    `json:"customerId"`
		}
		var orderBody Body
		err := json.NewDecoder(r.Body).Decode(&orderBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newOrder := orderStore.NewOrder(models.NewOrderParams{
			Name:       orderBody.Name,
			ProductIds: orderBody.ProductIds,
			CustomerId: orderBody.CustomerId,
		})
		publisher.PublishEvent(events.OrderReceivedEvent{
			EventId:        newOrder.ID,
			EventTimestamp: newOrder.CreatedAt,
			EventBody:      *newOrder,
		}, events.OrdersReceivedTopic)

	})
	fmt.Println("Listening on port 3100")
	http.ListenAndServe(":3100", r)
}
