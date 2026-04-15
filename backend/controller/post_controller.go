package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/proj_template/model"
	"github.com/lingojack/proj_template/pkg/response"
	"github.com/lingojack/proj_template/pkg/validator"
	"github.com/lingojack/proj_template/service"
)

type PostController struct {
	postSvc service.PostService
}

func NewPostController(postSvc service.PostService) *PostController {
	return &PostController{postSvc: postSvc}
}

// ---------- Request / Response DTOs ----------

type CreatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type PostResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListPostsResponse struct {
	Items    []PostResponse `json:"items"`
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
}

// ---------- Handlers ----------

func (pc *PostController) Create(c echo.Context) error {
	var req CreatePostRequest
	if err := validator.BindAndValidate(c, &req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	post, err := pc.postSvc.CreatePost(req.Title, req.Content)
	if err != nil {
		return response.InternalError(c)
	}

	return c.JSON(http.StatusCreated, response.Response{
		Code:    0,
		Message: "created",
		Data:    toPostResponse(post),
	})
}

func (pc *PostController) Get(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid id")
	}

	post, err := pc.postSvc.GetPost(uint(id))
	if err != nil {
		return response.NotFound(c, "post not found")
	}

	return response.OK(c, toPostResponse(post))
}

func (pc *PostController) List(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	posts, total, err := pc.postSvc.ListPosts(page, pageSize)
	if err != nil {
		return response.InternalError(c)
	}

	items := make([]PostResponse, 0, len(posts))
	for i := range posts {
		items = append(items, toPostResponse(&posts[i]))
	}

	return response.OK(c, ListPostsResponse{
		Items:    items,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

func (pc *PostController) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid id")
	}

	var req UpdatePostRequest
	if err := validator.BindAndValidate(c, &req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	post, err := pc.postSvc.UpdatePost(uint(id), req.Title, req.Content)
	if err != nil {
		return response.NotFound(c, "post not found")
	}

	return response.OK(c, toPostResponse(post))
}

func (pc *PostController) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid id")
	}

	if err := pc.postSvc.DeletePost(uint(id)); err != nil {
		return response.NotFound(c, "post not found")
	}

	return response.OK(c, nil)
}

// ---------- Helper ----------

func toPostResponse(p *model.Post) PostResponse {
	return PostResponse{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
	}
}
