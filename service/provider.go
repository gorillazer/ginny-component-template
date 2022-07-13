package service

import (
	"context"

	pb "MODULE_NAME/api/proto"
	"MODULE_NAME/internal/config"

	"github.com/google/wire"
	"github.com/goriller/ginny"
	"github.com/goriller/ginny/server/mux"
)

// ProviderSet
var ProviderSet = wire.NewSet(NewService, RegisterService)

// Service the instance for grpc proto.
type Service struct {
	pb.UnimplementedSERVICE_NAMEServer
	config *config.Config
	// Introduce new dependencies here, exp:
	// userRepository *repo.UserRepo
}

// NewService new service that implement hello
func NewService(
	config *config.Config,
) *Service {
	mux.RegisterErrorCodes(pb.ErrorCode_name)
	return &Service{
		config: config,
	}
}

// RegisterService
func RegisterService(ctx context.Context, sev *Service) ginny.RegistrarFunc {
	return func(app *ginny.Application) error {
		// 注册gRPC服务
		app.Server.RegisterService(&pb.SERVICE_NAME_ServiceDesc, sev)

		if app.Option.HttpAddr != "" {
			// 注册http服务
			if err := pb.RegisterSERVICE_NAMEHandlerServer(ctx, app.Server.ServeMux(), sev); err != nil {
				return err
			}
		}
		return nil
	}
}
