package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/hardeepnarang10/collector/cmd/api/internal/config"
	"github.com/hardeepnarang10/collector/common/service"
	"github.com/hardeepnarang10/collector/service/registry"
)

func main() {
	cfg := config.LoadConfig()
	service.SetName(cfg.ServiceName)
	initLogger(cfg.ServiceLogFilepath, cfg.ServiceLogLevel.GetLevel(), cfg.ServiceLogAddSource)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	svc := &registry.ServiceRegistry{}
	svc.QueueClient = initQueue(ctx, cfg.NATSServerEndpoint)
	svc.ServiceConfig = registry.ServiceConfig{
		RequestTimeout:   cfg.ServiceRequestTimeout,
		EnableValidation: cfg.ServiceRequestValidate,
		QueueTopicLog:    cfg.NATSSubjectIngest,
	}

	server := initServer(ctx, svc, serverConfig{
		serverName:    service.GetName(),
		serverPort:    cfg.ServerHTTPPort,
		enablePrefork: cfg.ServicePreforkEnable,
	})
	go func() {
		if err := server.Start(cfg.ServerHTTPPort); err != nil {
			report(ctx, err)
		}
	}()
	defer func() {
		graceCtx, cancel := context.WithTimeout(context.Background(), cfg.ServiceGracePeriod)
		defer cancel()
		if err := server.Shutdown(graceCtx); err != nil {
			report(ctx, err)
		}
	}()
	<-ctx.Done()
}
