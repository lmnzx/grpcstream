package main

import (
	"context"
	"encoding/json"
	"fmt"
	"grpcstreaming/notificationservice"
	"grpcstreaming/notificationservice/notificationproto"
	"io"
	"log"
)

func main() {
	client, err := notificationservice.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	stream, err := client.GetNotification(ctx, &notificationproto.NotificationRequest{
		UserId: "123",
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		notification, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("failed to read notification: %w", err)
		}
		b, err := json.MarshalIndent(notification, "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}
