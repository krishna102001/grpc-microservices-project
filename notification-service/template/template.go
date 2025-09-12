package template

import (
	"log"
	"os"

	typesdata "github.com/krishna102001/grpc-microservices-project/notification-service/types-data"
)

var login_subject_template = `You have login successfully`
var login_body_template = `Welcome To Microservices Blog Once Again Techie...`

var register_subject_template = `You have register successfully`
var register_body_template = `Welcome To Microservice Blog. Read like you never read this blog again & again`

func Email_Template(message typesdata.KafkaMessage) []byte {
	senderEmail := os.Getenv("FROM_SENDER_EMAIL")
	if senderEmail == "" {
		log.Fatalf("Failed to construct the template")
	}
	mp := make(map[string]string)

	mp["login"] = constructMailBody(senderEmail, message.MessageContent.RecieverEmail, login_subject_template, login_body_template)
	mp["register"] = constructMailBody(senderEmail, message.MessageContent.RecieverEmail, register_subject_template, register_body_template)

	return []byte(mp[message.MessageType])
}

func constructMailBody(senderEmail, recieverEmail, subject, msgBody string) string {
	var header string = "From: " + senderEmail + "\n" + "To: " + recieverEmail + "\n" + "Subject: " + subject + "\n"
	var body string = "Content-Type: text/html; charset=UTF-8\n\n" +
		"<div style='text-align: center; padding: 20px; background-color: #f4f4f4; font-family: Arial, sans-serif;'> " +
		"<div style='max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);'>" +
		"<h1 style='font-weight: bold; color: #ffffff; background-color: #0073e6; padding: 10px; border-radius: 8px;'>" + msgBody + "</h1>" +
		"</div>" +
		"</div>"
	totalMsg := header + body
	return totalMsg
}
