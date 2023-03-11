package src

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
)

// Init PubSub client
func InitQueue(ctx context.Context, project *string, credentialsFile *string) pubsub.Client {

	fmt.Println("Creating PubSub client")
	client, err := pubsub.NewClient(ctx, *project, option.WithCredentialsFile(*credentialsFile))

	if err != nil {
		log.Fatalf("error creating newclient: %v.\n", err)
	}
	return *client
}

// Init Subscribe to pubSub topic
func SubscribeTopic(client *pubsub.Client, psSubscription *string) *pubsub.Subscription {

	fmt.Println("Subscribing to the topic")
	sub := client.Subscription(*psSubscription)

	return sub
}

// Init Google chat client
func InitGoogleChat(ctx context.Context, sub *pubsub.Subscription) (context.Context, *chat.SpacesMessagesService) {
	fmt.Println("Creating google chat client")
	chatService, err := chat.NewService(ctx)
	if err != nil {
		log.Fatalf("Error creating chatService: %v.\n", err)
	}

	sms := chat.NewSpacesMessagesService(chatService)

	cctx, _ := context.WithCancel(ctx)

	ok, err := sub.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking if subscription exists. Err: %v.", err)
	}
	if !ok {
		log.Fatalln("Checked if subscription exists. It doesn't.")
	}

	return cctx, sms

}
