package consumer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

var kafka_client sarama.Consumer

func Connect(brokerURL []string, topic string, wg *sync.WaitGroup) error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokerURL, config)
	if err != nil {
		return err
	}
	kafka_client = conn
	consume_parition(topic, wg)
	return nil
}

func consume_parition(topic string, wg *sync.WaitGroup) {
	consume, err := kafka_client.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Failed to consume the partition %v \n", err)
	}
	log.Printf("Consumer started consuming the topics %v \n", topic)

	wg.Add(1)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// doneCh := make(chan struct{})

	totalNotificationCount := 0
	go func() {
		for {
			select {
			case err := <-consume.Errors():
				fmt.Println(err)
			case msg := <-consume.Messages():
				totalNotificationCount++
				sendMessage(msg)
			case <-sigchan:
				// doneCh <- struct{}{}
				wg.Done()
				fmt.Printf("\nTotal Notification recieved are %d \n", totalNotificationCount)
				if err := consume.Close(); err != nil {
					log.Fatalf("Failed to close the kafka service")
				}
				return
			}
		}
	}()
	// <-doneCh

}

func sendMessage(msg *sarama.ConsumerMessage) {
	fmt.Printf("Recieved the message from topics(%s) | paritions(%v) | Time(%v) \n", msg.Topic, msg.Partition, msg.Timestamp)
	fmt.Println("message are ", string(msg.Value))
}
