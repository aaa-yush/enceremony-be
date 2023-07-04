package repo

import (
	"context"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
)

type impl struct {
	logger *logger.Logger
	conf   *config.Config
}

type ProductRepo interface {
	GetProductsByEventId(ctx context.Context, eventId string)
}
