package configx

import (
	"encoding/json"
	"fmt"

	"github.com/blackhorseya/monorepo-go/pkg/logging"
	"github.com/blackhorseya/monorepo-go/pkg/netx"
)

// Application defines the application struct.
type Application struct {
	Name         string `json:"name" yaml:"name"`
	ClientID     string `json:"client_id" yaml:"clientID"`
	ClientSecret string `json:"client_secret" yaml:"clientSecret"`

	Log     logging.Config `json:"log" yaml:"log"`
	HTTP    HTTP           `json:"http" yaml:"http"`
	Storage struct {
		Redis struct {
			Addr     string `json:"addr" yaml:"addr"`
			Password string `json:"password" yaml:"password"`
			DB       int    `json:"db" yaml:"db"`
		} `json:"redis" yaml:"redis"`
	} `json:"storage" yaml:"storage"`
}

func (a *Application) String() string {
	msg, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(msg)
}

// Config defines the config struct.
type Config struct {
	Log     logging.Config `json:"log" yaml:"log"`
	HTTP    HTTP           `json:"http" yaml:"http"`
	GRPC    GRPC           `json:"grpc" yaml:"grpc"`
	Cronjob Cronjob        `json:"cronjob" yaml:"cronjob"`
	Storage Storage        `json:"storage" yaml:"storage"`

	ShortenURL Application `json:"shorten_url" yaml:"shortenURL"`
}

func (c *Config) String() string {
	msg, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(msg)
}

// HTTP defines the http config struct.
type HTTP struct {
	URL  string `json:"url" yaml:"url"`
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// GetAddr will get the http address.
func (http *HTTP) GetAddr() string {
	if http.Host == "" {
		http.Host = "0.0.0.0"
	}

	if http.Port == 0 {
		http.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", http.Host, http.Port)
}

// GRPC defines the grpc config struct.
type GRPC struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

// GetAddr will get the grpc address.
func (g *GRPC) GetAddr() string {
	if g.Host == "" {
		g.Host = "0.0.0.0"
	}

	if g.Port == 0 {
		g.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}

// Cronjob defines the cronjob config struct.
type Cronjob struct {
	Interval int `json:"interval" yaml:"interval"`
}

// Storage defines the storage config struct.
type Storage struct {
	DSN   string `json:"dsn" yaml:"dsn"`
	Conns int    `json:"conns" yaml:"conns"`
}
