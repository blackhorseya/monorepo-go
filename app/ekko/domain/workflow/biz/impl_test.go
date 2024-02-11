package biz

import (
	"errors"
	"reflect"
	"testing"

	idM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/identity/model"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/agg"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/biz"
	wfM "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/model"
	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	ctrl   *gomock.Controller
	issues *repo.MockIIssueRepo
	biz    biz.IWorkflowBiz
}

func (s *suiteTester) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())
	s.ctrl = gomock.NewController(s.T())
	s.issues = repo.NewMockIIssueRepo(s.ctrl)
	workflowBiz, err := NewWorkflowBiz(s.issues)
	s.Require().NoError(err)
	s.biz = workflowBiz
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_CreateTodo() {
	ticket1 := &wfM.Ticket{
		Title: "title1",
	}
	issue1, _ := agg.NewIssue(ticket1.Title)

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
			name: "create todo then error",
			args: args{title: ticket1.Title, mock: func() {
				s.issues.EXPECT().Create(gomock.Any(), issue1).
					Return("", errors.New("mock error")).
					Times(1)
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

			gotTodo, err := s.biz.CreateTodo(tt.args.ctx, tt.args.who, tt.args.title)
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
