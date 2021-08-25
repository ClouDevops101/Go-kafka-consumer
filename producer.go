package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("Kafka Producer using Golang")
	writer := kafka.Writer{
		Addr:  kafka.TCP("10.0.0.118:9092"),
		Topic: "sp-user",
	}
	defer writer.Close()

	for i := 100; i < 105; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprint("Key-%d", i+1)),
			Value: []byte(fmt.Sprint("Message : %d", i+1)),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("Message %d sent \n", i+1)
		time.Sleep(1 * time.Second)
	}
}
