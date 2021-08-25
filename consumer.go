package main

import ( 
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gocql/gocql"
	"net/http"
	"time"
	"log"
    "golang.org/x/crypto/bcrypt"
	"regexp"
	// kafka
	"context"
	"github.com/segmentio/kafka-go"

)

func main()  {
	fmt.Println("Kafka Consumer using Golang")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"0.0.0.0:9092"},
		Topic: "goscreencasts",
	})
	defer reader.close()

	fmt.Println(Consumer is runing..."")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Topic : %s Message : %s Partition :%d \n", msg.Topic, string(msg.Value), msg.Partition)
	}
}