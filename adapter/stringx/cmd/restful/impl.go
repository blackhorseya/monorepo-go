package restful

import (
	"net/http"

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

	svc biz.IStringBiz
}

func newImpl(viper *viper.Viper, logger *zap.Logger, svc biz.IStringBiz) adapterx.Servicer {
	return &impl{
		viper:  viper,
		logger: logger.With(zap.String("type", "restful")),
		svc:    svc,
	}
}

func (i *impl) Start() error {
	i.logger.Info("start restful service")

	// todo: 2023/10/12|sean|impl me

	uppercaseHandler := transport.MakeUppercaseHandler(contextx.Background(), endpoints.MakeUppercaseEndpoint(i.svc))
	countHandler := transport.MakeCountHandler(contextx.Background(), endpoints.MakeCountEndpoint(i.svc))

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	server := &http.Server{
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

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) AwaitSignal() error {
	i.logger.Info("await restful service signal")

	// todo: 2023/10/12|sean|impl me

	return nil
}
