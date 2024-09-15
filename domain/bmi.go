package domain

type BmiRequest struct {
	Weight float64
	Height float64
}

type BmiResponse struct {
	Bmi string `json:"bmi"`
}
