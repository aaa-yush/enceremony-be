package logger

import (
	"enceremony-be/internal/common/logger/conf"
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

func NewZapLogger(config *conf.LoggerConf) (*zap.SugaredLogger, error) {

	var cfg zap.Config

	switch strings.ToLower(config.Environment) {
	case "dev", "development", "stage":
		cfg = zap.NewDevelopmentConfig()
	case "prod", "production":
		cfg = zap.NewProductionConfig()
	default:
		return nil, errors.New("logger environment not supported")
	}

	zl, err := cfg.Build()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("zap logger build constructs failed. Err: %s", err))
	}

	// Disable stacktrace
	if config.RemoveStackTrace {
		zl = zl.WithOptions(zap.AddStacktrace(zap.DPanicLevel))
	}

	return zl.Sugar(), nil
}
