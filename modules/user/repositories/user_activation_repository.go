package repositories

import (
	"go-starterkit-project/database/driver"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/user/domain/interfaces"

	"gorm.io/gorm"
)

type UserActivationRepository struct {
	DB *gorm.DB
}

func NewUserActivationRepository() interfaces.UserActivationRepositoryInterface {
	return &UserActivationRepository{
		DB: driver.DB,
	}
}

func (repository UserActivationRepository) FindUserActivationCode(
	userActivation *stores.UserActivation,
	userId string,
	code string,
) *gorm.DB {
	return repository.DB.First(&userActivation, "user_id = ? AND code = ?", userId, code)
}

func (repository UserActivationRepository) CreateUserActivation(userActivate *stores.UserActivation) *gorm.DB {
	return repository.DB.Create(&userActivate)
}

func (repository UserActivationRepository) UpdateActivationCodeUsed(userActivate *stores.UserActivation) *gorm.DB {
	return repository.DB.Save(&userActivate)
}
