package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	KafkaServer1 = "localhost:9092"
	KafkaServer2 = "localhost:9093"
	KafkaTopic   = "orders-v1-topic"
	KafkaGroupId = "product-service-2"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s, %s", KafkaServer1, KafkaServer2),
		"group.id":          KafkaGroupId,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	topic := KafkaTopic
	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var order Order
			err := json.Unmarshal(msg.Value, &order)
			if err != nil {
				fmt.Printf("Error decoding message: %v\n", err)
				continue
			}

			fmt.Printf("Received Order: %+v\n", order)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

type Order struct {
	ID        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Amount    int64  `json:"amount"`
}
