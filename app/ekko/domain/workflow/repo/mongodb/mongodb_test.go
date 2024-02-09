package mongodb

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/agg"
	issueR "github.com/blackhorseya/monorepo-go/entity/ekko/domain/workflow/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	container *mongodbx.Container
	rw        *mongo.Client
	repo      issueR.IIssueRepo
}

func (s *suiteTester) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())
	ctx := contextx.Background()

	container, err := mongodbx.NewContainer(ctx)
	s.Require().NoError(err)
	s.container = container

	dsn, err := s.container.ConnectionString(ctx)
	s.Require().NoError(err)

	rw, err := mongodbx.NewClientWithDSN(dsn)
	s.Require().NoError(err)
	s.rw = rw

	repo, err := NewIssueRepoWithMongoDB(s.rw)
	s.Require().NoError(err)
	s.repo = repo
}

func (s *suiteTester) TearDownTest() {
	ctx := contextx.Background()

	if s.rw != nil {
		_ = s.rw.Disconnect(ctx)
	}

	if s.container != nil {
		_ = s.container.Terminate(ctx)
	}
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetByID() {
	issue1, _ := agg.NewIssue("title1")
	newIssue1 := newFromIssue(issue1)
	got1, _ := newIssue1.ToAggregate()

	type args struct {
		ctx  contextx.Contextx
		id   string
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantIssue agg.Issue
		wantErr   bool
	}{
		{
			name: "ok",
			args: args{id: newIssue1.ID.Hex(), mock: func() {
				_, err := s.rw.Database(dbName).Collection(collName).InsertOne(contextx.Background(), newIssue1)
				s.NoError(err)
			}},
			wantIssue: got1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotIssue, err := s.repo.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIssue, tt.wantIssue) {
				t.Errorf("GetByID() gotIssue = %v, want %v", gotIssue, tt.wantIssue)
			}
		})
	}
}

func (s *suiteTester) Test_impl_Create() {
	created1, _ := agg.NewIssue("title1")

	type args struct {
		ctx  contextx.Contextx
		item agg.Issue
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{item: created1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Create(tt.args.ctx, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
