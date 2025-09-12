package typesdata

type KafkaMessage struct {
	ServiceType    string         `json:"service_type"`
	MessageType    string         `json:"message_type"`
	MessageContent messageContent `json:"message_content"`
}

type messageContent struct {
	RecieverEmail string `json:"reciever_email"`
	Content       string `json:"content"`
}
