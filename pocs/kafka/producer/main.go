package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					continue
				}
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "testTopic"
	for i := 0; i < 100; i++ {

		err := p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(fmt.Sprintf("hello world from #%d", i)),
		}, nil)
		if err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
