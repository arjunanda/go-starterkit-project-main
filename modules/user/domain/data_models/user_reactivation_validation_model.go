package data_models

type UserReActivationValidation struct {
	Email string `validate:"required,email"`
}
