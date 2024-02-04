package biz

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	ctrl     *gomock.Controller
	storager *repo.MockStorager
	biz      biz.IIssueBiz
}

func (s *suiteTester) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())
	s.ctrl = gomock.NewController(s.T())
	s.storager = repo.NewMockStorager(s.ctrl)
	issueBiz, err := NewIssueBiz(s.storager)
	s.Require().NoError(err)
	s.biz = issueBiz
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_CreateTodo() {
	type args struct {
		ctx   contextx.Contextx
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantTodo *model.Ticket
		wantErr  bool
	}{
		{
			name: "create todo by title then error",
			args: args{title: "test", mock: func() {
				s.storager.EXPECT().Create(gomock.Any(), "test").Return(
					nil,
					errors.New("mock error"),
				).Times(1)
			}},
			wantTodo: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTodo, err := s.biz.CreateTodo(tt.args.ctx, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTodo, tt.wantTodo) {
				t.Errorf("CreateTodo() gotTodo = %v, want %v", gotTodo, tt.wantTodo)
			}
		})
	}
}
