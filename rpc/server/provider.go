package rpc_server

import (
	// SERVER_IMPORT 锚点请勿删除! Do not delete this line!

	"github.com/google/wire"
	"github.com/gorillazer/ginny-serve/grpc"
	stdGrpc "google.golang.org/grpc"
)

// CreateInitServerFn
func CreateInitServerFn(
// SERVER_PARAM 锚点请勿删除! Do not delete this line!
) grpc.InitServers {
	return func(s *stdGrpc.Server) {
		// SERVER_REGIST 锚点请勿删除! Do not delete this line!
	}
}

// ProviderSet
var ProviderSet = wire.NewSet(
	// SERVER_PROVIDER 锚点请勿删除! Do not delete this line!
	CreateInitServerFn,
)
