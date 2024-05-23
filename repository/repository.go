package repository

import (
	"sync"
	"context"
	"github.com/SirodgevAlex/comments-service/graph/model"
)

type Repository interface {
	SavePost(ctx context.Context, post *model.Post) (*model.Post, error)
	GetAllPosts(ctx context.Context)  ([]*model.Post, error)
}

type InMemoryRepository struct {
	Posts map[int]*model.Post
	mu    sync.Mutex
	postIDCounter int
}

func (r *InMemoryRepository) SavePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	// Защищаем доступ к общим данным
	r.mu.Lock()
	defer r.mu.Unlock()

	// Увеличиваем счетчик ID для поста
	r.postIDCounter++

	// Создаем копию поста с установленным ID
	newPost := &model.Post{
		ID:               r.postIDCounter,
		Title:            post.Title,
		Content:          post.Content,
		AuthorID:         post.AuthorID,
		CommentsDisabled: post.CommentsDisabled,
		CreatedAt:        post.CreatedAt,
	}

	// Сохраняем пост в словаре по его ID
	r.Posts[r.postIDCounter] = newPost

	return newPost, nil
}

func (r *InMemoryRepository) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
    // Защищаем доступ к общим данным
    r.mu.Lock()
    defer r.mu.Unlock()

    // Создаем слайс для хранения всех постов
    posts := make([]*model.Post, 0, len(r.Posts))

    // Проходим по всем постам в словаре и добавляем их в слайс
    for _, post := range r.Posts {
        posts = append(posts, post)
    }

    return posts, nil
}
