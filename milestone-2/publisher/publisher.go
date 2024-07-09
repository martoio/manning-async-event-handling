package publisher

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/martoio/manning-async-event-handling/events"
)

func PublishEvent(event events.Event, topic string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":   "localhost:9092",
		"socket.timeout.ms":   30_000,
		"delivery.timeout.ms": 30_000,
	})
	if err != nil {
		return err
	}

	deliveryChan := make(chan kafka.Event)

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	close(deliveryChan)
	return nil
}
