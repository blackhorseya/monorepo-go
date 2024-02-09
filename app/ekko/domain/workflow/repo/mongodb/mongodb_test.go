package mongodb_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/app/ekko/domain/workflow/repo/mongodb"
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

	repo, err := mongodb.NewIssueRepoWithMongoDB(s.rw)
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
