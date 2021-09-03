package repositories

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	// DATABASE_LIB 命令行工具锚点,请勿删除本行注释! Important! Do not delete this line
)

var REPO_NAMERepositoryProvider = wire.NewSet(NewREPO_NAMERepository, wire.Bind(new(IREPO_NAMERepository), new(*REPO_NAMERepository)))

type IREPO_NAMERepository interface {
	GetUser(ctx context.Context) (*REPO_NAMERepository, error)
}
type REPO_NAMERepository struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`

	logger *zap.Logger
	// STRUCT_ATTR 命令行工具锚点,请勿删除本行注释! Important! Do not delete this line
}

func NewREPO_NAMERepository(
	logger *zap.Logger,
	// FUNC_PARAM 命令行工具锚点,请勿删除本行注释! Important! Do not delete this line
) *REPO_NAMERepository {
	return &REPO_NAMERepository{
		logger: logger.With(zap.String("type", "REPO_NAMERepository")),
		// FUNC_ATTR 命令行工具锚点,请勿删除本行注释! Important! Do not delete this line
	}
}

func (p *REPO_NAMERepository) GetUser(ctx context.Context) (*REPO_NAMERepository, error) {
	user := &REPO_NAMERepository{}
	// err := p.mysql.Find(ctx, user, "user", nil)
	// if err != nil {
	// 	p.logger.Error("REPO_NAMERepository.GetUser", zap.Error(err))
	// 	return nil, err
	// }
	// p.redis.DB().Get(ctx, "aa").Result()
	return user, nil
}
