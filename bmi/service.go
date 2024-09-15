package bmi

import (
	"fmt"
	"math"

	"github.com/bxcodec/go-clean-arch/domain"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CalculateBmi(data domain.BmiRequest) domain.BmiResponse {
	bmi := data.Weight / math.Pow(data.Height, 2)

	return domain.BmiResponse{Bmi: fmt.Sprintf("%.2f", bmi)}
}
