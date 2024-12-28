package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/gabrielalbernazdev/rating-app-api/infra/database"
	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
)

const (
	userTable = "users"
	selectFields = "id, name, email, password, active"
	insertFields = "name, email, password, registration_date, active"
)

func FindUserByEmail(email string) (*models.User, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s u WHERE u.email = $1", selectFields, userTable)

	var user models.User
	err := database.DB.QueryRow(context.Background(), sql, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Active,
	)
	
	if err != nil {
		return nil, fmt.Errorf("error to get user: %v", err)
	}

	return &user, nil
}

func CreateUser(user models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("erro to generate hash: %v", err)
	}

	sql := fmt.Sprintf(`
		INSERT INTO users (%s)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`,insertFields)

	err = database.DB.QueryRow(context.Background(), sql, user.Name, user.Email, hashedPassword, time.Now(), user.Active).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("error to insert user: %v", err)
	}

	return nil
}

func UpdateUserLastLogin(user models.User) error {
	sql := fmt.Sprintf("UPDATE %s SET last_login = $1 WHERE id = $2",userTable)
	_, err := database.DB.Exec(context.Background(), sql, time.Now(), user.ID)
	return err;
}