package test

import (
	"testing"

	"github.com/SirodgevAlex/comments-service/repository"
)

func TestInMemoryRepository_GetAllPosts(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	testGetAllPosts(t, repo)
}

func TestInMemoryRepository_SavePost(t *testing.T) {
    repo := repository.NewInMemoryRepository()
    testSavePost(t, repo)
}

func TestInMemoryRepository_GetPost(t *testing.T) {
    repo := repository.NewInMemoryRepository()
    testGetPost(t, repo)
}

func TestInMemoryRepository_UpdatePostCommentsSettings(t *testing.T) {
    repo := repository.NewInMemoryRepository()
    testUpdatePostCommentsSettings(t, repo)
}

func TestInMemoryRepository_SaveComment(t *testing.T) {
    repo := repository.NewInMemoryRepository()
    testSaveComment(t, repo)
}

func TestInMemoryRepository_GetCommentsByID(t *testing.T) {
    repo := repository.NewInMemoryRepository()
    testGetCommentsByID(t, repo)
}

func TestInMemoryRepository_GetComment(t *testing.T) {
    repo := repository.NewInMemoryRepository()
    testGetComment(t, repo)
}