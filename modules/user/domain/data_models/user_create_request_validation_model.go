package data_models

type UserCreateRequestValidation struct {
	FullName string `validate:"required,min=3"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
	Image    string `validate:"required"`
	Phone    string `validate:"required,numeric,min=10"`
}
