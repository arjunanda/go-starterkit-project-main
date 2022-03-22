package interfaces

import (
	"go-starterkit-project/domain/stores"
)

type UserActivationServiceFactoryInterface interface {
	CreateUserActivation(user *stores.User) (*stores.UserActivation, error)
}
