//go:build external

package biz

import (
	"testing"

	idM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/identity/model"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/biz"
	wfM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/model"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/repo"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	issues repo.IIssueRepo
	biz    biz.IWorkflowBiz
}

func (s *suiteExternal) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	s.NoError(err)

	configx.ReplaceApplication(configx.C.Ekko)

	workflowBiz, err := NewForExternal()
	s.NoError(err)
	s.biz = workflowBiz
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) Test_impl_CreateTodo() {
	type args struct {
		ctx   contextx.Contextx
		who   *idM.User
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantTodo *wfM.Ticket
		wantErr  bool
	}{
		{
			name:     "ok",
			args:     args{title: "title1"},
			wantTodo: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTodo, err := s.biz.CreateTodo(tt.args.ctx, tt.args.who, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.args.ctx.Debug("gotTodo", zap.Any("gotTodo", &gotTodo))
		})
	}
}

func (s *suiteExternal) Test_impl_ListTodos() {
	type args struct {
		ctx  contextx.Contextx
		opts biz.ListTodosOptions
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantTodos []*wfM.Ticket
		wantTotal int
		wantErr   bool
	}{
		{
			name:      "ok",
			args:      args{},
			wantTodos: nil,
			wantTotal: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTodos, gotTotal, err := s.biz.ListTodos(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.args.ctx.Debug("gotTodos", zap.Any("gotTodos", gotTodos), zap.Int("gotTotal", gotTotal))
		})
	}
}
