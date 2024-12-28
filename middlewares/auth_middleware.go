package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gabrielalbernazdev/rating-app-api/dtos"
	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/services"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
)

type ContextKey string

const CurrentUserKey ContextKey = "currentUser"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			utils.WriteJson(
				w,
				dtos.GenericErrorResponseDto{Error: "Missing authorization header",Timestamp: time.Now()},
				http.StatusUnauthorized,
			)
			return
		}

		tokenString = tokenString[len("Bearer "):]

		err := services.VerifyToken(tokenString)
		if err != nil {
			utils.WriteJson(
				w,
				dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now(),},
				http.StatusUnauthorized,
			)
			return
		}

		claims, err := services.GetClaimsFromToken(tokenString)
		if err != nil {
			utils.WriteJson(
				w,
				dtos.GenericErrorResponseDto{Error: "Error to extract claims", Timestamp: time.Now(),},
				http.StatusInternalServerError,
			)
			return
		}

		username := claims["username"].(string)

		var roles []string
		
		if claims["roles"] != nil {
			for _, role := range claims["roles"].([]interface{}) {
				roleStr, ok := role.(string)
				if !ok {
					continue
				}
				roles = append(roles, roleStr)
			}
		}

		currentUser := &models.UserContext{
			Username: username,
			Roles:    roles,
		}

		ctx := context.WithValue(r.Context(), CurrentUserKey, currentUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
