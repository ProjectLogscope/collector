package queue

import (
	"fmt"

	"github.com/hardeepnarang10/collector/common/service"
	"github.com/nats-io/nats.go"
)

type queue struct {
	conn *nats.Conn
}

func New(endpoint string) (Queue, error) {
	conn, err := nats.Connect(endpoint, nats.Name(service.GetName()))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to NATS server: %w", err)
	}
	return &queue{
		conn: conn,
	}, nil
}
