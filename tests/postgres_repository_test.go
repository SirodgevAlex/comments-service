package test

import (
    "testing"
	"log"

	"github.com/SirodgevAlex/comments-service/repository"
	_ "github.com/lib/pq"
)


func TestPostgresRepository_GetAllPosts(t *testing.T) {
	repo, err := repository.NewPostgresRepository("postgres://postgres:1234@localhost:5432/test-comments-system?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to initialize Postgres repository: %v", err)
	}
	testGetAllPosts(t, repo)
}

func TestPostgresRepository_SavePost(t *testing.T) {
    repo, err := repository.NewPostgresRepository("postgres://postgres:1234@localhost:5432/test-comments-system?sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to create PostgreSQL repository: %v", err)
    }
    testSavePost(t, repo)
}

func TestPostgresRepository_GetPost(t *testing.T) {
    repo, err := repository.NewPostgresRepository("postgres://postgres:1234@localhost:5432/test-comments-system?sslmode=disable")
    if err != nil {
        t.Fatalf("Failed to initialize Postgres repository: %v", err)
    }
    testGetPost(t, repo)
}

func TestPostgresRepository_UpdatePostCommentsSettings(t *testing.T) {
    repo, err := repository.NewPostgresRepository("postgres://postgres:1234@localhost:5432/test-comments-system?sslmode=disable")
    if err != nil {
        t.Fatalf("Failed to initialize Postgres repository: %v", err)
    }
    testUpdatePostCommentsSettings(t, repo)
}

func TestPostgresRepository_SaveComment(t *testing.T) {
    repo, err := repository.NewPostgresRepository("postgres://postgres:1234@localhost:5432/test-comments-system?sslmode=disable")
    if err != nil {
        t.Fatalf("Failed to initialize Postgres repository: %v", err)
    }
    testSaveComment(t, repo)
}

func TestPostgresRepository_GetCommentsByID(t *testing.T) {
    repo, err := repository.NewPostgresRepository("postgres://postgres:1234@localhost:5432/test-comments-system?sslmode=disable")
    if err != nil {
        t.Fatalf("Failed to initialize Postgres repository: %v", err)
    }
    testGetCommentsByID(t, repo)
}

func TestPostgresRepository_GetComment(t *testing.T) {
    repo, err := repository.NewPostgresRepository("postgres://postgres:1234@localhost:5432/test-comments-system?sslmode=disable")
    if err != nil {
        t.Fatalf("Failed to initialize Postgres repository: %v", err)
    }
    testGetComment(t, repo)
}