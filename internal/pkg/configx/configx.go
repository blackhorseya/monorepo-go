package configx

import (
	"github.com/spf13/viper"
)

// Config defines the config struct.
type Config struct {
	HTTP HTTP `json:"http" yaml:"http"`
	GRPC GRPC `json:"grpc" yaml:"grpc"`
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

// NewExample will create a new example config instance.
func NewExample() *Config {
	return &Config{
		HTTP: HTTP{
			Host: "0.0.0.0",
			Port: 1992,
			Mode: "debug",
		},
		GRPC: GRPC{
			Host: "0.0.0.0",
			Port: 11992,
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
