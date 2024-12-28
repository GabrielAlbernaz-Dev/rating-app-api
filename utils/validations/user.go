package validations

import (
	"fmt"

	"github.com/gabrielalbernazdev/rating-app-api/models"
)

func ValidateUserLoginBody(user models.User) error {
	if(user.Email == "") {
		return fmt.Errorf("email is empty")
	}

	if(user.Password == "") {
		return fmt.Errorf("password is empty")
	}

	return nil
}

func ValidateUserRegisterBody(user models.User) error {
	if user.Email == "" {
		return fmt.Errorf("email is empty")
	}

	if user.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if user.Password == "" {
		return fmt.Errorf("password is empty")
	}

	return nil
}