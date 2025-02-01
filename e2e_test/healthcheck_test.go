package e2etest

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestHealthCheck() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/health-check")

	s.Equal(http.StatusOK, r.StatusCode)

	b, _ := io.ReadAll(r.Body)
	s.Equal("OK", string(b))
}
