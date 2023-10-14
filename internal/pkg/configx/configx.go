package configx

import (
	"github.com/blackhorseya/monorepo-go/pkg/netx"
	"github.com/spf13/viper"
)

// Config defines the config struct.
type Config struct {
	HTTP HTTP `json:"http" yaml:"http"`
	GRPC GRPC `json:"grpc" yaml:"grpc"`
	Log  Log  `json:"log" yaml:"log"`
}

// HTTP defines the http config struct.
type HTTP struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// GRPC defines the grpc config struct.
type GRPC struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

// Log defines the log config struct.
type Log struct {
	Level  string `json:"level" yaml:"level"`
	Format string `json:"format" yaml:"format"`
}

// NewExample will create a new example config instance.
func NewExample() *Config {
	return &Config{
		HTTP: HTTP{
			Host: "",
			Port: netx.GetAvailablePort(),
			Mode: "release",
		},
		GRPC: GRPC{
			Host: "",
			Port: netx.GetAvailablePort(),
		},
		Log: Log{
			Level:  "info",
			Format: "json",
		},
	}
}

// NewWithViper will create a new config instance with viper.
func NewWithViper(v *viper.Viper) (*Config, error) {
	cfg := NewExample()
	err := v.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
