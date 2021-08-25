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
		Addr:  kafka.TCP{"0.0.0.0:9092"},
		Topic: "goscreencasts",
	}
	defer writer.close()

	for i = 0; ; i++ {
		msg := kafka.Message{
			key:   []byte(fmt.Sprint("Key-%d", i+1)),
			Value: []byte(fmt.Sprint("Message : %d", i+1)),
		}
		err := writer.WriterMessage(context.Background(), msg)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("Message %d sent \n", i+1)
		time.Sleep(1 * time.Second)
	}
}
