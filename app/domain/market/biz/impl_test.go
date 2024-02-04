package biz_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/market/biz"
	marketB "github.com/blackhorseya/monorepo-go/entity/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	ctrl    *gomock.Controller
	finmind *finmindx.MockDialer
	biz     marketB.IMarketBiz
}

func (s *suiteTester) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())
	s.ctrl = gomock.NewController(s.T())
	s.finmind = finmindx.NewMockDialer(s.ctrl)
	marketBiz, err := biz.NewMarketBiz(s.finmind)
	s.NoError(err)
	s.biz = marketBiz
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
