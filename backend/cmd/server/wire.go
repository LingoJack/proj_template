//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/lingojack/proj_template/config"
	"github.com/lingojack/proj_template/controller"
	"github.com/lingojack/proj_template/pkg/database"
	"github.com/lingojack/proj_template/pkg/logger"
	"github.com/lingojack/proj_template/pkg/validator"
	"github.com/lingojack/proj_template/repository"
	"github.com/lingojack/proj_template/router"
	"github.com/lingojack/proj_template/service"
)

func initEcho(cfgPath string) (*echo.Echo, func(), error) {
	wire.Build(
		config.Load,
		logger.New,
		database.New,
		validator.New,
		repository.NewPostRepository,
		service.NewPostService,
		controller.NewHealthController,
		controller.NewPostController,
		router.NewControllers,
		router.NewEcho,
	)
	return nil, nil, nil
}
