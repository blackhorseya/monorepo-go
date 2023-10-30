package memory_test

import (
	"reflect"
	"testing"

	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo/memory"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	logger  *zap.Logger
	storage repo.Storager
}

func (s *suiteTester) SetupTest() {
	s.logger = zap.NewExample()
	s.storage = memory.New()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_ListRedPacket() {
	type args struct {
		ctx  contextx.Contextx
		cond repo.ListRedPacketCondition
		mock func()
	}
	tests := []struct {
		name        string
		args        args
		wantPackets []*eventM.RedPacket
		wantErr     bool
	}{
		{
			name:        "list not found then return empty",
			args:        args{},
			wantPackets: nil,
			wantErr:     false,
		},
		{
			name: "list found return packets",
			args: args{mock: func() {
				_ = s.storage.CreateRedPacket(contextx.Background(), &eventM.RedPacket{
					Id: "test",
				})
			}},
			wantPackets: []*eventM.RedPacket{{Id: "test"}},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.WithLogger(s.logger)
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotPackets, err := s.storage.ListRedPacket(tt.args.ctx, tt.args.cond)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRedPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPackets, tt.wantPackets) {
				t.Errorf("ListRedPacket() gotPackets = %v, want %v", gotPackets, tt.wantPackets)
			}
		})
	}
}
