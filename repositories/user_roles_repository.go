package repositories

import (
	"context"
	"fmt"

	"github.com/gabrielalbernazdev/rating-app-api/infra/database"
	"github.com/gabrielalbernazdev/rating-app-api/models"
)

func FindAllRolesByUserEmail(email string) ([]models.Role, error) {
    sql := `
        SELECT r.id, r.role
        FROM roles r
        JOIN user_roles ur ON ur.role_id = r.id
		JOIN users u ON u.id = ur.user_id
        WHERE u.email = $1
    `
    
    rows, err := database.DB.Query(context.Background(), sql, email)
    if err != nil {
        return nil, fmt.Errorf("error to get roles for user with email: %s: %v", email, err)
    }
    defer rows.Close()

    var roles []models.Role
    for rows.Next() {
        var role models.Role
        if err := rows.Scan(&role.ID, &role.Role); err != nil {
            return nil, fmt.Errorf("error scanning role for user with email %s: %v", email, err)
        }
        roles = append(roles, role)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating over roles: %v", err)
    }

    return roles, nil
}