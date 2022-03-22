package data_models

type UserAuthResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
	Token    string `json:"token"`
	Exp      int64  `json:"expires"`
}
