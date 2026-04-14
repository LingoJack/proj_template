package service

import (
	"context"
	"errors"

	"github.com/lingojack/proj_template/dao"
	"github.com/lingojack/proj_template/model/entity"
	"github.com/lingojack/proj_template/model/query"
)

type UserService struct {
	userDAO *dao.TUserDao
}

func NewUserService(userDAO *dao.TUserDao) *UserService {
	return &UserService{userDAO: userDAO}
}

type CreateUserInput struct {
	Username string
	Email    string
	Password string
}

func (s *UserService) Create(ctx context.Context, input CreateUserInput) (*entity.TUser, error) {
	if input.Username == "" || input.Email == "" || input.Password == "" {
		return nil, errors.New("username, email and password are required")
	}
	user := &entity.TUser{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
	if err := s.userDAO.Insert(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByID(ctx context.Context, id uint64) (*entity.TUser, error) {
	return s.userDAO.SelectById(ctx, id)
}

type ListResult struct {
	Items []*entity.TUser
	Total int64
}

func (s *UserService) List(ctx context.Context, page, pageSize int) (*ListResult, error) {
	if pageSize <= 0 {
		pageSize = 20
	}
	if page <= 0 {
		page = 1
	}
	dto := &query.TUserDto{
		PageOffset: (page - 1) * pageSize,
		PageSize:   pageSize,
	}
	users, err := s.userDAO.SelectList(ctx, dto)
	if err != nil {
		return nil, err
	}
	total, err := s.userDAO.SelectCount(ctx, dto)
	if err != nil {
		return nil, err
	}
	return &ListResult{Items: users, Total: total}, nil
}

func (s *UserService) Update(ctx context.Context, id uint64, updates map[string]any) (*entity.TUser, error) {
	if err := s.userDAO.UpdateByIdWithMap(ctx, id, updates); err != nil {
		return nil, err
	}
	return s.userDAO.SelectById(ctx, id)
}

func (s *UserService) Delete(ctx context.Context, id uint64) error {
	return s.userDAO.DeleteById(ctx, id)
}
