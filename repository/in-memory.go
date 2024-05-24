package repository

import (
	"context"
	"sync"
	"fmt"
	"github.com/SirodgevAlex/comments-service/graph/model"
)

type InMemoryRepository struct {
	mu        sync.Mutex
	Posts     map[int]*model.Post
	Comments  map[int][]*model.Comment	
	Comment   map[int]*model.Comment
	IDCounter int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		Posts:    make(map[int]*model.Post),
		Comments: make(map[int][]*model.Comment),
		Comment:  make(map[int]*model.Comment),
		IDCounter: 0,
	}
}

func (r *InMemoryRepository) SavePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.IDCounter++

	newPost := &model.Post{
		ID:               r.IDCounter,
		Title:            post.Title,
		Content:          post.Content,
		AuthorID:         post.AuthorID,
		CommentsDisabled: post.CommentsDisabled,
		CreatedAt:        post.CreatedAt,
	}

	r.Posts[r.IDCounter] = newPost

	return newPost, nil
}

func (r *InMemoryRepository) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	posts := make([]*model.Post, 0, len(r.Posts))

	for _, post := range r.Posts {
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *InMemoryRepository) SaveComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.IDCounter++

	newComment := &model.Comment{
		ID:        r.IDCounter,
		PostID:    comment.PostID,
		ParentID:  comment.ParentID,
		AuthorID:  comment.AuthorID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}

	if comment.ParentID != nil {
		r.Comments[*comment.ParentID] = append(r.Comments[*comment.ParentID], newComment)
	} else {
		r.Comments[comment.PostID] = append(r.Comments[comment.PostID], newComment)
	}
	r.Comment[r.IDCounter] = newComment

	return newComment, nil
}

func (r *InMemoryRepository) GetCommentsByID(ctx context.Context, postID int, after *int, first *int) ([]*model.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	comments := r.Comments[postID]

	var allComments []*model.Comment
	for _, comment := range comments {
		allComments = append(allComments, r.getCommentsRecursively(comment)...)
	}

	if after != nil && first != nil {
		allComments = allComments[*after:*after + *first]
	} else if after != nil {
		allComments = allComments[*after:]
	} else if first != nil {
		allComments = allComments[:*first]
	}

	return allComments, nil
}

func (r *InMemoryRepository) getCommentsRecursively(comment *model.Comment) []*model.Comment {
	var allComments []*model.Comment

	allComments = append(allComments, comment)

	for _, childComment := range r.Comments[comment.ID] {
		allComments = append(allComments, r.getCommentsRecursively(childComment)...)
	}

	return allComments
}

func (r *InMemoryRepository) GetPost(ctx context.Context, postID int) (*model.Post, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    if post, ok := r.Posts[postID]; ok {
        return post, nil
    }

    return nil, fmt.Errorf("post with ID %d not found", postID)
}

func (r *InMemoryRepository) UpdatePostCommentsSettings(ctx context.Context, postID int, commentsDisabled bool) (*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	post, ok := r.Posts[postID]
	if !ok {
		return nil, fmt.Errorf("post with ID %d not found", postID)
	}

	post.CommentsDisabled = commentsDisabled

	return post, nil
}

func (r *InMemoryRepository) GetComment(ctx context.Context, commentID int) (*model.Comment, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    if comment, ok := r.Comment[commentID]; ok {
        return comment, nil
    }

    return nil, fmt.Errorf("comment with ID %d not found", commentID)
}