package data_models

type UserAuthValidation struct {
	EmailValid    string `validate:"required"`
	PasswordValid string `validate:"required"`
}
