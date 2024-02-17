package kafkax

import (
	"crypto/tls"

	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// NewWriter returns a new Writer instance.
func NewWriter() (*kafka.Writer, error) {
	return &kafka.Writer{
		Addr:     kafka.TCP(configx.A.MessageQueue.Kafka.Brokers...),
		Balancer: &kafka.Hash{},
		Transport: &kafka.Transport{
			TLS: &tls.Config{
				InsecureSkipVerify: true, //nolint:gosec // it's okay to skip verification for internal services
			},
			SASL: plain.Mechanism{
				Username: configx.A.MessageQueue.Kafka.Username,
				Password: configx.A.MessageQueue.Kafka.Password,
			},
		},
	}, nil
}
