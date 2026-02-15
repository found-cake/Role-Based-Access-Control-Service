package httpx

import (
	"role-based-access-control-service/dto"

	"github.com/labstack/echo/v4"
)

func JSON(c echo.Context, status int, body dto.APIResponse) error {
	return c.JSON(status, body)
}

func Success(c echo.Context, status int, data interface{}, message string) error {
	return JSON(c, status, dto.APIResponse{Success: true, Data: data, Message: message})
}

func Failure(c echo.Context, status int, details string) error {
	return JSON(c, status, dto.APIResponse{Success: false, Error: &dto.ErrorBody{Code: status, Details: details}})
}
