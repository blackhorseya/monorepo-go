package biz

import (
	"errors"
	"time"

	"github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/google/uuid"
)

type impl struct {
	storage repo.Storager
}

// New create a new event biz.
func New(storage repo.Storager) biz.IEventBiz {
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
		return nil, errors.New("who is nil")
	}

	if amount == 0 {
		return nil, errors.New("amount is 0")
	}

	if count == 0 {
		return nil, errors.New("count is 0")
	}

	now := time.Now()
	ret := &eventM.RedPacket{
		Id:              uuid.New().String(),
		CreatorId:       who.Id,
		TotalAmount:     amount,
		RemainingAmount: amount,
		TotalCount:      count,
		RemainingCount:  count,
		CreatedAt:       now.UTC().Format(time.RFC3339),
		UpdatedAt:       now.UTC().Format(time.RFC3339),
	}

	err = i.storage.CreateRedPacket(ctx, ret)
	if err != nil {
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
