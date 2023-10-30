//go:build external

package biz_test

import (
	"testing"

	eventB "github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo/memory"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	logger  *zap.Logger
	storage repo.Storager
	biz     eventB.IEventBiz
}

func (s *suiteExternal) SetupTest() {
	s.logger = zap.NewExample()
	s.storage = memory.New()
	s.biz = biz.New(s.storage)
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) TestCreateRedPacket() {
	t := s.T()

	t.Skip()
}
