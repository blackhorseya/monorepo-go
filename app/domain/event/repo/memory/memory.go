package memory

import (
	eventR "github.com/blackhorseya/monorepo-go/app/domain/event/repo"
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	packets map[string]*eventM.RedPacket
}

// New create a new memory storage.
func New() eventR.Storager {
	return &impl{
		packets: make(map[string]*eventM.RedPacket),
	}
}

func (i *impl) CreateRedPacket(ctx contextx.Contextx, packet *eventM.RedPacket) error {
	i.packets[packet.Id] = packet

	return nil
}

func (i *impl) ListRedPacket(
	ctx contextx.Contextx,
	cond eventR.ListRedPacketCondition,
) (packets []*eventM.RedPacket, err error) {
	for _, packet := range i.packets {
		packets = append(packets, packet)
	}

	return packets, nil
}
