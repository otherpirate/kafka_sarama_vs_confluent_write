package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	start := time.Now()
	fmt.Println(start)
	for i := 1; i <= 10000; i++ {
		msg := &sarama.ProducerMessage{
			Topic: "benchmark_topic",
			Value: sarama.StringEncoder(i),
		}
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
	}

	end := time.Now()
	fmt.Println(end)
	elapsed := end.Sub(start)
	fmt.Println(elapsed)
}
