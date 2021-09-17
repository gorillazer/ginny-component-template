package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	// "github.com/gorillazer/ginny/errs"
	"github.com/gorillazer/ginny/res"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// HANDLE_NAMEHandlerProvider
var HANDLE_NAMEHandlerProvider = wire.NewSet(NewHANDLE_NAMEHandler, wire.Bind(new(iHANDLE_NAMEHandler), new(*HANDLE_NAMEHandler)))

// iHANDLE_NAMEHandler
type iHANDLE_NAMEHandler interface {
	Get(c *gin.Context) (*res.Response, error)
}

// HANDLE_NAMEHandler
type HANDLE_NAMEHandler struct {
	v      *viper.Viper
	logger *zap.Logger
	// Introduce new dependencies here, exp:
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
		logger: logger.With(zap.String("type", "HANDLE_NAMEHandler")),
		// testService:  testService,
	}
}

func (t *HANDLE_NAMEHandler) Get(c *gin.Context) (*res.Response, error) {
	t.logger.Debug("HANDLE_NAMEHandler", zap.Any("HANDLE_NAMEHandler.Get", c.Params))
	// name, err := t.testService.GetInfo(c)
	// if err != nil {
	// 	t.logger.Error("Get", zap.Error(err))
	// 	return nil, errs.New(constants.NOT_FOUND, constants.GetErrMsg(constants.NOT_FOUND))
	// }
	return res.Success("name"), nil
}
