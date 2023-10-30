package biz_test

import (
	"testing"

	eventB "github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/storage"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type tester struct {
	suite.Suite

	logger  *zap.Logger
	ctrl    *gomock.Controller
	storage *storage.MockStorager
	biz     eventB.IEventBiz
}

func (s *tester) SetupTest() {
	s.logger = zap.NewExample()
	s.ctrl = gomock.NewController(s.T())
	s.storage = storage.NewMockStorager(s.ctrl)
	s.biz = biz.New()
}

func (s *tester) TearDownSuite() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(tester))
}
