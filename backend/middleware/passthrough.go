package middleware

import "github.com/labstack/echo/v4"

// Passthrough is a no-op middleware used when a middleware is disabled in config.
func Passthrough() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return next
	}
}
