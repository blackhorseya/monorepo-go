//go:build external

package biz_test

import (
	"testing"

	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo/memory"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func TestImpl_CreateRedPacket(t *testing.T) {
	logger := zap.NewExample()
	storage := memory.New()
	event := biz.New(storage)

	ctx := contextx.WithLogger(logger)

	tester := &userM.UserAccount{
		Id:          uuid.New().String(),
		Username:    "tester",
		Email:       "tester@gmail.com",
		Password:    "",
		AccessToken: "",
		IdToken:     "",
		Profile:     nil,
	}

	packet, err := event.CreateRedPacket(ctx, tester, 100, 10)
	if err != nil {
		t.Error(err)
		return
	}

	ctx.Debug("got packet", zap.Any("packet", &packet))

	if packet.Id == "" {
		t.Error("packet id should not be empty")
		return
	}
}
