package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// producer config

	config := &kafka.ConfigMap{

		"bootstrap.servers": "localhost:9092", // kafka broker addr

	}

	topic := "coordinates" // target topic name

	// Create Kafka producer

	producer, err := kafka.NewProducer(config)

	if err != nil {

		panic(err)

	}

	// write 10 messages to the "coordinates" topic

	for i := 0; i < 10; i++ {

		value := fmt.Sprintf("message-%d", i)

		err := producer.Produce(&kafka.Message{

			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},

			Value: []byte(value),
		}, nil)

		if err != nil {

			fmt.Printf("Failed to produce message %d: %v\n", i, err)

		} else {

			fmt.Printf("Produced message %d: %s\n", i, value)

		}

	}

	// Close Kafka producer

	producer.Flush(15 * 1000)

	producer.Close()

}
