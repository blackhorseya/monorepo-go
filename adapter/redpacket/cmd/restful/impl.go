package restful

import (
	"net/http"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	config *configx.Config
	logger *zap.Logger

	router *gin.Engine
	server *http.Server
}

func newImpl(config *configx.Config, logger *zap.Logger) adapterx.Servicer {
	return &impl{
		config: config,
		logger: logger.With(zap.String("type", "restful")),
		router: gin.New(),
		server: nil,
	}
}

func (i *impl) Start() error {
	// todo: 2023/11/3|sean|implement me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// todo: 2023/11/3|sean|implement me
	panic("implement me")
}
