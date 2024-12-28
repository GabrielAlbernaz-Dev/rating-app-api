package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/gabrielalbernazdev/rating-app-api/dtos"
	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
)

func HasAnyRole(roles []string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(CurrentUserKey).(*models.UserContext)
			if !ok {
				utils.WriteJson(
					w,
					dtos.GenericErrorResponseDto{Error: "User not found in context", Timestamp: time.Now()},
					http.StatusUnauthorized,
				)
				return
			}

			if !hasRole(user.Roles, roles) {
				utils.WriteJson(
					w,
					dtos.GenericErrorResponseDto{Error: "Forbidden: insufficient role", Timestamp: time.Now()},
					http.StatusUnauthorized,
				)
				return
			}

			next(w, r)
		}
	}
}

func hasRole(userRoles []string, requiredRoles []string) bool {
	for _, userRole := range userRoles {
		for _, requiredRole := range requiredRoles {
			if strings.EqualFold(userRole, requiredRole) {
				return true
			}
		}
	}
	return false
}
