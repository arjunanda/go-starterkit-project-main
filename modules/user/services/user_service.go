package services

import (
	"errors"
	respModel "go-starterkit-project/domain/data_models"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/user/domain/data_models"
	"go-starterkit-project/modules/user/domain/interfaces"
	"go-starterkit-project/modules/user/services/factories"
	"go-starterkit-project/utils"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository           interfaces.UserRepositoryInterface
	UserActivationRepository interfaces.UserActivationRepositoryInterface
	RepositoryAggregate      interfaces.RepositoryAggregateInterface
	ActionFactory            factories.ActionFactoryInterface
}

func NewUserService(
	userRepository interfaces.UserRepositoryInterface,
	userActivationRepository interfaces.UserActivationRepositoryInterface,
	repositoryAggregate interfaces.RepositoryAggregateInterface,
	factory factories.ActionFactoryInterface,
) interfaces.UserServiceInterface {
	return &UserService{
		UserRepository:           userRepository,
		UserActivationRepository: userActivationRepository,
		RepositoryAggregate:      repositoryAggregate,
		ActionFactory:            factory,
	}
}

func (service UserService) CreateUser(user *data_models.UserCreateRequest, image string) (*data_models.UserCreateResponse, error) {
	hashPassword, _ := utils.HashPassword(user.Password)

	userData := stores.User{
		FullName: user.FullName,
		Email:    user.Email,
		Image:    image,
		Phone:    user.Phone,
		Password: hashPassword,
	}

	activationCode := utils.StringWithCharset(32)

	userAvtivate := stores.UserActivation{
		Code:    activationCode,
		ActType: stores.ACTIVATION_CODE,
	}

	result, err := service.RepositoryAggregate.CreateUser(&userData, &userAvtivate)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return &data_models.UserCreateResponse{}, &respModel.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    "User already register",
			}
		}

		return &data_models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong with our server",
		}
	}

	sendMail := respModel.Mail{
		To:           []string{user.Email},
		Subject:      "User Activation",
		TemplateHtml: "user_activation.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": activationCode,
		},
	}

	utils.SendMail(&sendMail)

	response := data_models.UserCreateResponse{
		ID:       userData.ID.String(),
		FullName: result.FullName,
		Email:    result.Email,
		Image:    result.Image,
		Phone:    result.Phone,
		IsActive: userData.IsActive,
	}

	return &response, nil
}

func (service UserService) UserActivation(email string, code string) (*data_models.UserCreateResponse, error) {
	var user stores.User
	var userAct stores.UserActivation

	errUser := service.UserRepository.FindUserByEmail(&user, email).Error

	log.Println(user)

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &data_models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "User not found",
		}
	}

	if user.IsActive {
		return &data_models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "User already active",
		}
	}

	errAct := service.UserActivationRepository.FindUserActivationCode(&userAct, user.ID.String(), code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return &data_models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "Activation code not found",
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return &data_models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    "The activation code has expired",
		}
	}

	userNew, errUserNew := service.RepositoryAggregate.UpdateUserActivation(user.ID.String(), true)

	if errUserNew != nil {
		return &data_models.UserCreateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Cannot activated user",
		}
	}

	service.RepositoryAggregate.UpdateActivationCodeUsed(user.ID.String(), code)

	response := data_models.UserCreateResponse{
		ID:       userNew.ID.String(),
		FullName: userNew.FullName,
		Email:    userNew.Email,
		Phone:    userNew.Phone,
		IsActive: userNew.IsActive,
	}

	return &response, nil
}

func (service UserService) CreateUserActivation(email string, actType stores.ActivationType) (map[string]interface{}, error) {
	var user stores.User

	errUser := service.UserRepository.FindUserByEmail(&user, email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "User not found",
		}
	}

	if user.IsActive && actType == stores.ACTIVATION_CODE {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "User already active",
		}
	}

	_, errActFactory := service.ActionFactory.Create(actType, &user)

	if errActFactory != nil {
		return nil, errActFactory
	}

	return map[string]interface{}{}, nil
}

func (service UserService) UpdatePassword(forgotPassReq *data_models.UserForgotPassActRequest) (map[string]interface{}, error) {
	var user stores.User
	var userAct stores.UserActivation

	if forgotPassReq.Password != forgotPassReq.RepeatPassword {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Password validation does not match",
		}
	}

	errUser := service.UserRepository.FindUserByEmail(&user, forgotPassReq.Email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "User not found",
		}
	}

	errAct := service.UserActivationRepository.FindUserActivationCode(&userAct, user.ID.String(), forgotPassReq.Code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "Activation code not found",
		}
	}

	if userAct.IsUsed {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Password reset code has been used",
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    "The activation code has expired",
		}
	}

	go func() {
		hashPassword, _ := utils.HashPassword(user.Password)

		userData := stores.User{
			FullName: user.FullName,
			Email:    user.Email,
			Phone:    user.Phone,
			Password: hashPassword,
		}

		service.UserRepository.UpdatePassword(&userData)
		service.RepositoryAggregate.UpdateActivationCodeUsed(user.ID.String(), forgotPassReq.Code)
	}()

	return map[string]interface{}{}, nil
}
