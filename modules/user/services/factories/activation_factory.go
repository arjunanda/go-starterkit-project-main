package factories

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/user/domain/interfaces"
)

type ActionFactory struct {
	UserActivationServiceFactory interfaces.UserActivationServiceFactoryInterface
	UserForgotPassServiceFactory interfaces.UserForgotPassServiceFactoryInterface
}

type ActionFactoryInterface interface {
	Create(actionType stores.ActivationType, user *stores.User) (*stores.UserActivation, error)
}

func NewActionFactory(
	userActivationServiceFactory interfaces.UserActivationServiceFactoryInterface,
	userForgotPassServiceFactory interfaces.UserForgotPassServiceFactoryInterface,
) ActionFactoryInterface {
	return &ActionFactory{
		UserActivationServiceFactory: userActivationServiceFactory,
		UserForgotPassServiceFactory: userForgotPassServiceFactory,
	}
}

func (factory ActionFactory) Create(actionType stores.ActivationType, user *stores.User) (*stores.UserActivation, error) {

	if actionType == stores.ACTIVATION_CODE {
		userAct, err := factory.UserActivationServiceFactory.CreateUserActivation(user)

		if err != nil {
			return nil, err
		}

		return userAct, nil
	}

	userAct, err := factory.UserForgotPassServiceFactory.CreateUserForgotPass(user)

	if err != nil {
		return nil, err
	}

	return userAct, nil
}
