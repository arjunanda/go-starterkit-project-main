package handlers

import (
	respModel "go-starterkit-project/domain/data_models"
	"go-starterkit-project/modules/auth/domain/data_models"
	"go-starterkit-project/modules/auth/domain/interfaces"
	"go-starterkit-project/utils"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	AuthService interfaces.UserAuthServiceInterface
}

func NewAuthHandler(authService interfaces.UserAuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

/**
Authentication handler
*/
func (handler *AuthHandler) Authentication(c *fiber.Ctx) error {
	var request data_models.UserAuthRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	userValidation := data_models.UserAuthValidation{
		EmailValid:    request.Email,
		PasswordValid: request.Password,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.AuthService.Authenticate(&request)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Invalid authentication", re.StatusCode, err)
	}

	return utils.ApiOk(c, "Authentication successful", response)
}

/**
Get user profile
*/
func (handler *AuthHandler) GetProfile(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	response, err := handler.AuthService.GetProfile(id)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Failed to get user data", re.StatusCode, err)
	}

	return utils.ApiOk(c, "Load user successful", response)
}

/**
Refresh token
*/
func (handler *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	response, err := handler.AuthService.RefreshToken(token)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Failed to refresh token", re.StatusCode, err)
	}

	return utils.ApiOk(c, "Refresh token successful", response)
}

func (handler *AuthHandler) UpdateProfile(c *fiber.Ctx) error {

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	var request data_models.UserUpdateRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ApiUnprocessableEntity(c, "Failed body parser", err)
	}

	upload, err := c.FormFile("image")

	var image string

	if err != nil {
		image = request.Image
	} else {
		contentType := filepath.Ext(upload.Filename)

		if contentType != ".jpg" && contentType != ".png" && contentType != ".jpeg" {
			err := utils.ApiUnprocessableEntity(c, "Invalid Extension Uploaded", contentType)

			return err

		}

		err = utils.UploadImageUser(c, upload)

		if err != nil {
			return err
		}

		// fmt.Println(err)

		image = upload.Filename

	}

	userValidation := data_models.UserUpdateRequestValidation{
		FullName:        request.FullName,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
		Image:           image,
		Phone:           request.Phone,
	}

	errors := utils.ValidateStruct(userValidation)
	if errors != nil {
		return utils.ApiErrorValidation(c, "Error validation request", errors)
	}

	response, err := handler.AuthService.UpdateProfile(&request, image, id)

	if err != nil {
		re := err.(*respModel.ApiErrorResponse)
		return utils.ApiResponseError(c, "Error Update User", re.StatusCode, err)
	}

	return utils.ApiUpdated(c, "Successfully Updated User", response)
}
