package service

import (
	"fmt"
	"testing"

	"github.com/lingojack/proj_template/model"
)

// mockPostRepository is a mock implementation of repository.PostRepository
type mockPostRepository struct {
	posts map[uint]*model.Post
	nextID uint
}

func newMockPostRepository() *mockPostRepository {
	return &mockPostRepository{
		posts:  make(map[uint]*model.Post),
		nextID: 1,
	}
}

func (m *mockPostRepository) Create(post *model.Post) error {
	post.ID = m.nextID
	m.nextID++
	m.posts[post.ID] = post
	return nil
}

func (m *mockPostRepository) GetByID(id uint) (*model.Post, error) {
	p, ok := m.posts[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return p, nil
}

func (m *mockPostRepository) List(offset, limit int) ([]model.Post, int64, error) {
	var all []model.Post
	for _, p := range m.posts {
		all = append(all, *p)
	}
	total := int64(len(all))
	start := offset
	if start > len(all) {
		start = len(all)
	}
	end := start + limit
	if end > len(all) {
		end = len(all)
	}
	return all[start:end], total, nil
}

func (m *mockPostRepository) Update(post *model.Post) error {
	m.posts[post.ID] = post
	return nil
}

func (m *mockPostRepository) Delete(id uint) error {
	delete(m.posts, id)
	return nil
}

func TestPostService_CreatePost(t *testing.T) {
	repo := newMockPostRepository()
	svc := NewPostService(repo)

	post, err := svc.CreatePost("Title", "Content")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if post.ID != 1 {
		t.Errorf("expected ID 1, got %d", post.ID)
	}
}

func TestPostService_GetPost(t *testing.T) {
	repo := newMockPostRepository()
	svc := NewPostService(repo)

	created, _ := svc.CreatePost("Title", "Content")
	got, err := svc.GetPost(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Title != "Title" {
		t.Errorf("expected title 'Title', got '%s'", got.Title)
	}
}

func TestPostService_ListPosts(t *testing.T) {
	repo := newMockPostRepository()
	svc := NewPostService(repo)

	for i := 0; i < 5; i++ {
		svc.CreatePost("Title", "Content")
	}

	posts, total, err := svc.ListPosts(1, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 5 {
		t.Errorf("expected total 5, got %d", total)
	}
	if len(posts) != 3 {
		t.Errorf("expected 3 posts, got %d", len(posts))
	}
}

func TestPostService_DeletePost(t *testing.T) {
	repo := newMockPostRepository()
	svc := NewPostService(repo)

	created, _ := svc.CreatePost("Title", "Content")
	if err := svc.DeletePost(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := svc.GetPost(created.ID)
	if err == nil {
		t.Error("expected error after delete, got nil")
	}
}
