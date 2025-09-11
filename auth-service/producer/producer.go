package producer

import (
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
)

var kafkaProducerClient sarama.SyncProducer

func connectKafkaProducer() error {
	brk_url := os.Getenv("BROKER_URL")
	if brk_url == "" {
		log.Fatal("broker-url not found")
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer([]string{brk_url}, config)
	if err != nil {
		return err
	}

	kafkaProducerClient = conn
	log.Println("--------------Kafka new Producer is created---------------------")
	return nil
}

func PushToKafkaQueue(topic string, message []byte) error {
	if err := connectKafkaProducer(); err != nil {
		log.Fatalf("Failed to connect to new producer %v", err)
		return err
	}
	log.Println("----------------Kafka connected successfully---------------------")
	defer kafkaProducerClient.Close()
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(message),
		Timestamp: time.Now(),
	}

	partition, offset, err := kafkaProducerClient.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message sent successfully to the topics(%s) | partition(%d) | offset(%d) \n", topic, partition, offset)
	return nil
}
