package interfaces

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/user/domain/data_models"
)

type UserServiceInterface interface {
	CreateUser(user *data_models.UserCreateRequest, image string) (*data_models.UserCreateResponse, error)

	UserActivation(email string, code string) (*data_models.UserCreateResponse, error)

	CreateUserActivation(email string, actType stores.ActivationType) (map[string]interface{}, error)

	UpdatePassword(forgotPassReq *data_models.UserForgotPassActRequest) (map[string]interface{}, error)
}
