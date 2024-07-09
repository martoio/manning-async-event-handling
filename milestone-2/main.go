package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	http.ListenAndServe(":3000", r)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "OrdersReceived"

	msg := OrderReceivedMessage{
		Id:        1,
		Name:      "Test",
		Timestamp: time.Now(),
		Body:      OrderReceivedMessageBody{},
	}

	msgJson, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msgJson,
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
