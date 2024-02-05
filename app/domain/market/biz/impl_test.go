package biz_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/app/domain/market/repo"
	marketB "github.com/blackhorseya/monorepo-go/entity/domain/market/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/blackhorseya/monorepo-go/pkg/finmindx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	ctrl    *gomock.Controller
	finmind *finmindx.MockDialer
	storage *repo.MockStorager
	biz     marketB.IMarketBiz
}

func (s *suiteTester) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())
	s.ctrl = gomock.NewController(s.T())
	s.finmind = finmindx.NewMockDialer(s.ctrl)
	s.storage = repo.NewMockStorager(s.ctrl)
	marketBiz, err := biz.NewMarketBiz(s.finmind, s.storage)
	s.NoError(err)
	s.biz = marketBiz
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetStockBySymbol() {
	stock1 := &model.Stock{
		Symbol: "2330",
		Name:   "",
		Price:  100,
	}

	type args struct {
		ctx    contextx.Contextx
		symbol string
		mock   func()
	}
	tests := []struct {
		name      string
		args      args
		wantStock *model.Stock
		wantErr   bool
	}{
		{
			name: "get by symbol then error",
			args: args{symbol: stock1.Symbol, mock: func() {
				s.finmind.EXPECT().TaiwanStockPrice(
					gomock.Any(),
					stock1.Symbol,
					gomock.Any(),
					gomock.Any(),
				).Return(nil, errors.New("mock error")).Times(1)
			}},
			wantStock: nil,
			wantErr:   true,
		},
		{
			name: "get by symbol then ok",
			args: args{symbol: stock1.Symbol, mock: func() {
				s.finmind.EXPECT().TaiwanStockPrice(
					gomock.Any(),
					stock1.Symbol,
					gomock.Any(),
					gomock.Any(),
				).Return(&finmindx.TaiwanStockPriceResponse{
					Data: []*finmindx.TaiwanStockPrice{{StockID: stock1.Symbol, Close: stock1.Price}},
				}, nil).Times(1)
			}},
			wantStock: stock1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotStock, err := s.biz.GetStockBySymbol(tt.args.ctx, tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStockBySymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStock, tt.wantStock) {
				t.Errorf("GetStockBySymbol() gotStock = %v, want %v", gotStock, tt.wantStock)
			}
		})
	}
}
