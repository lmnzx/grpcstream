package notificationservice

import (
	"fmt"
	"grpcstreaming/notificationservice/notificationproto"
	"log/slog"
	"time"

	"github.com/valkey-io/valkey-go"
	"google.golang.org/grpc"
)

var _ notificationproto.NotificationServiceServer = (*Handler)(nil)

type Handler struct {
	notificationproto.UnimplementedNotificationServiceServer
	valkeyClient valkey.Client
}

func NewHandler(valkeyClient valkey.Client) *Handler {
	return &Handler{
		valkeyClient: valkeyClient,
	}
}

func (h *Handler) GetNotification(req *notificationproto.NotificationRequest, stream grpc.ServerStreamingServer[notificationproto.Notification]) error {
	channelName := fmt.Sprintf("notifications/%s", req.GetUserId())
	ctx := stream.Context()
	sub := h.valkeyClient.B().Subscribe().Channel(channelName).Build()

	ch := make(chan valkey.PubSubMessage)
	errCh := make(chan error, 1)

	go func() {
		err := h.valkeyClient.Receive(ctx, sub, func(msg valkey.PubSubMessage) {
			slog.Info("received a new message ", msg.Message, msg.Channel)
			ch <- msg
		})
		if err != nil && ctx.Err() == nil {
			errCh <- err
		}
		close(ch)
		close(errCh)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if ok {
				err := stream.Send(&notificationproto.Notification{
					UserId:    req.GetUserId(),
					Content:   fmt.Sprintf("New Notification at %s: %s", time.Now().String(), msg.Message),
					CreatedAt: time.Now().UnixMilli(),
				})
				if err != nil {
					return fmt.Errorf("could not send notification: %s", err)
				}
			}
		case err := <-errCh:
			if err != nil {
				return fmt.Errorf("failed subsciption: %s", err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
