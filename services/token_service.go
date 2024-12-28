package services

import (
	"fmt"
	"os"
	"time"

	"github.com/gabrielalbernazdev/rating-app-api/repositories"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("TOKEN_SECRET_KEY"))

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"roles": extractUserRolesByUsername(username),
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		},
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	
	return nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not extract claims from token")
	}

	return claims, nil
}

func extractUserRolesByUsername(username string) []string {
	roles ,err := repositories.FindAllRolesByUserEmail(username)
	if err != nil {
		return []string{}
	}

	var roleNames []string
    for _, role := range roles {
        roleNames = append(roleNames, role.Role)
    }

	return roleNames
}