package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092,localhost:9093,localhost:9094"})
	if err != nil {
		panic(err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "benchmark_topic"
	start := time.Now()
	fmt.Println(start)
	for i := 1; i <= 10000; i++ {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(strconv.Itoa(i)),
		}, nil)
	}

	// Wait for message deliveries
	p.Flush(15 * 1000)

	end := time.Now()
	fmt.Println(end)
	elapsed := end.Sub(start)
	fmt.Println(elapsed)
}
