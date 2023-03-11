package main

import (
	"flag"
	"log"
	"os"

	"github.com/mancubus77/red-hat-openai-bot/src"
	"golang.org/x/net/context"
)

var (
	credentialsFile = flag.String("credentialsFile", "", "path to credentials file")
	project         = flag.String("project", "", "Google Cloud Project name")
	psTopic         = flag.String("topic", "", "Pub/Sub Topic for Hangouts Chat")
	psSubscription  = flag.String("subscription", "", "Pub/Sub Subscription for Hangouts Chat")
	apiKey          = flag.String("gtp-key", "", "ChatGTP3 API Key")
)

func InitClients() {
	log.Println("Initializing GTP3 Chat client")
	src.InitGTPClient(*apiKey)

}

func main() {
	flag.Parse()
	log.Println("Hangbot Starting.")
	InitClients()
	log.Printf("Configuration: Credentials File: %s, Project: %s, Topic: %s, Subscription: %s.\n", *credentialsFile, *project, *psTopic, *psSubscription)

	// This seems like a hack, but some of the oauth libraries expect an environment variable
	// if you use the JSON file, as opposed to being able to specify the path
	// as part of client creation.
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", *credentialsFile)
	ctx := context.Background()

	// Init PubSub client
	client := src.InitQueue(ctx, project, credentialsFile)

	// Get subscription
	sub := src.SubscribeTopic(&client, psSubscription)

	// Init google chat and contexts
	cctx, sms := src.InitGoogleChat(ctx, sub)

	// Read Queue and process messages
	src.QueueReader(cctx, sub, sms)

	log.Println("Exiting.")
}
