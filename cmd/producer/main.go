package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
    p, err := kafka.NewProducer(&kafka.ConfigMap{
       "bootstrap.servers": "localhost", 
    })
    if err != nil {
        panic(err)
    }
    defer p.Close()

    topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	p.Flush(15 * 1000)

    fmt.Println("Successfully created producer")
}
