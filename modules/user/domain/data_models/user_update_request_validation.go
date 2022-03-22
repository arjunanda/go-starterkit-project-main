package data_models

type UserUpdateRequestValidation struct {
	FullName        string `validate:"required,min=3"`
	Password        string
	ConfirmPassword string
	Email           string `validate:"required,email"`
	Image           string
	Phone           string `validate:"required,numeric,min=10"`
}
