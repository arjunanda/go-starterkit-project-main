package data_models

type UserCreateResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}
