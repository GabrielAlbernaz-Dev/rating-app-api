package services

import (
	"fmt"

	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/repositories"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
)

func AuthVerifyUser(user models.User) (bool, error) {
	var err error
	verifyUser, _ := repositories.FindUserByEmail(user.Email)

	if verifyUser == nil || verifyUser.ID == 0 {
		return false, fmt.Errorf("user with email %s not found", user.Email)
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