package biz_test

import (
	"errors"
	"reflect"
	"testing"

	shortB "github.com/blackhorseya/monorepo-go/entity/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/entity/domain/shortening/model"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/biz"
	"github.com/blackhorseya/monorepo-go/internal/app/domain/shortening/repo"
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
	biz     shortB.IShorteningBiz
}

func (s *suiteTester) SetupTest() {
	s.logger = zap.NewExample()
	s.ctrl = gomock.NewController(s.T())
	s.storage = repo.NewMockStorager(s.ctrl)
	s.biz = biz.NewShortening(s.storage)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetURLRecordByShortURL() {
	shortURL1 := "https://www.google.com"
	record1 := &model.ShortenedUrl{
		Id:          0,
		OriginalUrl: shortURL1,
		ShortUrl:    shortURL1,
		CreatedAt:   nil,
		ExpiredAt:   nil,
	}

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
		{
			name: "error",
			args: args{shortURL: shortURL1, mock: func() {
				s.storage.EXPECT().GetURLRecordByShortURL(gomock.Any(), shortURL1).
					Return(nil, errors.New("error")).Times(1)
			}},
			wantRecord: nil,
			wantErr:    true,
		},
		{
			name: "ok",
			args: args{shortURL: shortURL1, mock: func() {
				s.storage.EXPECT().GetURLRecordByShortURL(gomock.Any(), shortURL1).
					Return(record1, nil).Times(1)
			}},
			wantRecord: record1,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.WithLogger(s.logger)
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecord, err := s.biz.GetURLRecordByShortURL(tt.args.ctx, tt.args.shortURL)
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

func (s *suiteTester) Test_impl_CreateShortenedURL() {
	longURL1 := "https://www.google.com"

	type args struct {
		ctx         contextx.Contextx
		originalURL string
		mock        func()
	}
	tests := []struct {
		name       string
		args       args
		wantRecord *model.ShortenedUrl
		wantErr    bool
	}{
		{
			name: "create a record then error",
			args: args{originalURL: longURL1, mock: func() {
				s.storage.EXPECT().CreateURLRecord(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
			}},
			wantRecord: nil,
			wantErr:    true,
		},
		{
			name: "ok",
			args: args{originalURL: longURL1, mock: func() {
				s.storage.EXPECT().CreateURLRecord(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			}},
			wantRecord: &model.ShortenedUrl{
				Id:          0,
				OriginalUrl: longURL1,
				ShortUrl:    longURL1,
				CreatedAt:   nil,
				ExpiredAt:   nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.WithLogger(s.logger)
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecord, err := s.biz.CreateShortenedURL(tt.args.ctx, tt.args.originalURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShortenedURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRecord != nil {
				gotRecord.CreatedAt = nil
				gotRecord.ExpiredAt = nil
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("CreateShortenedURL() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}
