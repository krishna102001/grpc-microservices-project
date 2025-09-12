package mailservice

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/krishna102001/grpc-microservices-project/notification-service/template"
	typesdata "github.com/krishna102001/grpc-microservices-project/notification-service/types-data"
)

func connectMailServer(username, password, host string) smtp.Auth {
	return smtp.PlainAuth("", username, password, host)
}

func SendMail(recieverEmail []string, message typesdata.KafkaMessage) error {
	username, password, host, email_port := os.Getenv("FROM_SENDER_EMAIL"), os.Getenv("FROM_SENDER_PASSWORD"), os.Getenv("HOST_EMAIL"), os.Getenv("EMAIL_PORT")
	if username == "" || password == "" || host == "" {
		return fmt.Errorf("failed to get the username or password or host")
	}

	auth := connectMailServer(username, password, host)
	fmt.Println("Mail server is connected successfully")

	msg := template.Email_Template(message)

	if err := smtp.SendMail(host+":"+email_port, auth, username, recieverEmail, msg); err != nil {
		return fmt.Errorf("failed to send the email")
	}

	fmt.Println("Mail is send successfully")
	return nil
}
