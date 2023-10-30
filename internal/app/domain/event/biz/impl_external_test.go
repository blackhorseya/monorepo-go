//go:build external

package biz_test

import (
	"reflect"
	"testing"

	eventB "github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	logger  *zap.Logger
	storage repo.Storager
	biz     eventB.IEventBiz
}

func (s *suiteExternal) SetupTest() {
	s.logger = zap.NewExample()
	s.biz = biz.New(s.storage)
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) Test_impl_CreateRedPacket() {
	type args struct {
		ctx    contextx.Contextx
		who    *userM.UserAccount
		amount uint64
		count  uint32
		mock   func()
	}
	tests := []struct {
		name       string
		args       args
		wantPacket *eventM.RedPacket
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.WithLogger(s.logger)
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotPacket, err := s.biz.CreateRedPacket(tt.args.ctx, tt.args.who, tt.args.amount, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRedPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPacket, tt.wantPacket) {
				t.Errorf("CreateRedPacket() gotPacket = %v, want %v", gotPacket, tt.wantPacket)
			}
		})
	}
}
