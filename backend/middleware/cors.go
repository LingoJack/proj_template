package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/lingojack/proj_template/config"
)

func CORS(cfg *config.Config) echo.MiddlewareFunc {
	c := cfg.Middleware.CORS
	if !c.Enabled {
		return Passthrough()
	}

	methods := c.AllowedMethods
	if len(methods) == 0 {
		methods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	}

	return echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins:     c.AllowedOrigins,
		AllowMethods:     methods,
		AllowCredentials: c.AllowCredentials,
	})
}
