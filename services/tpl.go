package services

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

// SERVICE_NAMEServiceProvider
var SERVICE_NAMEServiceProvider = wire.NewSet(NewSERVICE_NAMEService, wire.Bind(new(ISERVICE_NAMEService), new(*SERVICE_NAMEService)))

// ISERVICE_NAMEService
type ISERVICE_NAMEService interface {
	Get(ctx context.Context, Id uint64) (string, error)
}

// SERVICE_NAMEService
type SERVICE_NAMEService struct {
	logger *zap.Logger
	// Introduce new dependencies here, exp:
	// userRepository *repositories.UserRepository
}

// NewSERVICE_NAMEService
func NewSERVICE_NAMEService(
	logger *zap.Logger,
	// userRepository *repositories.UserRepository,
) *SERVICE_NAMEService {
	return &SERVICE_NAMEService{
		logger: logger.With(zap.String("type", "Hello")),
		// userRepository: userRepository,
	}
}

//
func (p *SERVICE_NAMEService) Get(ctx context.Context, Id uint64) (string, error) {
	return "name", nil
}
