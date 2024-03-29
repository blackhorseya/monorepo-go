package httpx

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server is a http server.
type Server struct {
	httpserver *http.Server
	Router     *gin.Engine
}

// NewServer is used to create a new http server.
func NewServer() (*Server, error) {
	return NewServerWithContextx(contextx.Background())
}

// NewServerWithContextx is used to create a new http server.
func NewServerWithContextx(ctx contextx.Contextx) (*Server, error) {
	return NewServerWithAPP(ctx, configx.A)
}

// NewServerWithAPP is used to create a new http server with app.
func NewServerWithAPP(ctx contextx.Contextx, app *configx.Application) (*Server, error) {
	gin.SetMode(app.HTTP.Mode)

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
		Addr:              app.HTTP.GetAddr(),
		Handler:           router,
		ReadHeaderTimeout: time.Second,
	}

	return &Server{
		httpserver: httpserver,
		Router:     router,
	}, nil
}

// Start starts the http server.
func (s *Server) Start(ctx contextx.Contextx) error {
	ctx.Info("start listen and serve", zap.String("addr", s.httpserver.Addr))

	go func() {
		err := s.httpserver.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			ctx.Fatal("start http server error", zap.Error(err))
		}
	}()

	return nil
}

// Stop stops the http server.
func (s *Server) Stop(ctx contextx.Contextx) error {
	ctx.Info("shutdown http server")

	timeout, cancelFunc := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	err := s.httpserver.Shutdown(timeout)
	if err != nil {
		ctx.Error("shutdown http server error", zap.Error(err))
		return err
	}

	return nil
}
