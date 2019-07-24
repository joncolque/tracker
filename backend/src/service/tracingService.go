package service

import (
	"fmt"
	"log"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func InitTracing(id_tracked string) {
	fmt.Println("id_tracked", id_tracked)
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
	}

	ctx := context.Background()

	client, err := app.Messaging(ctx)
	if err != nil {
			log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
			Data: map[string]string{
					"score": "1234",
					"time":  "4567",
			},
			Token: id_tracked, // This registration token comes from the client FCM SDKs.
			Notification: &messaging.Notification{
				Title: "Hello2",
				Body:  "World",
			},
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
			log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

}
