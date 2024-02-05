package mongodb_test

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/app/domain/market/repo/mongodb"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	dbName   = "orianna"
	collName = "stocks"
)

type suiteTester struct {
	suite.Suite

	container *mongodbx.Container
	rw        *mongo.Client
	storage   repo.Storager
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

	storager, err := mongodb.NewStorager(s.rw)
	s.Require().NoError(err)
	s.storage = storager
}

func (s *suiteTester) TearDownTest() {
	ctx := contextx.Background()

	if s.rw != nil {
		err := s.rw.Disconnect(ctx)
		s.NoError(err)
	}

	if s.container != nil {
		err := s.container.Terminate(ctx)
		s.NoError(err)
	}
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetBySymbol() {
	stock1 := &model.StockInfo{
		Symbol: "2330",
		Name:   "2330",
	}

	type args struct {
		ctx    contextx.Contextx
		symbol string
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *model.StockInfo
		wantErr  bool
	}{
		{
			name:     "not found",
			args:     args{symbol: "not_found"},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "ok",
			args: args{symbol: stock1.Symbol, mock: func() {
				_, _ = s.rw.Database(dbName).Collection(collName).InsertOne(contextx.Background(), stock1)
			}},
			wantInfo: stock1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.storage.GetBySymbol(tt.args.ctx, tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBySymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetBySymbol() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
