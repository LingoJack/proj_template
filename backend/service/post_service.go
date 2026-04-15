package service

import (
	"github.com/lingojack/proj_template/model"
	"github.com/lingojack/proj_template/repository"
)

type PostService interface {
	CreatePost(title, content string) (*model.Post, error)
	GetPost(id uint) (*model.Post, error)
	ListPosts(page, pageSize int) ([]model.Post, int64, error)
	UpdatePost(id uint, title, content string) (*model.Post, error)
	DeletePost(id uint) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(title, content string) (*model.Post, error) {
	post := &model.Post{
		Title:   title,
		Content: content,
	}
	if err := s.repo.Create(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) GetPost(id uint) (*model.Post, error) {
	return s.repo.GetByID(id)
}

func (s *postService) ListPosts(page, pageSize int) ([]model.Post, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return s.repo.List(offset, pageSize)
}

func (s *postService) UpdatePost(id uint, title, content string) (*model.Post, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	post.Title = title
	post.Content = content
	if err := s.repo.Update(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) DeletePost(id uint) error {
	return s.repo.Delete(id)
}
