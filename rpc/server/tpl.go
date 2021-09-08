package rpc_server

import (
	"MODULE_NAME/api/proto"
	"MODULE_NAME/internal/constants"
	"context"

	"github.com/google/wire"
	"github.com/gorillazer/ginny/errs"
	"go.uber.org/zap"
)

// SERVER_NAMEServerProvider
var SERVER_NAMEServerProvider = wire.NewSet(NewSERVER_NAMEServer, wire.Bind(new(ISERVER_NAMEServer), new(*SERVER_NAMEServer)))

// ISERVER_NAMEServer
type ISERVER_NAMEServer interface {
	Get(ctx context.Context, req *proto.SERVER_NAMEReq) (*proto.SERVER_NAMERes, error)
}

// SERVER_NAMEServer
type SERVER_NAMEServer struct {
	logger *zap.Logger
	// Introduce new dependencies here, exp:
	// testService  *services.TestService
}

// NewSERVER_NAMEServer
func NewSERVER_NAMEServer(
	logger *zap.Logger,
	// testService *services.TestService,
) (*SERVER_NAMEServer, error) {
	return &SERVER_NAMEServer{
		logger: logger.With(zap.String("type", "SERVER_SERVER_NAMEServer")),
		// testService:  testService,
	}, nil
}

func (s *SERVER_NAMEServer) Get(ctx context.Context, req *proto.SERVER_NAMEReq) (*proto.SERVER_NAMERes, error) {
	if req == nil {
		return nil, errs.New(constants.ERR_GETINFO, constants.GetErrMsg(constants.ERR_GETINFO))
	}
	// p, err := s.testService.Get(ctx, req.Id)
	// if err != nil {
	// 	return nil, errs.New(constants.ERR_GETINFO, constants.GetErrMsg(constants.ERR_GETINFO))
	// }

	resp := &proto.Detail{}

	return resp, nil
}
