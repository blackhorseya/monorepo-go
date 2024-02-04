package linebot

import (
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

// NewClient is used to create a new line bot client.
func NewClient() (*linebot.Client, error) {
	return linebot.New(configx.A.LineBot.Secret, configx.A.LineBot.Token)
}
