package data_models

type UserForgotPassActValidation struct {
	Email          string `validate:"required,email"`
	Password       string `validate:"required,min=8"`
	RepeatPassword string `validate:"required,min=8"`
	Code           string `validate:"required"`
}
