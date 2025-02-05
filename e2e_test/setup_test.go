package e2etest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type e2eSuite struct {
	ctx context.Context

	suite.Suite
}

func TestE2ESuite(t *testing.T) {
	suite.Run(t, new(e2eSuite))
}

func (s *e2eSuite) SetupSuite() {
	s.ctx = context.Background()
}
