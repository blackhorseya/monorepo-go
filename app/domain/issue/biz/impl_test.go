package biz

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/biz"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteTester struct {
	suite.Suite

	ctrl     *gomock.Controller
	storager *repo.MockStorager
	biz      biz.IIssueBiz
}

func (s *suiteTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.storager = repo.NewMockStorager(s.ctrl)
	issueBiz, err := NewIssueBiz()
	s.Require().NoError(err)
	s.biz = issueBiz
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
