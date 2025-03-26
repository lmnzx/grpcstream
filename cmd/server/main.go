package main

import (
	"context"
	"grpcstreaming"
	"grpcstreaming/notificationservice"
	"grpcstreaming/notificationservice/notificationproto"
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", notificationservice.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	ctx := context.Background()

	valkeyClient := grpcstreaming.NewValkeyClient(ctx)
	handler := notificationservice.NewHandler(valkeyClient)

	notificationproto.RegisterNotificationServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	slog.Info("listening on " + notificationservice.Address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
