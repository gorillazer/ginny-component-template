package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gorillazer/ginny/res"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// HANDLE_NAMEHandlerProvider
var HANDLE_NAMEHandlerProvider = wire.NewSet(NewHANDLE_NAMEHandler, wire.Bind(new(IHANDLE_NAMEHandler), new(*HANDLE_NAMEHandler)))

// IHANDLE_NAMEHandler
type IHANDLE_NAMEHandler interface {
	Get(c *gin.Context) (*res.Response, error)
}

// HANDLE_NAMEHandler
type HANDLE_NAMEHandler struct {
	v      *viper.Viper
	logger *zap.Logger
	// testService  *services.TestService
}

// NewHANDLE_NAMEHandler
func NewHANDLE_NAMEHandler(
	v *viper.Viper,
	logger *zap.Logger,
	// testService *services.TestService,
) *HANDLE_NAMEHandler {
	return &HANDLE_NAMEHandler{
		v:      v,
		logger: logger.With(zap.String("type", "TestHandler")),
		// testService:  testService,
	}
}

func (t *HANDLE_NAMEHandler) Get(c *gin.Context) (*res.Response, error) {
	t.logger.Debug(zap.Any("HANDLE_NAMEHandler.Get", c.Params))
	// name, err := t.testService.GetInfo(c)
	// if err != nil {
	// 	t.logger.Error(zap.Error(err))
	// 	return nil, errs.New(configs.ERR_GETINFO, configs.GetErrMsg(configs.ERR_GETINFO))
	// }
	return res.Success("name"), nil
}
