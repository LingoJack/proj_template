//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/lingojack/proj_template/config"
	"github.com/lingojack/proj_template/controller"
	"github.com/lingojack/proj_template/dao"
	"github.com/lingojack/proj_template/pkg/database"
	applogger "github.com/lingojack/proj_template/pkg/logger"
	"github.com/lingojack/proj_template/router"
	"github.com/lingojack/proj_template/service"
)

func initEcho(cfgPath string) (*echo.Echo, func(), error) {
	wire.Build(
		// Config
		config.Load,

		// Logger
		applogger.New,

		// Database
		database.New,

		// DAO layer
		dao.NewTUserDao,

		// Service layer
		service.NewUserService,

		// Controller layer
		controller.NewHealthController,
		controller.NewUserController,

		// Assemble Controllers struct
		provideControllers,

		// Echo + routes
		provideEcho,
	)
	return nil, nil, nil
}

func provideControllers(
	health *controller.HealthController,
	user *controller.UserController,
) *router.Controllers {
	return &router.Controllers{
		Health: health,
		User:   user,
	}
}

func provideEcho(cfg *config.Config, ctrl *router.Controllers) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	router.Register(e, cfg, ctrl)
	return e
}
