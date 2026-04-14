package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/lingojack/proj_template/config"
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Auth returns a JWT validation middleware.
// When disabled in config, it passes through all requests.
func Auth(cfg *config.Config) echo.MiddlewareFunc {
	if !cfg.Middleware.Auth.Enabled {
		return passthrough()
	}
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cfg.Middleware.Auth.JWTSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &JWTClaims{}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return errors.New("unauthorized")
		},
	})
}

// ClaimsFromContext extracts JWTClaims from Echo context after Auth middleware runs.
func ClaimsFromContext(c echo.Context) *JWTClaims {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok || token == nil {
		return nil
	}
	claims, _ := token.Claims.(*JWTClaims)
	return claims
}
