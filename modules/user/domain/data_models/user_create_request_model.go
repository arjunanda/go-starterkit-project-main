package data_models

type UserCreateRequest struct {
	FullName string `json:"full_name" xml:"full_name" form:"full_name"`
	Password string `json:"password" xml:"password" form:"password"`
	Email    string `json:"email" xml:"email" form:"email"`
	Image    string `json:"image" xml:"image" form:"image"`
	Phone    string `json:"phone" xml:"phone" form:"phone"`
}
