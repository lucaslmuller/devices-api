package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lucaslmuller/technical-test/internal/app/device/domain/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ServiceSuite struct {
	ctx            context.Context
	cache          *mocks.MockCache
	ctrl           *gomock.Controller
	mockRepository *mocks.MockIRepository
	service        *DeviceService
	suite.Suite
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.cache = mocks.NewMockCache(s.ctrl)
	s.mockRepository = mocks.NewMockIRepository(s.ctrl)
	l, _ := zap.NewProduction()
	logger := l.Sugar()
	s.service = NewService(s.mockRepository, s.cache, logger)
	s.ctx = context.Background()
}
