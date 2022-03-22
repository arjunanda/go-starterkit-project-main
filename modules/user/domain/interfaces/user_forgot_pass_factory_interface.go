package interfaces

import "go-starterkit-project/domain/stores"

type UserForgotPassServiceFactoryInterface interface {
	CreateUserForgotPass(user *stores.User) (*stores.UserActivation, error)
}
