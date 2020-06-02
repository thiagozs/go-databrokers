package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = c.Close()
	}()

	if err := c.SubscribeTopics([]string{
		"testTopic",
	}, nil); err != nil {
		fmt.Printf("Subscribe Topic error: %v\n", err)
		return
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}
}
