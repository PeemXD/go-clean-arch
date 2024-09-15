package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/rest/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBmi(t *testing.T) {
	t.Run("should return status ok", func(t *testing.T) {
		s := mocks.BmiService{}
		bmiReq := domain.BmiRequest{Weight: 100, Height: 2}
		bmiRes := domain.BmiResponse{Bmi: "25.00"}
		s.On("CalculateBmi", bmiReq).Return(bmiRes)

		body := domain.BmiRequest{Weight: 100, Height: 2}
		given, err := json.Marshal(body)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), echo.POST, "/bmi/calculate", strings.NewReader(string(given)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := BmiHandler{svc: &s}
		want := domain.BmiResponse{Bmi: "25.00"}

		err = h.CalculateBmi(c)
		assert.NoError(t, err)
		var got domain.BmiResponse
		_ = json.Unmarshal(rec.Body.Bytes(), &got)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want, got)
	})
	t.Run("should return status bad request when given context type is not json", func(t *testing.T) {
		s := mocks.BmiService{}

		body := domain.BmiRequest{Weight: 100, Height: 2}
		given, err := json.Marshal(body)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), echo.POST, "/bmi/calculate", strings.NewReader(string(given)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJavaScript)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := BmiHandler{svc: &s}

		err = h.CalculateBmi(c)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return status bad request when given height 0", func(t *testing.T) {
		s := mocks.BmiService{}

		body := domain.BmiRequest{Weight: 100, Height: 0}
		given, err := json.Marshal(body)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), echo.POST, "/bmi/calculate", strings.NewReader(string(given)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := BmiHandler{svc: &s}

		err = h.CalculateBmi(c)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return status bad request when given height -1", func(t *testing.T) {
		s := mocks.BmiService{}

		body := domain.BmiRequest{Weight: 100, Height: -1}
		given, err := json.Marshal(body)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), echo.POST, "/bmi/calculate", strings.NewReader(string(given)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := BmiHandler{svc: &s}

		err = h.CalculateBmi(c)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return status bad request when given Weight 0", func(t *testing.T) {
		s := mocks.BmiService{}

		body := domain.BmiRequest{Weight: 0, Height: 2}
		given, err := json.Marshal(body)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), echo.POST, "/bmi/calculate", strings.NewReader(string(given)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := BmiHandler{svc: &s}

		err = h.CalculateBmi(c)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return status bad request when given Weight -1", func(t *testing.T) {
		s := mocks.BmiService{}

		body := domain.BmiRequest{Weight: -1, Height: 2}
		given, err := json.Marshal(body)
		assert.NoError(t, err)

		e := echo.New()
		req, err := http.NewRequestWithContext(context.TODO(), echo.POST, "/bmi/calculate", strings.NewReader(string(given)))
		assert.NoError(t, err)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := BmiHandler{svc: &s}

		err = h.CalculateBmi(c)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
