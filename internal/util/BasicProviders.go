package util

import (
	"fmt"
	"github.com/caarlos0/env"
	"go.uber.org/zap"
	"net/http"
)

var (
	// Logger is the defaut logger
	logger *zap.SugaredLogger
	//FIXME: remove this
	//defer Logger.Sync()
)

// Deprecated: instead calling this method inject logger from wire
func GetLogger() *zap.SugaredLogger {
	return logger
}

type SentryConfig struct {
	DSN           string `env:"DSN" envDefault:"https://0e57adc987494c3f99c56e8287475c20@sentry.io/1887839"`
	SentryEnv     string `env:"SENTRY_ENV" envDefault:"Staging"`
	SentryEnabled bool   `env:"SENTRY_ENABLED" envDefault:"false"`
}

func init() {
	cfg := &SentryConfig{}
	err := env.Parse(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := zap.NewProduction()
	if err != nil {
		panic("failed to create the default logger: " + err.Error())
	}
	if cfg.SentryEnabled {
		logger = l.Sugar() //modifyToSentryLogger(l, cfg.DSN, cfg.SentryEnv)
	} else {
		logger = l.Sugar()
	}
}

func NewSugardLogger() *zap.SugaredLogger {
	return logger
}

func NewHttpClient() *http.Client {
	return http.DefaultClient
}
