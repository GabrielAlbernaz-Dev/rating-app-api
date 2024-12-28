package services

import (
	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/repositories"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
)

func AuthVerifyUser(user models.User) (bool, error) {
	verifyUser, err := repositories.FindUserByEmail(user.Email)
	if err != nil {
		return false, err
	}

	verifyPassword := utils.CheckPasswordHash(user.Password, verifyUser.Password)
	if !verifyPassword{
		return false, nil
	}

	err = repositories.UpdateUserLastLogin(*verifyUser)
	if err != nil {
		return false, err
	}

	return true, nil
}

func AuthRegisterUser(user models.User) error {
	return repositories.CreateUser(user)
}