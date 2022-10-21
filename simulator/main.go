package main

import (
	"fmt"
	"log"

	kafka2 "github.com/codeedu/imersaofsfc2-simulator/application/kafka"
	route2 "github.com/codeedu/imersaofsfc2-simulator/application/route"
	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func main() {
	//teste()

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}

}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func teste() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("Ola", "readtest", producer)

	for {
		_ = 1
	}
}

func teste2() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[1])
}
