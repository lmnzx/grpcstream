package grpcstreaming

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/valkey-io/valkey-go"
)

func NewValkeyClient(ctx context.Context) valkey.Client {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{"127.0.0.1:6379"},
	})
	if err != nil {
		slog.Error(fmt.Errorf("error creating new valkey client: %v", err).Error())
	}

	if err := client.Do(ctx, client.B().Ping().Build()).Error(); err != nil {
		slog.Error(fmt.Errorf("failed to pind valkey: %v", err).Error())
	}

	return client
}
