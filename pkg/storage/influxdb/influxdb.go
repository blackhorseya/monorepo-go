package influxdb

import (
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
)

// NewClient is used to create a new influxdb client.
func NewClient() (*influxdb3.Client, error) {
	return influxdb3.New(influxdb3.ClientConfig{
		Host:         configx.A.Storage.Influxdb.URL,
		Token:        configx.A.Storage.Influxdb.Token,
		Organization: "",
		Database:     "",
		HTTPClient:   nil,
		WriteOptions: nil,
		Headers:      nil,
	})
}
