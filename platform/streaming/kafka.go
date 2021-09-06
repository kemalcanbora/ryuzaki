package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	helpers "ryuzaki/helpers"
	schema "ryuzaki/helpers/schema"
	"time"
)

func producer(topic string) {
	helpers.GetEnv()

	kafkaHost := os.Getenv("DOCKER_KAFKA_HOST")
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	dummyData := &schema.Comment{
		Text:   "tes",
		Bos:    "pear",
		Fruits: []string{"a", "b", "c"}}

	res1B, _ := json.Marshal(dummyData)

	_, err = conn.WriteMessages(
		kafka.Message{Value: res1B},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func main() {
	producer("jungle")
}
