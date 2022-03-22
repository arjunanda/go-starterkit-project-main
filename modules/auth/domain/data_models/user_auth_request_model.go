package data_models

type UserAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
