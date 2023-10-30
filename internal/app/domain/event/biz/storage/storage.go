//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package storage

import (
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// ListRedPacketCondition defines the condition for list red packet.
type ListRedPacketCondition struct {
}

// Storager defines the storage interface.
type Storager interface {
	// CreateRedPacket create a new red packet.
	CreateRedPacket(ctx contextx.Contextx, packet *eventM.RedPacket) error

	// ListRedPacket list red packet by user account.
	ListRedPacket(ctx contextx.Contextx, cond ListRedPacketCondition) (packets []*eventM.RedPacket, err error)
}
