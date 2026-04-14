package middleware

import "github.com/labstack/echo/v4"

// passthrough is a no-op middleware used when a middleware is disabled in config.
func passthrough() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return next
	}
}
