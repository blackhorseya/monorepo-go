package logx

import (
	"os"

	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewExample will create a new example logger instance.
func NewExample() *zap.Logger {
	return zap.NewExample()
}

// NewLoggerWithConfig will create a new logger instance with config.
func NewLoggerWithConfig(config *configx.Config) (logger *zap.Logger, err error) {
	level := zap.NewAtomicLevel()
	err = level.UnmarshalText([]byte(config.Log.Level))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)
	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(zapConfig)
	if config.Log.Format == "json" {
		zapConfig = zap.NewProductionEncoderConfig()
		enc = zapcore.NewJSONEncoder(zapConfig)
	}

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(enc, cw, level))

	core := zapcore.NewTee(cores...)
	ret := zap.New(core)

	zap.ReplaceGlobals(ret)

	return ret, nil
}
