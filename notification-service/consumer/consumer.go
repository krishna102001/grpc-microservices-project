package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"

	mailservice "github.com/krishna102001/grpc-microservices-project/notification-service/mail-service"
	typesdata "github.com/krishna102001/grpc-microservices-project/notification-service/types-data"
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
	// fmt.Println("message are ", string(msg.Value))
	var data typesdata.KafkaMessage
	if err := json.Unmarshal(msg.Value, &data); err != nil {
		log.Fatalf("Failed to unmarshal the json %v", err)
	}
	fmt.Printf("data service-type(%s) | message-type(%s) | message-content(%s)\n", data.ServiceType, data.MessageType, data.MessageContent)
	switch data.ServiceType {
	case "email":
		log.Printf("email service called")
		if err := mailservice.SendMail([]string{data.MessageContent.RecieverEmail}, data); err != nil {
			log.Println("Planning to template the retry mechanism save in database")
			//  save to email database
			// run for loop forever and check every second if any email is in db
			// if have then capture that email and send the mail again if succeed then delete it from the table
		}
	case "sms":
		log.Printf("sms service called")
	default:
		log.Printf("invalid service type please check")
	}
}
