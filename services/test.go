package services

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

// SERVICE_NAMEServiceProvider
var SERVICE_NAMEServiceProvider = wire.NewSet(NewSERVICE_NAME, wire.Bind(new(ISERVICE_NAME), new(*SERVICE_NAME)))

// ISERVICE_NAME
type ISERVICE_NAME interface {
	Get(ctx context.Context, Id uint64) (string, error)
}

// SERVICE_NAME
type SERVICE_NAME struct {
	logger *zap.Logger
	// Introduce new dependencies here, exp:
	// userRepository *repositories.UserRepository
}

// NewSERVICE_NAME
func NewSERVICE_NAME(
	logger *zap.Logger,
	// userRepository *repositories.UserRepository,
) *SERVICE_NAME {
	return &SERVICE_NAME{
		logger: logger.With(zap.String("type", "SERVICE_NAME")),
		// userRepository: userRepository,
	}
}

// Get
func (p *SERVICE_NAME) Get(ctx context.Context, Id uint64) (string, error) {
	return "name", nil
}
