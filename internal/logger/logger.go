package logger

import (
	"go.uber.org/zap"
	"net/http"
)

func NewSugardLogger() *zap.SugaredLogger {
	l, err := zap.NewProduction()
	if err != nil {
		panic("failed to create the default logger: " + err.Error())
	}
	return l.Sugar()
}

func NewHttpClient() *http.Client {
	return http.DefaultClient
}