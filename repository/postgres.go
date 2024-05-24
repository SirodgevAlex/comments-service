package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/SirodgevAlex/comments-service/db"
	"github.com/SirodgevAlex/comments-service/graph/model"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(connStr string) (*PostgresRepository, error) {
    db, err := db.ConnectPostgresDB(connStr)
	if err != nil {
        return nil, err
    }
    return &PostgresRepository{DB: db}, nil
}

func ClosePostgresRepository(db *sql.DB) {
    if db != nil {
        db.Close()
        log.Println("Disconnected from PostgreSQL database")
    }
}

func (r *PostgresRepository) SavePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	query := `INSERT INTO posts (title, content, author_id, comments_disabled, created_at)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, post.Title, post.Content, post.AuthorID, post.CommentsDisabled, post.CreatedAt).Scan(&post.ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostgresRepository) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	query := `SELECT id, title, content, author_id, comments_disabled, created_at FROM posts`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*model.Post
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CommentsDisabled, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostgresRepository) SaveComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	query := `INSERT INTO comments (post_id, parent_id, author_id, content, created_at)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRowContext(ctx, query, comment.PostID, comment.ParentID, comment.AuthorID, comment.Content, comment.CreatedAt).Scan(&comment.ID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *PostgresRepository) GetPost(ctx context.Context, postID int) (*model.Post, error) {
	query := `SELECT id, title, content, author_id, comments_disabled, created_at FROM posts WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, postID)

	var post model.Post
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CommentsDisabled, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostgresRepository) UpdatePostCommentsSettings(ctx context.Context, postID int, commentsDisabled bool) (*model.Post, error) {
	query := `UPDATE posts SET comments_disabled = $1 WHERE id = $2 RETURNING id, title, content, author_id, comments_disabled, created_at`
	row := r.DB.QueryRowContext(ctx, query, commentsDisabled, postID)

	var post model.Post
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CommentsDisabled, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostgresRepository) GetCommentsByID(ctx context.Context, postID int, after *int, first *int) ([]*model.Comment, error) {
    query := `
        WITH RECURSIVE comment_tree AS (
            SELECT id, post_id, parent_id, author_id, content, created_at
            FROM comments ct
            WHERE post_id = $1 AND parent_id IS NULL
            
            UNION ALL
            
            SELECT c.id, c.post_id, c.parent_id, c.author_id, c.content, c.created_at
            FROM comments c
            INNER JOIN comment_tree ct ON c.parent_id = ct.id
        )
        SELECT id, post_id, parent_id, author_id, content, created_at
        FROM comment_tree
        WHERE post_id = $1
    `

    args := []interface{}{postID}

    if first != nil && after != nil {
        query += " LIMIT $2 OFFSET $3"
        args = append(args, *first, *after)
    } else if first != nil {
        query += " LIMIT $2"
        args = append(args, *first)
    } else if after != nil {
        query += " OFFSET $2"
        args = append(args, *after)
    }

    rows, err := r.DB.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var comments []*model.Comment
    for rows.Next() {
        var comment model.Comment
        err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.AuthorID, &comment.Content, &comment.CreatedAt)
        if err != nil {
            return nil, err
        }
        comments = append(comments, &comment)
    }
    return comments, nil
}

func (r *PostgresRepository) GetComment(ctx context.Context, commentID int) (*model.Comment, error) {
	query := `SELECT id, post_id, parent_id, author_id, content, created_at FROM comments WHERE id = $1`
	row := r.DB.QueryRowContext(ctx, query, commentID)

	var comment model.Comment
	err := row.Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.AuthorID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &comment, nil
}