package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"time"

	"github.com/SirodgevAlex/comments-service/graph/model"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, authorID int) (*model.Post, error) {
	// Create a new Post instance
	post := &model.Post{
		Title:            title,
		Content:          content,
		AuthorID:         authorID,
		CommentsDisabled: false,
		CreatedAt:        time.Now(),
	}

	// Save the post using repository
	savedPost, err := r.Repo.SavePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return savedPost, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, postID int, parentID *int, authorID int, content string) (*model.Comment, error) {
	// Проверяем длину текста комментария
	if len(content) > maxCommentLength {
		return nil, fmt.Errorf("comment length exceeds the maximum allowed length of %d characters", maxCommentLength)
	}

	// Проверяем, существует ли родительский комментарий
	// if parentID != nil {
	// 	parentComment, err := r.Repo.GetCommentsByID(ctx, *parentID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if parentComment == nil {
	// 		return nil, fmt.Errorf("parent comment not found")
	// 	}
	// }

	// Создаем новый комментарий
	newComment := &model.Comment{
		PostID:    postID,
		ParentID:  parentID,
		AuthorID:  authorID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	// Сохраняем новый комментарий
	savedComment, err := r.Repo.SaveComment(ctx, newComment)
	if err != nil {
		return nil, err
	}

	return savedComment, nil
}

// UpdatePostCommentsSettings is the resolver for the updatePostCommentsSettings field.
func (r *mutationResolver) UpdatePostCommentsSettings(ctx context.Context, postID int, authorID int, commentsDisabled bool) (*model.Post, error) {
    // Получаем пост по его ID
    post, err := r.Repo.GetPost(ctx, postID)
    if err != nil {
        return nil, err
    }

    // Проверяем, соответствует ли автор ID автору поста
    if post.AuthorID != authorID {
        return nil, fmt.Errorf("user %d is not the author of post %d", authorID, postID)
    }

    // Обновляем настройки комментариев для поста
    updatedPost, err := r.Repo.UpdatePostCommentsSettings(ctx, postID, commentsDisabled)
    if err != nil {
        return nil, err
    }

    return updatedPost, nil
}


// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.Repo.GetAllPosts(ctx)
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID int, after *int, first *int) ([]*model.Comment, error) {
	return r.Repo.GetCommentsByID(ctx, postID, after, first)
}

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID int) (<-chan *model.Comment, error) {
	panic(fmt.Errorf("not implemented: CommentAdded - commentAdded"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
const maxCommentLength = 2000
