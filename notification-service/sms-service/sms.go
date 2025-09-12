package smsservice

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(phoneNumber, message string) error {

	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}

	params.SetBody("Your OTP to reset the password is " + message + "\n -gRPC Microservices ")

	sent_from := os.Getenv("SENT_FROM")
	if sent_from == "" {
		log.Fatalf("Sent from phone number not found in env")
	}

	params.SetFrom(sent_from)
	params.SetTo(phoneNumber)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatalf("Failed to send the sms %v", err)
		return err
	} else {
		if resp.Body != nil {
			fmt.Println(*resp.Body)
		} else {
			fmt.Println(*resp.Body)
		}
	}
	return nil
}
