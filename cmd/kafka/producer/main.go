package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

const (
	KafkaServer1 = "localhost:9092"
	KafkaServer2 = "localhost:9093"
	KafkaTopic   = "orders-v1-topic"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s, %s", KafkaServer1, KafkaServer2),
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	topic := KafkaTopic
	order := Order{
		ID:        uuid.New().String(),
		ProductId: uuid.New().String(),
		UserId:    uuid.New().String(),
		Amount:    456000,
	}

	value, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	if err != nil {
		panic(err)
	}
}

type Order struct {
	ID        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Amount    int64  `json:"amount"`
}
