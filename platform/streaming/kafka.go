package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	helpers "ryuzaki/helpers"
	"time"
)

func Producer(topic string, data []byte) {
	helpers.GetEnv()
	kafkaHost := os.Getenv("DOCKER_KAFKA_HOST")
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	_, err = conn.WriteMessages(
		kafka.Message{Value: data},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
