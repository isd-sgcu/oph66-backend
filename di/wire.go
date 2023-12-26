//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/isd-sgcu/oph66-backend/cache"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"github.com/isd-sgcu/oph66-backend/database"
	event "github.com/isd-sgcu/oph66-backend/internal/event"
	featureflag "github.com/isd-sgcu/oph66-backend/internal/feature_flag"
	healthcheck "github.com/isd-sgcu/oph66-backend/internal/health_check"
	"github.com/isd-sgcu/oph66-backend/logger"
	"go.uber.org/zap"
)

type Container struct {
	EventHandler       event.Handler
	HcHandler          healthcheck.Handler
	FeatureflagHandler featureflag.Handler
	Config             *cfgldr.Config
	Logger             *zap.Logger
	CorsHandler        cfgldr.CorsHandler
}

func newContainer(eventHandler event.Handler, hcHandler healthcheck.Handler, featureflagHandler featureflag.Handler, config *cfgldr.Config, logger *zap.Logger, corsHandler cfgldr.CorsHandler) Container {
	return Container{
		eventHandler,
		hcHandler,
		featureflagHandler,
		config,
		logger,
		corsHandler,
	}
}

func Init() (Container, error) {
	wire.Build(
		newContainer,
		cfgldr.LoadConfig,
		cfgldr.MakeCorsConfig,
		event.NewHandler,
		event.NewService,
		event.NewRepository,
		event.NewCache,
		healthcheck.NewHandler,
		database.New,
		cache.New,
		featureflag.NewHandler,
		featureflag.NewCache,
		featureflag.NewService,
		featureflag.NewRepository,
		logger.InitLogger,
	)

	return Container{}, nil
}
