package services

import (
	"errors"
	respModel "go-starterkit-project/domain/data_models"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/auth/domain/data_models"
	"go-starterkit-project/modules/auth/domain/interfaces"
	userInterface "go-starterkit-project/modules/user/domain/interfaces"
	"go-starterkit-project/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthService struct {
	UserRepository userInterface.UserRepositoryInterface
}

func NewAuthService(userRepository userInterface.UserRepositoryInterface) interfaces.UserAuthServiceInterface {
	return &AuthService{
		UserRepository: userRepository,
	}
}

/**
This function is used to handle authentication
*/
func (service AuthService) Authenticate(auth *data_models.UserAuthRequest) (*data_models.UserAuthResponse, error) {
	var user stores.User

	// Get user by email
	errUser := service.UserRepository.FindUserByEmail(&user, auth.Email).Error

	// Check if the user is not found
	// then displayan error message
	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &data_models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "Invalid email or password",
		}
	}

	// Check if a query operation error occurs
	if errUser != nil {
		return &data_models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	// Check if the user status is not active
	if !user.IsActive {
		return &data_models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "User is not active, please activate the user first",
		}
	}

	// Match password hashes
	match := utils.CheckPasswordHash(auth.Password, user.Password)

	// Check if it doesn't match, show an error message
	if !match {
		return &data_models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    "Invalid email or password",
		}
	}

	token, exp, errToken := utils.CreateToken(user.ID.String())

	if errToken != nil {
		return &data_models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Error token",
		}
	}

	// Set response message to succeed
	response := data_models.UserAuthResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Image:    user.Image,
		Phone:    user.Phone,
		IsActive: user.IsActive,
		Token:    token,
		Exp:      exp,
	}

	return &response, nil
}

/**
This function is used to authorize users and display logged in user data
*/
func (service AuthService) GetProfile(id string) (*data_models.UserAuthProfileResponse, error) {
	var user stores.User

	// Get user from database
	errUser := service.UserRepository.FindUserById(&user, id).Error

	// Check if there is a query error
	if errUser != nil {
		return &data_models.UserAuthProfileResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	// Set response message
	response := data_models.UserAuthProfileResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Image:    user.Image,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}

	return &response, nil
}

/**
This function is used to refresh token
*/
func (service AuthService) RefreshToken(tokenUser *jwt.Token) (*data_models.UserAuthResponse, error) {
	var user stores.User

	// Get data from token then convert to string
	beforeClaims := tokenUser.Claims.(jwt.MapClaims)
	id := beforeClaims["id"].(string)

	// Get user data
	errUser := service.UserRepository.FindUserById(&user, id).Error

	// Check if something went wrong with query
	if errUser != nil {
		return &data_models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	token, exp, errToken := utils.CreateToken(user.ID.String())
	if errToken != nil {
		return &data_models.UserAuthResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Error token",
		}
	}

	// Set response message
	response := data_models.UserAuthResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Image:    user.Image,
		Phone:    user.Phone,
		IsActive: user.IsActive,
		Token:    token,
		Exp:      exp,
	}

	return &response, nil
}

func (service AuthService) UpdateProfile(auth *data_models.UserUpdateRequest, image string, id string) (*data_models.UserUpdateResponse, error) {

	var user stores.User

	// Get user data
	errUser := service.UserRepository.FindUserById(&user, id).Error

	// Check if something went wrong with query
	if errUser != nil {
		return &data_models.UserUpdateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Something went wrong",
		}
	}

	if auth.Password != "" && auth.Password != auth.ConfirmPassword {
		return &data_models.UserUpdateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Password And Password Confirmation Are Not The Samed",
		}
	}

	if image != "" {
		err := utils.DeleteImageUser(user.Image)

		if err != nil {
			log.Printf("Error In Delete File : %s", err)
		}
		user.Image = image
	}
	if auth.Password != "" {
		user.Password, _ = utils.HashPassword(auth.Password)
	}
	if auth.Phone != "" {
		user.Phone = auth.Phone
	}
	if auth.FullName != "" {
		user.FullName = auth.FullName
	}

	// ErrUpdateUser := service.UserRepository.UpdateUser(&user)

	ErrUpdateUser := service.UserRepository.UpdateUser(&user).Error

	// fmt.Println(ErrUpdateUser)

	if ErrUpdateUser != nil {
		return &data_models.UserUpdateResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Cannot Update users",
		}
	}

	response := data_models.UserUpdateResponse{
		ID:       user.ID.String(),
		FullName: user.FullName,
		Email:    user.Email,
		Image:    user.Image,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}

	return &response, nil
}
