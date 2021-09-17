package services

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

// SERVICE_NAMEServiceProvider
var SERVICE_NAMEServiceProvider = wire.NewSet(NewSERVICE_NAMEService, wire.Bind(new(iSERVICE_NAMEService), new(*SERVICE_NAMEService)))

// iSERVICE_NAMEService
type iSERVICE_NAMEService interface {
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
		logger: logger.With(zap.String("type", "SERVICE_NAMEService")),
		// userRepository: userRepository,
	}
}

//
func (p *SERVICE_NAMEService) Get(ctx context.Context, Id uint64) (string, error) {
	// user, err := p.userRepository.GetUser(ctx)
	// if err != nil {
	// 	p.logger.Error("Get", zap.Error(err))
	// 	return "", errs.New(constants.NOT_FOUND, constants.GetErrMsg(constants.NOT_FOUND))
	// }
	return "name", nil
}
