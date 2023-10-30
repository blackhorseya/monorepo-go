package memory

import (
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

type impl struct {
	packets map[string]*eventM.RedPacket
}

// New create a new memory storage.
func New() repo.Storager {
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
	cond repo.ListRedPacketCondition,
) (packets []*eventM.RedPacket, err error) {
	for _, packet := range i.packets {
		packets = append(packets, packet)
	}

	return packets, nil
}
