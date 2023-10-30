package memory

import (
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/storage"
)

type impl struct {
	packets map[string]*eventM.RedPacket
}

// New create a new memory storage.
func New() storage.Storager {
	// todo: 2023/10/30|sean|impl me
	panic("implement me")
}
