package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/agg"
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/biz"
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/model"
	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	ctrl   *gomock.Controller
	assets *repo.MockIAssetRepo
	biz    biz.IRentalBiz
}

func (s *suiteTester) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())
	s.ctrl = gomock.NewController(s.T())
	s.assets = repo.NewMockIAssetRepo(s.ctrl)
	rentalBiz, err := NewRentalBiz(s.assets)
	s.Require().NoError(err)
	s.biz = rentalBiz
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_ListByLocation() {
	type args struct {
		ctx      contextx.Contextx
		location *model.Location
		opts     biz.ListByLocationOptions
		mock     func()
	}
	tests := []struct {
		name        string
		args        args
		wantRentals []*agg.Asset
		wantTotal   int
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRentals, gotTotal, err := s.biz.ListByLocation(tt.args.ctx, tt.args.location, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListByLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRentals, tt.wantRentals) {
				t.Errorf("ListByLocation() gotRentals = %v, want %v", gotRentals, tt.wantRentals)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("ListByLocation() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
