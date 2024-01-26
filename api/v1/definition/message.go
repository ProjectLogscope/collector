package definition

import "time"

type Log struct {
	Level      string    `json:"level" validate:"required"`
	Message    string    `json:"message" validate:"required"`
	ResourceID string    `json:"resourceId" validate:"required"`
	Timestamp  time.Time `json:"timestamp" validate:"required"`
	TraceID    string    `json:"traceId" validate:"required"`
	SpanID     string    `json:"spanId" validate:"required"`
	Commit     string    `json:"commit" validate:"required"`
	Metadata   struct {
		ParentResourceID *string `json:"parentResourceId,omitempty"`
	} `json:"metadata" validate:"required"`
}
