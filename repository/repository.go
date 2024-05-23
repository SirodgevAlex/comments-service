package repository

import (
	"context"
	"sync"
	"fmt"
	"github.com/SirodgevAlex/comments-service/graph/model"
)

type Repository interface {
	SavePost(ctx context.Context, post *model.Post) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]*model.Post, error)
	SaveComment(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	GetCommentsByID(ctx context.Context, id int, after *int, first *int) ([]*model.Comment, error)
	GetPost(ctx context.Context, postID int) (*model.Post, error)
	UpdatePostCommentsSettings(ctx context.Context, postID int, commentsDisabled bool) (*model.Post, error)
}

type InMemoryRepository struct {
	mu        sync.Mutex
	Posts     map[int]*model.Post
	Comments  map[int][]*model.Comment
	IDCounter int
}

func (r *InMemoryRepository) SavePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	// Защищаем доступ к общим данным
	r.mu.Lock()
	defer r.mu.Unlock()

	// Увеличиваем счетчик ID для поста
	r.IDCounter++

	// Создаем копию поста с установленным ID
	newPost := &model.Post{
		ID:               r.IDCounter,
		Title:            post.Title,
		Content:          post.Content,
		AuthorID:         post.AuthorID,
		CommentsDisabled: post.CommentsDisabled,
		CreatedAt:        post.CreatedAt,
	}

	// Сохраняем пост в словаре по его ID
	r.Posts[r.IDCounter] = newPost

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

func (r *InMemoryRepository) SaveComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	// Защищаем доступ к общим данным
	r.mu.Lock()
	defer r.mu.Unlock()

	// Увеличиваем счетчик ID для комментария
	r.IDCounter++

	// Создаем копию комментария с установленным ID
	newComment := &model.Comment{
		ID:        r.IDCounter,
		PostID:    comment.PostID,
		ParentID:  comment.ParentID,
		AuthorID:  comment.AuthorID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}

	// Сохраняем комментарий в словаре по его ID
	if comment.ParentID != nil {
		r.Comments[*comment.ParentID] = append(r.Comments[*comment.ParentID], newComment)
	} else {
		r.Comments[comment.PostID] = append(r.Comments[comment.PostID], newComment)
	}

	return newComment, nil
}

func (r *InMemoryRepository) GetCommentsByID(ctx context.Context, postID int, after *int, first *int) ([]*model.Comment, error) {
	// Защищаем доступ к общим данным
	r.mu.Lock()
	defer r.mu.Unlock()

	// Извлекаем слайс комментариев по заданному postID
	comments := r.Comments[postID]

	// Возвращаем комментарии для заданного postID
	return comments, nil
}


// func (r *InMemoryRepository) GetCommentsByID(ctx context.Context, postID int, after *int, first *int) ([]*model.Comment, error) {
// 	// Защищаем доступ к общим данным
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	// Извлекаем слайс комментариев по заданному postID
// 	comments := r.Comments[postID]

// 	// Рекурсивно собираем все комментарии
// 	var allComments []*model.Comment
// 	for _, comment := range comments {
// 		allComments = append(allComments, r.getCommentsRecursively(comment)...)
// 	}

// 	return allComments, nil
// }

// func (r *InMemoryRepository) getCommentsRecursively(comment *model.Comment) []*model.Comment {
// 	// Создаем слайс для хранения комментариев
// 	var allComments []*model.Comment

// 	// Добавляем текущий комментарий в слайс
// 	allComments = append(allComments, comment)

// 	// Рекурсивно добавляем комментарии в текущий комментарий
// 	for _, childComment := range r.Comments[comment.ID] {
// 		allComments = append(allComments, r.getCommentsRecursively(childComment)...)
// 	}

// 	return allComments
// }

func (r *InMemoryRepository) GetPost(ctx context.Context, postID int) (*model.Post, error) {
    // Защищаем доступ к общим данным
    r.mu.Lock()
    defer r.mu.Unlock()

    // Проверяем, существует ли пост с заданным ID
    if post, ok := r.Posts[postID]; ok {
        return post, nil
    }

    // Возвращаем ошибку, если пост не найден
    return nil, fmt.Errorf("post with ID %d not found", postID)
}

func (r *InMemoryRepository) UpdatePostCommentsSettings(ctx context.Context, postID int, commentsDisabled bool) (*model.Post, error) {
	// Защищаем доступ к общим данным
	r.mu.Lock()
	defer r.mu.Unlock()

	// Получаем пост по его ID
	post, ok := r.Posts[postID]
	if !ok {
		return nil, fmt.Errorf("post with ID %d not found", postID)
	}

	// Обновляем настройки комментариев для поста
	post.CommentsDisabled = commentsDisabled

	return post, nil
}