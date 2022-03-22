package interfaces

import (
	"go-starterkit-project/domain/stores"
)

type RepositoryAggregateInterface interface {
	CreateUser(user *stores.User, userActivate *stores.UserActivation) (*stores.User, error)

	UpdateUser(id string, stat bool) (*stores.User, error)

	UpdateUserActivation(id string, stat bool) (*stores.User, error)

	UpdatePassword(id string, password string) (*stores.User, error)

	UpdateActivationCodeUsed(userId string, code string) (*stores.UserActivation, error)
}
