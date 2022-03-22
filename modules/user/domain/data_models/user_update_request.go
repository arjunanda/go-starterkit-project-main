package data_models

type UserUpdateRequest struct {
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Image           string `json:"image"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Code            string `json:"code"`
}
