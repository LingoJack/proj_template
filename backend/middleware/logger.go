package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

// Logger returns Echo's built-in request logger middleware.
func Logger() echo.MiddlewareFunc {
	return echomw.Logger()
}
