package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/lingojack/proj_template/pkg/response"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Check(c echo.Context) error {
	return response.OK(c, map[string]string{
		"status": "ok",
	})
}
