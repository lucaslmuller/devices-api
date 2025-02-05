package e2etest

import (
	"io"
	"net/http"
)

func (s *e2eSuite) TestHealthCheck() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/health-check")

	s.Equal(http.StatusOK, r.StatusCode)

	b, _ := io.ReadAll(r.Body)
	s.Equal("OK", string(b))
}
