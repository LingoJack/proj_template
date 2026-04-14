package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/proj_template/pkg/response"
	"github.com/lingojack/proj_template/pkg/validator"
	"github.com/lingojack/proj_template/service"
)

type UserController struct {
	userSvc *service.UserService
}

func NewUserController(userSvc *service.UserService) *UserController {
	return &UserController{userSvc: userSvc}
}

type createUserReq struct {
	Username string `json:"username" validate:"required,min=2,max=64"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (uc *UserController) Create(c echo.Context) error {
	var req createUserReq
	if err := validator.BindAndValidate(c, &req); err != nil {
		return response.BadRequest(c, err.Error())
	}
	user, err := uc.userSvc.Create(service.CreateUserInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return response.InternalError(c)
	}
	return response.Created(c, user)
}

func (uc *UserController) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid id")
	}
	user, err := uc.userSvc.GetByID(uint(id))
	if err != nil {
		return response.NotFound(c, "user not found")
	}
	return response.OK(c, user)
}

func (uc *UserController) List(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	result, err := uc.userSvc.List(page, pageSize)
	if err != nil {
		return response.InternalError(c)
	}
	return response.OK(c, map[string]any{
		"items": result.Items,
		"total": result.Total,
	})
}

type updateUserReq struct {
	Email string `json:"email" validate:"omitempty,email"`
}

func (uc *UserController) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid id")
	}
	var req updateUserReq
	if err := validator.BindAndValidate(c, &req); err != nil {
		return response.BadRequest(c, err.Error())
	}
	user, err := uc.userSvc.Update(uint(id), map[string]any{"email": req.Email})
	if err != nil {
		return response.NotFound(c, "user not found")
	}
	return response.OK(c, user)
}

func (uc *UserController) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid id")
	}
	if err := uc.userSvc.Delete(uint(id)); err != nil {
		return response.NotFound(c, "user not found")
	}
	return c.NoContent(http.StatusNoContent)
}
