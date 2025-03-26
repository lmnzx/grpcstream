package main

import (
	"context"
	"fmt"
	"grpcstreaming"
	"log"
)

func main() {
	ctx := context.Background()

	valkeyClient := grpcstreaming.NewValkeyClient(ctx)

	channelName := fmt.Sprintf("notifications/%s", "123")
	message := "hello world"

	fmt.Printf("Publishing message '%s' to channel '%s'\n", message, channelName)

	err := valkeyClient.Do(ctx, valkeyClient.B().Publish().Channel(channelName).Message(message).Build()).Error()
	if err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}
}
