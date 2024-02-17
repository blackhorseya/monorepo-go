package kafkax

import (
	"crypto/tls"

	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// NewReader returns a new Reader instance.
func NewReader() (*kafka.Reader, error) {
	id := uuid.New().String()
	if configx.A.MessageQueue.Kafka.GroupID != "" {
		id = configx.A.MessageQueue.Kafka.GroupID
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: configx.A.MessageQueue.Kafka.Brokers,
		GroupID: id,
		Topic:   configx.A.MessageQueue.Kafka.Topic,
		Dialer: &kafka.Dialer{
			DualStack: true,
			TLS: &tls.Config{
				InsecureSkipVerify: true, //nolint:gosec // skip verify
			},
			SASLMechanism: plain.Mechanism{
				Username: configx.A.MessageQueue.Kafka.Username,
				Password: configx.A.MessageQueue.Kafka.Password,
			},
		},
	}), nil
}
