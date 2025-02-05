package endpoint

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/mocks"
	"github.com/stretchr/testify/suite"
)

type EndpointSuite struct {
	ctx         context.Context
	cache       *mocks.MockCache
	ctrl        *gomock.Controller
	mockService *mocks.MockIService
	endpoints   *DeviceEndpoints
	suite.Suite
}

func TestEndpointSuite(t *testing.T) {
	suite.Run(t, new(EndpointSuite))
}

func (s *EndpointSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.cache = mocks.NewMockCache(s.ctrl)
	s.mockService = mocks.NewMockIService(s.ctrl)

	// l, _ := zap.NewProduction()
	// logger := l.Sugar()

	s.endpoints = NewDeviceEndpoints(s.mockService)
	s.ctx = context.Background()
}
