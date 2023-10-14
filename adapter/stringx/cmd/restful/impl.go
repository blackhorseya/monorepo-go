package restful

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/endpoints"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/stringx/transport"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type impl struct {
	viper  *viper.Viper
	logger *zap.Logger

	server *http.Server
	svc    biz.IStringBiz
}

func newImpl(viper *viper.Viper, logger *zap.Logger, svc biz.IStringBiz) adapterx.Servicer {
	return &impl{
		viper:  viper,
		logger: logger.With(zap.String("type", "restful")),
		svc:    svc,
	}
}

func (i *impl) Start() error {
	uppercaseHandler := transport.MakeUppercaseHandler(contextx.Background(), endpoints.MakeUppercaseEndpoint(i.svc))
	countHandler := transport.MakeCountHandler(contextx.Background(), endpoints.MakeCountEndpoint(i.svc))

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	i.server = &http.Server{
		Addr:                         "0.0.0.0:8080",
		Handler:                      nil,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}

	go func() {
		i.logger.Info("start restful service", zap.String("addr", i.server.Addr))

		err := i.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			i.logger.Fatal("restful service error", zap.Error(err))
		}
	}()

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		i.logger.Info("receive signal", zap.String("signal", sig.String()))

		timeout, cancelFunc := contextx.WithTimeout(contextx.Background(), 5*time.Second)
		defer cancelFunc()

		err := i.server.Shutdown(timeout)
		if err != nil {
			i.logger.Error("shutdown restful server error", zap.Error(err))
			return err
		}
	}

	return nil
}
