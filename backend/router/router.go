package router

import (
	"github.com/labstack/echo/v4"

	"github.com/lingojack/proj_template/config"
	"github.com/lingojack/proj_template/controller"
	mw "github.com/lingojack/proj_template/middleware"
)

type Controllers struct {
	Health *controller.HealthController
	User   *controller.UserController
}

func Register(e *echo.Echo, cfg *config.Config, ctrl *Controllers) {
	// 全局中间件 — 所有请求都经过
	e.Use(mw.Recover(cfg))
	e.Use(mw.RequestID(cfg))
	e.Use(mw.Logger())

	// 公开路由 — 无需鉴权
	open := e.Group(cfg.API.Prefix)
	open.Use(mw.CORS(cfg))
	open.GET("/health", ctrl.Health.Check)

	// 需要鉴权的路由
	private := e.Group(cfg.API.Prefix)
	private.Use(mw.CORS(cfg))
	private.Use(mw.Auth(cfg))

	// User 路由
	private.GET("/users", ctrl.User.List)
	private.GET("/users/:id", ctrl.User.GetByID)
	private.POST("/users", ctrl.User.Create)
	private.PUT("/users/:id", ctrl.User.Update)
	private.DELETE("/users/:id", ctrl.User.Delete)
}
