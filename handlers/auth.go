package handlers

import (
	"log"
	"net/http"

	"role-based-access-control-service/dto"
	"role-based-access-control-service/middleware"
	"role-based-access-control-service/pkg/auth"
	"role-based-access-control-service/pkg/apperrors"
	"role-based-access-control-service/pkg/httpx"
	"role-based-access-control-service/service"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) handleServiceError(c echo.Context, err error) error {
	status, details := apperrors.ToHTTP(err)
	if status == http.StatusInternalServerError {
		log.Printf("internal server error: method=%s path=%s err=%v", c.Request().Method, c.Path(), err)
	}

	return httpx.Failure(c, status, details)
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return httpx.Failure(c, http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(&req); err != nil {
		return httpx.Failure(c, http.StatusBadRequest, err.Error())
	}

	data, err := h.authService.Register(c.Request().Context(), req)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return httpx.Success(c, http.StatusCreated, data, "User registered successfully")
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return httpx.Failure(c, http.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(&req); err != nil {
		return httpx.Failure(c, http.StatusBadRequest, err.Error())
	}

	data, err := h.authService.Login(c.Request().Context(), req)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return httpx.Success(c, http.StatusOK, data, "Login successful")
}

func (h *AuthHandler) Me(c echo.Context) error {
	rawClaims := c.Get(middleware.UserClaimsKey)
	claims, ok := rawClaims.(*auth.Claims)
	if !ok || claims == nil {
		return httpx.Failure(c, http.StatusUnauthorized, "Unauthorized")
	}

	user, err := h.authService.Me(c.Request().Context(), claims.ID)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return httpx.Success(c, http.StatusOK, user, "Current user fetched successfully")
}
