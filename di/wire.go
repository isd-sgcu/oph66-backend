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
	login "github.com/isd-sgcu/oph66-backend/internal/login"
	register "github.com/isd-sgcu/oph66-backend/internal/register"
	"github.com/isd-sgcu/oph66-backend/logger"
	"go.uber.org/zap"
)

type Container struct {
	EventHandler       event.Handler
	HcHandler          healthcheck.Handler
	FeatureflagHandler featureflag.Handler
	RegisterHandler    register.Handler
	LoginHandler       login.Handler
	Config             *cfgldr.Config
	Logger             *zap.Logger
}

func newContainer(eventHandler event.Handler, hcHandler healthcheck.Handler, featureflagHandler featureflag.Handler, RegisterHandler register.Handler, LoginHandler login.Handler, config *cfgldr.Config, logger *zap.Logger) Container {
	return Container{
		eventHandler,
		hcHandler,
		featureflagHandler,
		RegisterHandler,
		LoginHandler,
		config,
		logger,
	}
}

func Init() (Container, error) {
	wire.Build(
		newContainer,
		cfgldr.LoadConfig,
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
		register.NewHandler,
		register.NewService,
		register.NewRepository,
		login.NewHandler,
		login.NewService,
		login.NewRepository,
		logger.InitLogger,
	)

	return Container{}, nil
}
