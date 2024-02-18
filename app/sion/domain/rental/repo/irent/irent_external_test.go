//go:build external

package irent

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/entity/sion/domain/rental/repo"
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteExternal struct {
	suite.Suite

	repo repo.IAssetRepo
}

func (s *suiteExternal) SetupTest() {
	zap.ReplaceGlobals(zap.NewExample())

	err := configx.Load("", "sean")
	s.NoError(err)

	configx.ReplaceApplication(configx.C.Sion)

	s.repo, err = NewAssetRepo()
	s.NoError(err)
}

func TestExternal(t *testing.T) {
	suite.Run(t, new(suiteExternal))
}

func (s *suiteExternal) TestFetchAvailableCars() {
	ctx := contextx.Background()
	cars, err := s.repo.FetchAvailableCars(ctx)
	s.NoError(err)
	ctx.Debug("cars", zap.Any("cars", cars))
}
