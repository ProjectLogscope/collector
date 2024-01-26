package registry

import (
	"github.com/hardeepnarang10/collector/pkg/queue"
)

type ServiceRegistry struct {
	ServiceConfig ServiceConfig
	QueueClient   queue.Queue
}
