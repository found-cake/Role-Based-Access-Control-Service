package middleware

import (
	"net/http"
	"strings"

	"role-based-access-control-service/pkg/auth"
	"role-based-access-control-service/pkg/httpx"

	"github.com/labstack/echo/v4"
)

const UserClaimsKey = "user_claims"

type AuthMiddleware struct {
	JWTSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{JWTSecret: jwtSecret}
}

func (m *AuthMiddleware) RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return httpx.Failure(c, http.StatusUnauthorized, "Unauthorized")
			}

			token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			if token == "" {
				return httpx.Failure(c, http.StatusUnauthorized, "Unauthorized")
			}

			claims, err := auth.ParseToken(m.JWTSecret, token)
			if err != nil {
				return httpx.Failure(c, http.StatusUnauthorized, "Invalid token")
			}

			c.Set(UserClaimsKey, claims)
			return next(c)
		}
	}
}
