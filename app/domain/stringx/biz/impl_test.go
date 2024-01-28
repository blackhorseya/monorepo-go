package biz_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/stringx/biz"
	stringxB "github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteUnit struct {
	suite.Suite
	logger *zap.Logger

	biz stringxB.IStringBiz
}

func (s *suiteUnit) SetupTest() {
	s.logger = zap.NewExample()

	s.biz = biz.New()
}

func (s *suiteUnit) TearDownTest() {
}

func TestUnit(t *testing.T) {
	suite.Run(t, new(suiteUnit))
}

func (s *suiteUnit) Test_impl_Uppercase() {
	type args struct {
		ctx   contextx.Contextx
		value string
		mock  func()
	}
	tests := []struct {
		name      string
		args      args
		wantUpper string
		wantErr   bool
	}{
		{
			name:      "empty string then error",
			args:      args{value: ""},
			wantUpper: "",
			wantErr:   true,
		},
		{
			name:      "normal string then success",
			args:      args{value: "hello"},
			wantUpper: "HELLO",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.WithLogger(s.logger)
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotUpper, err := s.biz.Uppercase(tt.args.ctx, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uppercase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUpper != tt.wantUpper {
				t.Errorf("Uppercase() gotUpper = %v, want %v", gotUpper, tt.wantUpper)
			}
		})
	}
}

func (s *suiteUnit) Test_impl_Count() {
	type args struct {
		ctx   contextx.Contextx
		value string
		mock  func()
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty string then 0",
			args: args{value: ""},
			want: 0,
		},
		{
			name: "normal string then success",
			args: args{value: "hello"},
			want: 5,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.WithLogger(s.logger)
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if got := s.biz.Count(tt.args.ctx, tt.args.value); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
