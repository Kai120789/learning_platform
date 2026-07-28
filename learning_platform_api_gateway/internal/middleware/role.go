package middleware

import (
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

func MinNeededRole(minNeededRole enum.UserRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := r.Context().Value("role").(enum.UserRole)
			if !ok {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			switch minNeededRole {
			case enum.RoleStudent:
				break
			case enum.RoleTutor:
				if userRole == enum.RoleStudent {
					http.Error(w, "forbidden", http.StatusForbidden)
					return
				}
				break
			case enum.RoleAdmin:
				if userRole != enum.RoleAdmin {
					http.Error(w, "forbidden", http.StatusForbidden)
					return
				}
				break
			}

			next.ServeHTTP(w, r)
		})
	}
}
