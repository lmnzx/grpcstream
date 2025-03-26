package notificationservice

import (
	"fmt"
	"grpcstreaming/notificationservice/notificationproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (notificationproto.NotificationServiceClient, error) {
	conn, err := grpc.NewClient(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to construct client: %v", err)
	}

	return notificationproto.NewNotificationServiceClient(conn), nil
}
