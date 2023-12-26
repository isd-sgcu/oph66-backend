package logger

import (
	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"go.uber.org/zap"
)

func InitLogger(cfg *cfgldr.Config) *zap.Logger {
	var logger *zap.Logger

	if cfg.AppConfig.IsDevelopment() {
		logger = zap.Must(zap.NewDevelopment())
	} else {
		logger = zap.Must(zap.NewProduction())
	}

	return logger
}
