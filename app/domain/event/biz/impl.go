package biz

import (
	"errors"
	"time"

	eventR "github.com/blackhorseya/monorepo-go/app/domain/event/repo"
	eventB "github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type impl struct {
	storage eventR.Storager
}

// New create a new event eventB.
func New(storage eventR.Storager) eventB.IEventBiz {
	return &impl{
		storage: storage,
	}
}

func (i *impl) CreateRedPacket(
	ctx contextx.Contextx,
	who *userM.UserAccount,
	amount uint64,
	count uint32,
) (packet *eventM.RedPacket, err error) {
	if who == nil {
		ctx.Error("who is nil")
		return nil, errors.New("who is nil")
	}

	if amount == 0 {
		ctx.Error("amount is 0")
		return nil, errors.New("amount is 0")
	}

	if count == 0 {
		ctx.Error("count is 0")
		return nil, errors.New("count is 0")
	}

	totalAmount := amount * uint64(count)
	now := time.Now()
	ret := &eventM.RedPacket{
		Id:              uuid.New().String(),
		CreatorId:       who.Id,
		TotalAmount:     totalAmount,
		RemainingAmount: totalAmount,
		TotalCount:      count,
		RemainingCount:  count,
		CreatedAt:       now.UTC().Format(time.RFC3339),
		UpdatedAt:       now.UTC().Format(time.RFC3339),
	}

	err = i.storage.CreateRedPacket(ctx, ret)
	if err != nil {
		ctx.Error("create red packet failed", zap.Error(err))
		return nil, err
	}

	return ret, nil
}

func (i *impl) ListRedPacket(
	ctx contextx.Contextx,
	cond eventB.ListRedPacketCondition,
) (list []*eventM.RedPacket, err error) {
	ret, err := i.storage.ListRedPacket(ctx, eventR.ListRedPacketCondition{})
	if err != nil {
		ctx.Error("list red packet failed", zap.Error(err))
		return nil, err
	}

	return ret, nil
}

func (i *impl) GrabRedPacket(
	ctx contextx.Contextx,
	who *userM.UserAccount,
	packetID string,
) (record *eventM.GrabRecord, err error) {
	// todo: 2023/10/28|sean|impl me
	panic("implement me")
}
