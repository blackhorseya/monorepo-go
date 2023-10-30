package biz_test

import (
	"errors"
	"reflect"
	"testing"

	eventB "github.com/blackhorseya/monorepo-go/entity/domain/event/biz"
	eventM "github.com/blackhorseya/monorepo-go/entity/domain/event/model"
	userM "github.com/blackhorseya/monorepo-go/entity/domain/user/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/event/biz/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	logger  *zap.Logger
	ctrl    *gomock.Controller
	storage *repo.MockStorager
	biz     eventB.IEventBiz
}

func (s *suiteTester) SetupTest() {
	s.logger = zap.NewExample()
	s.ctrl = gomock.NewController(s.T())
	s.storage = repo.NewMockStorager(s.ctrl)
	s.biz = biz.New(s.storage)
}

func (s *suiteTester) TearDownSuite() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_CreateRedPacket() {
	tester := &userM.UserAccount{
		Id:          "1",
		Username:    "tester",
		Email:       "tester@gmail.com",
		Password:    "",
		AccessToken: "",
		IdToken:     "",
		Profile:     nil,
	}

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
		{
			name:       "who is nil then error",
			args:       args{who: nil, amount: 100, count: 10},
			wantPacket: nil,
			wantErr:    true,
		},
		{
			name:       "invalid amount then error",
			args:       args{who: tester, amount: 0, count: 10},
			wantPacket: nil,
			wantErr:    true,
		},
		{
			name:       "invalid count then error",
			args:       args{who: tester, amount: 100, count: 0},
			wantPacket: nil,
			wantErr:    true,
		},
		{
			name: "create a new red packet then error",
			args: args{who: tester, amount: 100, count: 10, mock: func() {
				s.storage.EXPECT().CreateRedPacket(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			}},
			wantPacket: nil,
			wantErr:    true,
		},
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
