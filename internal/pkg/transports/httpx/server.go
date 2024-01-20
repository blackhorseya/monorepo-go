package httpx

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// Server is a http server.
type Server struct {
	httpserver *http.Server
	Router     *gin.Engine
}

// NewServer is used to create a new http server.
func NewServer(ctx contextx.Contextx) (*Server, error) {
	router := gin.New()

	// register middleware
	router.Use(ginzap.GinzapWithConfig(ctx, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/healthz"},
		Context:    nil,
	}))
	router.Use(ginzap.CustomRecoveryWithZap(ctx, true, func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Err.WithMessage(fmt.Sprintf("%v", err)))
	}))
	router.Use(contextx.AddContextxMiddleware())
	router.Use(response.AddErrorHandlingMiddleware())

	// init http server
	httpserver := &http.Server{
		Addr:    configx.C.HTTP.GetAddr(),
		Handler: router,
	}

	return &Server{
		httpserver: httpserver,
		Router:     router,
	}, nil
}
