package mocks

import (
	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

type BmiService struct {
	mock.Mock
}

func (m *BmiService) CalculateBmi(data domain.BmiRequest) domain.BmiResponse {
	args := m.Called(data)
	return args.Get(0).(domain.BmiResponse)
}
