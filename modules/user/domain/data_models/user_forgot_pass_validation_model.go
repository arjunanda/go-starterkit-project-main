package data_models

type UserForgotPassValidation struct {
	Email string `validate:"required,email"`
}
