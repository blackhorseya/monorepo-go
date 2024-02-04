package mongodb

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/monorepo-go/app/domain/issue/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/issue/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
)

type suiteTester struct {
	suite.Suite

	storage repo.Storager
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_List() {
	type args struct {
		ctx  contextx.Contextx
		opts repo.ListOptions
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantTodos []*model.Ticket
		wantTotal int
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTodos, gotTotal, err := s.storage.List(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTodos, tt.wantTodos) {
				t.Errorf("List() gotTodos = %v, want %v", gotTodos, tt.wantTodos)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("List() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
