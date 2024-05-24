package repository

import (
	"context"
	"github.com/SirodgevAlex/comments-service/graph/model"
)

type Repository interface {
	SavePost(ctx context.Context, post *model.Post) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]*model.Post, error)
	SaveComment(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	GetCommentsByID(ctx context.Context, id int, after *int, first *int) ([]*model.Comment, error)
	GetPost(ctx context.Context, postID int) (*model.Post, error)
	UpdatePostCommentsSettings(ctx context.Context, postID int, commentsDisabled bool) (*model.Post, error)
	GetComment(ctx context.Context, commentID int) (*model.Comment, error)
}