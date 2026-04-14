package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	v *validator.Validate
}

func New() *CustomValidator {
	return &CustomValidator{v: validator.New()}
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.v.Struct(i)
}

// BindAndValidate binds request body and validates it.
func BindAndValidate(c echo.Context, dst any) error {
	if err := c.Bind(dst); err != nil {
		return err
	}
	return c.Validate(dst)
}
