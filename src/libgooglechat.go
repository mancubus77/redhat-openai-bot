package src

import (
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/chat/v1"
)

var (
	incomingMessage *chat.DeprecatedEvent
	responseMessage *chat.Message
)

// Read queue and pass received message to ChatGTP
func QueueReader(cctx context.Context,
	sub *pubsub.Subscription,
	sms *chat.SpacesMessagesService) {

	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("Received Message %s.\n", string(msg.Data))
		msg.Ack()

		err := json.Unmarshal(msg.Data, &incomingMessage)

		// var result map[string]interface{}
		// json.Unmarshal(msg.Data, &result)
		// message := gjson.Get(string(msg.Data), "message.text")

		message := incomingMessage.Message.Text

		fmt.Printf(">>> %v", RemoveName(message))
		if err != nil {
			msg.Ack()
			log.Fatalf("Unable to decode Chat Message JSON: %v.\n", err)
		}
		responseMessage = new(chat.Message)
		gtpResponse := GetResponse(message)
		responseMessage.Text = gtpResponse

		response, err := sms.Create(incomingMessage.Space.Name, responseMessage).Do()
		if err != nil {
			log.Printf("There was an error sending a response back to Hangouts Chat: %v.\n", err)
		}
		log.Printf("Hangouts Response: %+v.\n", response)
	})

	if err != nil {
		log.Panicf("Error, can not subscribe: %v", err)
	}

}
