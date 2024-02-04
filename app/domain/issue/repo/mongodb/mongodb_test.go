package mongodb

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteTester struct {
	suite.Suite

	container *mongodb.Container
	rw        *mongo.Client
	storage   repo.Storager
}

func (s *suiteTester) SetupTest() {
	ctx := contextx.Background()

	container, err := mongodb.NewContainer(ctx)
	s.Require().NoError(err)
	s.container = container

	dsn, err := s.container.ConnectionString(ctx)
	s.Require().NoError(err)

	rw, err := mongodb.NewClientWithDSN(dsn)
	s.Require().NoError(err)
	s.rw = rw

	storager, err := NewStorager(s.rw)
	s.Require().NoError(err)
	s.storage = storager
}

func (s *suiteTester) TearDownTest() {
	ctx := contextx.Background()

	err := s.rw.Disconnect(ctx)
	s.Require().NoError(err)

	err = s.container.Terminate(ctx)
	s.Require().NoError(err)
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_List() {
	type args struct {
		ctx  contextx.Contextx
		opts repo.ListOptions
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantTodos []*model.Ticket
		wantTotal int
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTodos, gotTotal, err := s.storage.List(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTodos, tt.wantTodos) {
				t.Errorf("List() gotTodos = %v, want %v", gotTodos, tt.wantTodos)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("List() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func (s *suiteTester) Test_impl_Create() {
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
			name:     "",
			args:     args{},
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

			gotTodo, err := s.storage.Create(tt.args.ctx, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTodo, tt.wantTodo) {
				t.Errorf("Create() gotTodo = %v, want %v", gotTodo, tt.wantTodo)
			}
		})
	}
}
