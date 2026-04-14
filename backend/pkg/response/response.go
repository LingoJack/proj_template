package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func OK(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

func Created(c echo.Context, data any) error {
	return c.JSON(http.StatusCreated, Response{
		Code:    0,
		Message: "created",
		Data:    data,
	})
}

func Fail(c echo.Context, httpStatus int, code int, msg string) error {
	return c.JSON(httpStatus, Response{
		Code:    code,
		Message: msg,
	})
}

func BadRequest(c echo.Context, msg string) error {
	return Fail(c, http.StatusBadRequest, 400, msg)
}

func Unauthorized(c echo.Context) error {
	return Fail(c, http.StatusUnauthorized, 401, "unauthorized")
}

func Forbidden(c echo.Context) error {
	return Fail(c, http.StatusForbidden, 403, "forbidden")
}

func NotFound(c echo.Context, msg string) error {
	return Fail(c, http.StatusNotFound, 404, msg)
}

func InternalError(c echo.Context) error {
	return Fail(c, http.StatusInternalServerError, 500, "internal server error")
}
