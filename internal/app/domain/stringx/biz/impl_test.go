package biz_test

import (
	"testing"

	"github.com/blackhorseya/monorepo-go/entity/domain/stringx/biz"
	"github.com/stretchr/testify/suite"
)

type suiteUnit struct {
	suite.Suite

	biz biz.IStringBiz
}

func TestUnit(t *testing.T) {
	suite.Run(t, new(suiteUnit))
}
