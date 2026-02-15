package apperrors

import (
	"errors"
	"net/http"
)

func ToHTTP(err error) (status int, details string) {
	switch {
	case errors.Is(err, ErrEmailAlreadyExists):
		return http.StatusBadRequest, "Email already exists"
	case errors.Is(err, ErrInvalidCredentials):
		return http.StatusUnauthorized, "Invalid credentials"
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized, "Unauthorized"
	default:
		return http.StatusInternalServerError, "Internal server error"
	}
}
