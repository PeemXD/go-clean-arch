package rest

import (
	"fmt"
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type BmiService interface {
	CalculateBmi(domain.BmiRequest) domain.BmiResponse
}

type BmiHandler struct {
	svc BmiService
}

func NewBmiHandler(e *echo.Echo, svc BmiService) {
	h := &BmiHandler{svc: svc}
	e.POST("/bmi/calculate", h.CalculateBmi)
}

func (h *BmiHandler) CalculateBmi(c echo.Context) error {
	var body domain.BmiRequest
	if err := c.Bind(&body); err != nil {
		log.Error(fmt.Sprintf("[CalculateBmi handler] unmarshall body error: %s", err))
		return c.JSON(http.StatusBadRequest, nil)
	}

	if body.Weight <= 0 || body.Height <= 0 {
		log.Error("[CalculateBmi handler] weight and height must greater than 0")
		return c.JSON(http.StatusBadRequest, "weight and height must greater than 0")
	}

	return c.JSON(http.StatusOK, h.svc.CalculateBmi(body))
}
