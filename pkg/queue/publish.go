package queue

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/nats-io/nats.go"
)

func (q *queue) Publish(ctx context.Context, subject string, body []byte) error {
	msg := nats.NewMsg(subject)
	msg.Data = body
	if err := q.conn.PublishMsg(msg); err != nil {
		slog.ErrorContext(ctx, "unable to publish message", slog.String("subject", subject), slog.Any("error", err))
		return fmt.Errorf("unable to publish message on subject %q: %w", msg.Subject, err)
	}
	return nil
}
