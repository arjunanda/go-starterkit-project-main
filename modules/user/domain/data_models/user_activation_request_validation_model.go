package data_models

type UserActivationRequestValidation struct {
	Email string `validate:"required,email"`
	Code  string `validate:"required"`
}
