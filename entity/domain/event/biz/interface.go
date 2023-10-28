//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IEventBiz event biz interface
type IEventBiz interface {
	// CreateRedPacket serve caller to create a red packet
	CreateRedPacket(
		ctx contextx.Contextx,
		who *userM.UserAccount,
		amount uint64,
		count uint32,
	) (packet *eventM.RedPacket, err error)
}
