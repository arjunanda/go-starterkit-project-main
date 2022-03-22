package interfaces

import (
	"go-starterkit-project/modules/auth/domain/data_models"

	"github.com/golang-jwt/jwt/v4"
)

type UserAuthServiceInterface interface {
	Authenticate(auth *data_models.UserAuthRequest) (*data_models.UserAuthResponse, error)

	GetProfile(id string) (*data_models.UserAuthProfileResponse, error)

	UpdateProfile(auth *data_models.UserUpdateRequest, image string, id string) (*data_models.UserUpdateResponse, error)

	RefreshToken(token *jwt.Token) (*data_models.UserAuthResponse, error)
}
