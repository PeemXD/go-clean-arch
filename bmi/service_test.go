package bmi

import (
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBmi(t *testing.T) {
	t.Run("should return 25.00 when given weigth 100kg and heigh 2m", func(t *testing.T) {
		given := domain.BmiRequest{
			Weight: 100,
			Height: 2,
		}
		want := domain.BmiResponse{Bmi: "25.00"}
		s := NewService()

		got := s.CalculateBmi(given)

		assert.Equal(t, want, got)
	})

	t.Run("should return 25.02 when given weigth 74.9kg and heigh 1.73m", func(t *testing.T) {
		given := domain.BmiRequest{
			Weight: 74.9,
			Height: 1.73,
		}
		want := domain.BmiResponse{Bmi: "25.03"}
		s := NewService()

		got := s.CalculateBmi(given)

		assert.Equal(t, want, got)
	})
}
