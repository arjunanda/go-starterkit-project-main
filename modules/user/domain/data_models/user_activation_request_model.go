package data_models

type UserActivationRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
