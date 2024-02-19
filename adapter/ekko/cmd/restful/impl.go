package restful

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/blackhorseya/monorepo-go/adapter/ekko/api/docs" // swagger docs
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/biz"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/model"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/response"
	"github.com/blackhorseya/monorepo-go/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type impl struct {
	server *httpx.Server
	bot    *linebot.Client
	svc    biz.IWorkflowBiz
}

func newRestful(bot *linebot.Client, svc biz.IWorkflowBiz) (adapterx.Servicer, error) {
	ctx := contextx.Background()

	server, err := httpx.NewServer(ctx)
	if err != nil {
		return nil, err
	}

	return &impl{
		server: server,
		bot:    bot,
		svc:    svc,
	}, nil
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	// register router
	api := i.server.Router.Group("/api")
	{
		api.GET("/healthz", i.healthz)
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		api.GET("/v1/todos", i.ListTodos)
		api.POST("/v1/todos", i.CreateTodo)

		api.POST("/callback", i.callback)
	}

	err := i.server.Start(ctx)
	if err != nil {
		return err
	}

	ctx.Info(
		"swagger docs",
		zap.String("url", fmt.Sprintf(
			"http://%s/api/docs/index.html",
			strings.ReplaceAll(configx.A.HTTP.GetAddr(), "0.0.0.0", "localhost"),
		)),
	)

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		ctx := contextx.Background()
		ctx.Info("receive signal", zap.String("signal", sig.String()))

		err := i.server.Stop(ctx)
		if err != nil {
			ctx.Error("shutdown restful server error", zap.Error(err))
			return err
		}
	}

	return nil
}

// healthz is used to check the health of the service.
// @Summary healthz
// @Description Check the health of the service.
// @Tags healthz
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /healthz [get]
func (i *impl) healthz(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}

// ListTodos is used to list all todos.
// @Summary List todos
// @Description List all todos
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/todos [get]
func (i *impl) ListTodos(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ret, total, err := i.svc.ListTodos(ctx, biz.ListTodosOptions{})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Header("X-Total-Count", fmt.Sprintf("%d", total))
	c.JSON(http.StatusOK, response.OK.WithData(ret))
}

type createTodoPayload struct {
	Title string `json:"title"`
}

// CreateTodo is used to create a todo.
// @Summary Create a todo
// @Description Create a todo
// @Tags todos
// @Accept json
// @Produce json
// @Param payload body createTodoPayload true "payload"
// @Success 201 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/todos [post]
func (i *impl) CreateTodo(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload createTodoPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ret, err := i.svc.CreateTodo(ctx, nil, payload.Title)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response.OK.WithData(ret))
}

// callback is used to handle the callback from the third-party service.
// @Summary callback
// @Description Handle the callback from the third-party service.
// @Tags callback
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /callback [post]
func (i *impl) callback(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	events, err := i.bot.ParseRequest(c.Request)
	if err != nil {
		if errors.Is(err, linebot.ErrInvalidSignature) {
			ctx.Error("invalid line bot signature", zap.Error(err))
			_ = c.Error(err)
		} else {
			ctx.Error("parse line bot request error", zap.Error(err))
			_ = c.Error(err)
		}

		return
	}

	var messages []linebot.SendingMessage
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			message, ok := event.Message.(*linebot.TextMessage)
			if !ok {
				continue
			}

			messages, err = i.handleMessage(ctx, message)
			if err != nil {
				ctx.Warn("handle message error", zap.Error(err), zap.String("text", message.Text))
				continue
			}

			_, err = i.bot.ReplyMessage(event.ReplyToken, messages...).Do()
			if err != nil {
				_ = c.Error(err)
				return
			}
		}
	}

	c.JSON(http.StatusOK, response.OK)
}

func (i *impl) handleMessage(ctx contextx.Contextx, message *linebot.TextMessage) ([]linebot.SendingMessage, error) {
	text := message.Text
	if text == "ping" {
		return []linebot.SendingMessage{
			linebot.NewTextMessage("pong"),
		}, nil
	}

	if text == "list" {
		var todos model.Tickets
		var err error
		todos, _, err = i.svc.ListTodos(ctx, biz.ListTodosOptions{
			Page: 1,
			Size: 5,
		})
		if err != nil {
			return nil, err
		}

		return []linebot.SendingMessage{
			todos.FlexMessage(),
		}, nil
	}

	if strings.HasPrefix(text, "create.") {
		title := strings.TrimPrefix(text, "create.")
		title = strings.TrimSpace(title)
		if title == "" {
			return nil, errors.New("title is required")
		}

		ticket, err := i.svc.CreateTodo(ctx, nil, title)
		if err != nil {
			return nil, err
		}

		return []linebot.SendingMessage{
			ticket.FlexMessage(),
		}, nil
	}

	return nil, errors.New("unknown message")
}
