package config

import (
	"time"

	"github.com/alexflint/go-arg"
)

type config struct {
	ServiceName            string          `arg:"--service-name,env:SERVICE_NAME" help:"Service name" default:"collector" placeholder:"collector"`
	ServiceLogLevel        ServiceLogLevel `arg:"--service-log-level,env:SERVICE_LOG_LEVEL" help:"Specify log level" default:"info" placeholder:"info"`
	ServiceLogFilepath     string          `arg:"--service-log-filepath,env:SERVICE_LOG_FILEPATH" help:"Relative filepath for logger output" default:"/service/log/collector.log" placeholder:"/service/log/collector.log"`
	ServiceLogAddSource    bool            `arg:"--service-log-add-source,env:SERVICE_LOG_ADD_SOURCE" help:"Add source info to logger output" default:"true" placeholder:"true"`
	ServiceRequestTimeout  time.Duration   `arg:"--service-request-timeout,env:SERVICE_REQUEST_TIMEOUT" help:"Service request timeout duration" default:"30s" placeholder:"30s"`
	ServiceRequestValidate bool            `arg:"--service-request-validate,env:SERVICE_REQUEST_VALIDATE" help:"Enable request validation" default:"true" placeholder:"true"`
	ServiceGracePeriod     time.Duration   `arg:"--service-grace-period,env:SERVICE_GRACE_PERIOD" help:"Service shutdown grace period" default:"1m" placeholder:"1m"`
	ServicePreforkEnable   bool            `arg:"--service-prefork-enable,env:SERVICE_PREFORK_ENABLE" help:"Enable service prefork" default:"true" placeholder:"true"`

	ServerHostname string `arg:"--server-hostname,env:SERVER_HOSTNAME" help:"Specify service hostname" default:"localhost" placeholder:"localhost"`
	ServerHTTPPort uint16 `arg:"--server-http-port,env:SERVER_HOSTPORT" help:"Port to listen on" default:"3000" placeholder:"3000"`

	NATSServerEndpoint string `arg:"--nats-server-endpoint,env:NATS_SERVER_ENDPOINT" help:"NATS server endpoint" placeholder:"nats://localhost:4222"`
	NATSSubjectIngest  string `arg:"--nats-subject-ingest,env:NATS_SUBJECT_INGEST" help:"NATS ingest subject" placeholder:"subject-ingest"`
}

func (config) Version() string {
	return "Log Ingest - Collector v1.0"
}

func LoadConfig() config {
	var cfg config
	p := arg.MustParse(&cfg)
	validateStrings(cfg, p)
	return cfg
}
