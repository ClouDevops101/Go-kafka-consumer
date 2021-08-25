package main

import (

	// kafka
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("Kafka Consumer using Golang")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"0.0.0.0:9092"},
		Topic:   "sp-user",
		GroupID: "sp-user-1",
	})
	defer reader.Close()

	fmt.Println("Consumer is runing...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Time: %s |  Topic: %s | Key : %s |  Message : %s | Partition :%d \n", msg.Time, msg.Topic, msg.Key, string(msg.Value), msg.Partition)
	}
}
