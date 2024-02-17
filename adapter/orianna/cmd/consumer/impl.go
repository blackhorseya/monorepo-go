package consumer

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/transports/kafkax"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var (
	partitionCount = 6

	topicName = "daily_quote"
)

type impl struct {
	client *lambda.Lambda
	done   chan struct{}
}

func newConsumer() (adapterx.Servicer, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	client := lambda.New(sess)

	return &impl{
		client: client,
		done:   make(chan struct{}),
	}, nil
}

func (i *impl) Start() error {
	for id := 0; id < partitionCount; id++ {
		go i.execute(id)
	}

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		ctx := contextx.Background()
		ctx.Info("receive signal", zap.String("signal", sig.String()))

		ctx.Info("consumer is stopped")
	}

	return nil
}

func (i *impl) execute(id int) {
	ctx := contextx.Background()

	reader, err := kafkax.NewReaderWithTopic(topicName)
	if err != nil {
		ctx.Error("new reader error", zap.Error(err))
		return
	}

	var message kafka.Message
	for {
		message, err = reader.ReadMessage(ctx)
		if err != nil {
			ctx.Error("read message error", zap.Error(err))
			continue
		}

		var result *lambda.InvokeOutput
		result, err = i.client.Invoke(&lambda.InvokeInput{
			FunctionName: aws.String("prod-calcLongUp"),
			LogType:      aws.String("Tail"),
			Payload:      message.Value,
		})
		if err != nil {
			ctx.Error("invoke lambda error", zap.Error(err))
			continue
		}

		ctx.Info("invoke lambda success", zap.Int("id", id), zap.Any("result", &result))
	}
}