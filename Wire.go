//+build wireinject

package main

import (
	"github.com/devtron-labs/external-app-crawler/api"
	"github.com/devtron-labs/external-app-crawler/client"
	"github.com/devtron-labs/external-app-crawler/internal/logger"
	"github.com/devtron-labs/external-app-crawler/internal/sql"
	"github.com/devtron-labs/external-app-crawler/pubsub"
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		NewApp,
		api.NewMuxRouter,
		logger.NewSugardLogger,
		sql.GetConfig,
		sql.NewDbConnection,
		api.NewRestHandlerImpl,
		wire.Bind(new(api.RestHandler), new(*api.RestHandlerImpl)),
		client.NewPubSubClient,
		pubsub.NewNatSubscription,

	)
	return &App{}, nil
}
