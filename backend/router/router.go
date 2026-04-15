package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/lingojack/proj_template/config"
	"github.com/lingojack/proj_template/controller"
	mw "github.com/lingojack/proj_template/middleware"
	"github.com/lingojack/proj_template/pkg/validator"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Controllers struct {
	Health *controller.HealthController
	Post   *controller.PostController
}

func NewControllers(
	health *controller.HealthController,
	post *controller.PostController,
) *Controllers {
	return &Controllers{
		Health: health,
		Post:   post,
	}
}

// NewEcho 创建 Echo 实例，注册中间件和路由，返回 (echo实例, cleanup函数, error)
func NewEcho(
	cfg *config.Config,
	log zerolog.Logger,
	db *gorm.DB,
	v *validator.CustomValidator,
	ctrl *Controllers,
) (*echo.Echo, func(), error) {
	e := echo.New()
	e.Validator = v
	e.HideBanner = true

	// 全局中间件
	e.Use(mw.Recover(cfg))
	e.Use(mw.RequestID(cfg))
	e.Use(mw.Logger())
	e.Use(mw.RateLimit(cfg))

	// 注册路由
	Register(e, cfg, ctrl)

	// cleanup 函数 — 关闭数据库等资源
	cleanup := func() {
		log.Info().Msg("running cleanup")
	}

	return e, cleanup, nil
}

func Register(e *echo.Echo, cfg *config.Config, ctrl *Controllers) {
	// 公开路由 — 无需鉴权
	open := e.Group(cfg.API.Prefix)
	open.Use(mw.CORS(cfg))
	open.GET("/health", ctrl.Health.Check)

	// 文章 CRUD（公开示例，生产环境可移至 private 组）
	open.GET("/posts", ctrl.Post.List)
	open.GET("/posts/:id", ctrl.Post.Get)
	open.POST("/posts", ctrl.Post.Create)
	open.PUT("/posts/:id", ctrl.Post.Update)
	open.DELETE("/posts/:id", ctrl.Post.Delete)

	// Swagger API 文档
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 需要鉴权的路由
	private := e.Group(cfg.API.Prefix)
	private.Use(mw.CORS(cfg))
	private.Use(mw.Auth(cfg))
}
