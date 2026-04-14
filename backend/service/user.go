package service

import (
	"errors"

	"github.com/lingojack/proj_template/dao"
	"github.com/lingojack/proj_template/model"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}
}

type CreateUserInput struct {
	Username string
	Email    string
	Password string
}

func (s *UserService) Create(input CreateUserInput) (*model.User, error) {
	if input.Username == "" || input.Email == "" || input.Password == "" {
		return nil, errors.New("username, email and password are required")
	}
	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password, // hash in production
	}
	if err := s.userDAO.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	return s.userDAO.GetByID(id)
}

func (s *UserService) List(page, pageSize int) (*dao.ListResult, error) {
	return s.userDAO.List(dao.ListParams{Page: page, PageSize: pageSize})
}

func (s *UserService) Update(id uint, updates map[string]any) (*model.User, error) {
	user, err := s.userDAO.GetByID(id)
	if err != nil {
		return nil, err
	}
	if v, ok := updates["email"].(string); ok {
		user.Email = v
	}
	if err := s.userDAO.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Delete(id uint) error {
	return s.userDAO.Delete(id)
}
