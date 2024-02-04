package biz_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/issue/biz"
	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo"
	issueB "github.com/blackhorseya/monorepo-go/entity/domain/issue/biz"
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
	biz      issueB.IIssueBiz
}

func (s *suiteTester) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())
	s.ctrl = gomock.NewController(s.T())
	s.storager = repo.NewMockStorager(s.ctrl)
	issueBiz, err := biz.NewIssueBiz(s.storager)
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

func (s *suiteTester) Test_impl_ListTodos() {
	type args struct {
		ctx  contextx.Contextx
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantTodos []*model.Ticket
		wantTotal int
		wantErr   bool
	}{
		{
			name: "list todos then error",
			args: args{mock: func() {
				s.storager.EXPECT().List(gomock.Any(), gomock.Any()).Return(
					nil,
					0,
					errors.New("mock error"),
				).Times(1)
			}},
			wantTodos: nil,
			wantTotal: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTodos, gotTotal, err := s.biz.ListTodos(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTodos, tt.wantTodos) {
				t.Errorf("ListTodos() gotTodos = %v, want %v", gotTodos, tt.wantTodos)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListTodos() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
