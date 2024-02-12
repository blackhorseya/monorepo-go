//go:build external

package restful

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	client Dialer
}

func (s *suiteExternal) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	s.Require().NoError(err)

	configx.ReplaceApplication(configx.C.Orianna)

	client, err := NewClient()
	s.Require().NoError(err)
	s.client = client
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) Test_impl_GetStockBySymbol() {
	type args struct {
		ctx    contextx.Contextx
		symbol string
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		want    agg.Stock
		wantErr bool
	}{
		{
			name:    "2330",
			args:    args{symbol: "2330"},
			want:    agg.Stock{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.client.GetStockBySymbol(tt.args.ctx, tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStockBySymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.args.ctx.Debug("got", zap.Any("got", &got))
		})
	}
}
