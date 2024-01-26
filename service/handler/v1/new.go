package handler

import (
	"errors"

	"github.com/hardeepnarang10/collector/service/registry"
)

type handler struct {
	service *registry.ServiceRegistry
}

func New(svc *registry.ServiceRegistry) (Handler, error) {
	if svc == nil {
		return nil, errors.New("handler: empty service registery passed")
	}

	return &handler{
		service: svc,
	}, nil
}
