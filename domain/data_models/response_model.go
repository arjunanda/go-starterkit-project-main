package data_models

import "fmt"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Message     string `json:"message"`
}

type ApiErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func (r *ApiErrorResponse) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Message)
}

func (r *ApiErrorResponse) GetStatusCode() int {
	return r.StatusCode
}
