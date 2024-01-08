package biz_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteTester struct {
	suite.Suite
}

func (s *suiteTester) SetupTest() {
}

func (s *suiteTester) TearDownTest() {
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
