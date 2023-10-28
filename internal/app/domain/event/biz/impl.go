package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
}

// New create a new event biz.
func New() biz.IEventBiz {
	return &impl{}
}

func (i *impl) CreateRedPacket(
	ctx contextx.Contextx,
	who *userM.UserAccount,
	amount uint64,
	count uint32,
) (packet *eventM.RedPacket, err error) {
	// todo: 2023/10/28|sean|impl me
	panic("implement me")
}

func (i *impl) GrabRedPacket(
	ctx contextx.Contextx,
	who *userM.UserAccount,
	packetID string,
) (record *eventM.GrabRecord, err error) {
	// todo: 2023/10/28|sean|impl me
	panic("implement me")
}
