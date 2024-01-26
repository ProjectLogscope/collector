package queue

import "context"

type Queue interface {
	Publish(ctx context.Context, subject string, body []byte) error
}
