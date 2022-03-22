package data_models

type UserUpdateRequest struct {
	FullName        string `json:"full_name" xml:"full_name" form:"full_name"`
	Image           string `json:"image" xml:"image" form:"image"`
	Phone           string `json:"phone" xml:"phone" form:"phone"`
	Password        string `json:"password" xml:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" xml:"confirm_password" form:"confirm_password"`
}
