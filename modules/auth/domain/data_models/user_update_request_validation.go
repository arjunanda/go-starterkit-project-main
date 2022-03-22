package data_models

type UserUpdateRequestValidation struct {
	FullName        string `json:"full_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Image           string `json:"image"`
	Phone           string `json:"phone"`
}
