package message

import (
	"log/slog"
	"time"
)

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

func (l *Log) LogValue() slog.Value {
	if l == nil {
		return slog.Value{}
	}

	attrs := []slog.Attr{
		slog.String("Level", l.Level),
		slog.String("Message", l.Message),
		slog.String("ResourceID", l.ResourceID),
		slog.Time("Timestamp", l.Timestamp),
		slog.String("TraceID", l.TraceID),
		slog.String("SpanID", l.SpanID),
		slog.String("Commit", l.Commit),
	}
	if l.Metadata.ParentResourceID != nil {
		attrs = append(attrs,
			slog.Any("Metadata",
				slog.GroupValue([]slog.Attr{
					slog.String("ParentResourceID", *l.Metadata.ParentResourceID)}...,
				),
			),
		)
	}

	return slog.GroupValue(attrs...)
}
