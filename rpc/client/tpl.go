package client

import (
	// 调用第三方系统的sever请修改此地址
	proto "MODULE_NAME/api/proto"

	consul "github.com/gorillazer/ginny-consul"
	"github.com/gorillazer/ginny-serve/grpc"
	"github.com/pkg/errors"
)

// NewCLIENT_NAMEClient
func NewCLIENT_NAMEClient(
	client *grpc.Client,
	consul *consul.Client,
) (proto.CLIENT_NAMEClient, error) {
	// conn, err := client.Dial("CLIENT_NAME", grpc.WithTarget("127.0.0.1:9000"))
	conn, err := client.Dial("CLIENT_NAME", grpc.WithConsulConfig(consul.Config))
	if err != nil {
		return nil, errors.Wrap(err, "client dial error")
	}
	c := proto.NewCLIENT_NAMEClient(conn)
	return c, nil
}
