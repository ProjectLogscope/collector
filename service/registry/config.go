package registry

import "time"

type ServiceConfig struct {
	RequestTimeout   time.Duration
	EnableValidation bool
	QueueTopicLog    string
}
