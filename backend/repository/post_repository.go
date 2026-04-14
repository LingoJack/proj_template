package repository

import (
	"github.com/lingojack/proj_template/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	GetByID(id uint) (*model.Post, error)
	List(offset, limit int) ([]model.Post, int64, error)
	Update(post *model.Post) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) List(offset, limit int) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	if err := r.db.Model(&model.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(limit).Order("id DESC").Find(&posts).Error; err != nil {
		return nil, 0, err
	}
	return posts, total, nil
}

func (r *postRepository) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&model.Post{}, id).Error
}
