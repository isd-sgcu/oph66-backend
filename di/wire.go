//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/isd-sgcu/oph66-backend/cache"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"github.com/isd-sgcu/oph66-backend/database"
	auth "github.com/isd-sgcu/oph66-backend/internal/auth"
	event "github.com/isd-sgcu/oph66-backend/internal/event"
	featureflag "github.com/isd-sgcu/oph66-backend/internal/feature_flag"
	healthcheck "github.com/isd-sgcu/oph66-backend/internal/health_check"
	"github.com/isd-sgcu/oph66-backend/internal/middleware"
	"github.com/isd-sgcu/oph66-backend/internal/router"
	"github.com/isd-sgcu/oph66-backend/internal/evtreg"
	"github.com/isd-sgcu/oph66-backend/logger"
	"go.uber.org/zap"
)

type Container struct {
	EventHandler       event.Handler
	HcHandler          healthcheck.Handler
	FeatureflagHandler featureflag.Handler
	AuthHandler        auth.Handler
	EvtregHandler      evtreg.Handler
	Config             *cfgldr.Config
	Logger             *zap.Logger
	CorsHandler        cfgldr.CorsHandler
	Router             *router.Router
}

func newContainer(
	eventHandler event.Handler,
	hcHandler healthcheck.Handler,
	featureflagHandler featureflag.Handler,
	authHandler auth.Handler,
	evtregHandler evtreg.Handler,
	config *cfgldr.Config,
	logger *zap.Logger,
	corsHandler cfgldr.CorsHandler,
	router *router.Router,
) Container {
	return Container{
		eventHandler,
		hcHandler,
		featureflagHandler,
		authHandler,
		evtregHandler,
		config,
		logger,
		corsHandler,
		router,
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
		evtreg.NewRepository,
		evtreg.NewService,
		evtreg.NewHandler,
		auth.NewHandler,
		auth.NewService,
		auth.NewRepository,
		logger.InitLogger,
		router.NewRouter,
		middleware.NewAuthMiddleware,
	)

	return Container{}, nil
}
