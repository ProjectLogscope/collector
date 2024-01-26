package main

import (
	"context"

	"github.com/hardeepnarang10/collector/pkg/queue"
)

func initQueue(ctx context.Context, endpoint string) queue.Queue {
	q, err := queue.New(endpoint)
	if err != nil {
		report(ctx, err)
	}
	return q
}
