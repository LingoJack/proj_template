package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/lingojack/proj_template/config"
)

func RequestID(cfg *config.Config) echo.MiddlewareFunc {
	if !cfg.Middleware.RequestID.Enabled {
		return passthrough()
	}
	return echomw.RequestID()
}
