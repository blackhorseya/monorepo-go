package memory_test

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/repo"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/repo/memory"
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
	s.storage = memory.NewStorager()
}

func (s *suiteTester) TearDownTest() {
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetURLRecordByShortURL() {
	type args struct {
		ctx      contextx.Contextx
		shortURL string
		mock     func()
	}
	tests := []struct {
		name       string
		args       args
		wantRecord *model.ShortenedUrl
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

			gotRecord, err := s.storage.GetURLRecordByShortURL(tt.args.ctx, tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetURLRecordByShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("GetURLRecordByShortURL() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}

func (s *suiteTester) Test_impl_CreateURLRecord() {
	type args struct {
		ctx    contextx.Contextx
		record *model.ShortenedUrl
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.WithLogger(s.logger)
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.storage.CreateURLRecord(tt.args.ctx, tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("CreateURLRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
