package repository

import (
	"testing"

	"github.com/lingojack/proj_template/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&model.Post{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestPostRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)

	post := &model.Post{Title: "Test", Content: "Content"}
	if err := repo.Create(post); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if post.ID == 0 {
		t.Error("expected post ID to be set after create")
	}
}

func TestPostRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)

	post := &model.Post{Title: "Test", Content: "Content"}
	_ = repo.Create(post)

	got, err := repo.GetByID(post.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Title != "Test" {
		t.Errorf("expected title 'Test', got '%s'", got.Title)
	}
}

func TestPostRepository_List(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)

	for i := 0; i < 15; i++ {
		_ = repo.Create(&model.Post{Title: "Post", Content: "Content"})
	}

	posts, total, err := repo.List(0, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 15 {
		t.Errorf("expected total 15, got %d", total)
	}
	if len(posts) != 10 {
		t.Errorf("expected 10 posts, got %d", len(posts))
	}
}

func TestPostRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)

	post := &model.Post{Title: "Old", Content: "Old"}
	_ = repo.Create(post)

	post.Title = "New"
	if err := repo.Update(post); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, _ := repo.GetByID(post.ID)
	if got.Title != "New" {
		t.Errorf("expected title 'New', got '%s'", got.Title)
	}
}

func TestPostRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewPostRepository(db)

	post := &model.Post{Title: "Delete", Content: "Me"}
	_ = repo.Create(post)

	if err := repo.Delete(post.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := repo.GetByID(post.ID)
	if err == nil {
		t.Error("expected error after delete, got nil")
	}
}
