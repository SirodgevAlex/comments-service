package test

import (
    "context"
    "testing"
    "time"

    "github.com/SirodgevAlex/comments-service/graph/model"
    "github.com/SirodgevAlex/comments-service/repository"
)

func testGetAllPosts(t *testing.T, repo repository.Repository) {
	// Создание тестовых постов
	posts := []*model.Post{
		{
			Title:            "Test Post 1",
			Content:          "This is test post 1.",
			AuthorID:         1,
			CommentsDisabled: false,
			CreatedAt:        time.Now(),
		},
		{
			Title:            "Test Post 2",
			Content:          "This is test post 2.",
			AuthorID:         2,
			CommentsDisabled: false,
			CreatedAt:        time.Now(),
		},
	}

	// Сохранение тестовых постов
	for _, post := range posts {
		_, err := repo.SavePost(context.Background(), post)
		if err != nil {
			t.Fatalf("Failed to save post: %v", err)
		}
	}

	// Получение всех постов из базы данных
	allPosts, err := repo.GetAllPosts(context.Background())
	if err != nil {
		t.Fatalf("Failed to get all posts: %v", err)
	}

	// Проверка, что количество полученных постов совпадает с ожидаемым
	if len(allPosts) != len(posts) {
		t.Errorf("Unexpected number of posts: got %d, want %d", len(allPosts), len(posts))
	}

	// Проверка, что данные полученных постов соответствуют ожидаемым
	for i := range posts {
		if allPosts[i].Title != posts[i].Title || allPosts[i].Content != posts[i].Content || allPosts[i].AuthorID != posts[i].AuthorID || allPosts[i].CommentsDisabled != posts[i].CommentsDisabled {
			t.Errorf("Unexpected post data: got %v, want %v", allPosts[i], posts[i])
		}
	}
}

func testSavePost(t *testing.T, repo repository.Repository) {
    post := &model.Post{
        Title:            "Test Post",
        Content:          "This is a test post.",
        AuthorID:         1,
        CommentsDisabled: false,
        CreatedAt:        time.Now(),
    }

    savedPost, err := repo.SavePost(context.Background(), post)
    if err != nil {
        t.Fatalf("Failed to save post: %v", err)
    }

    if savedPost.ID == 0 {
        t.Error("Expected saved post to have a non-zero ID")
    }

    if savedPost.Title != post.Title || savedPost.Content != post.Content || savedPost.AuthorID != post.AuthorID || savedPost.CommentsDisabled != post.CommentsDisabled {
        t.Errorf("Unexpected post data: got %v, want %v", savedPost, post)
    }
}

func testGetPost(t *testing.T, repo repository.Repository) {
    // Создание тестового поста
    post := &model.Post{
        Title:            "Test Post",
        Content:          "This is a test post.",
        AuthorID:         1,
        CommentsDisabled: false,
        CreatedAt:        time.Now(),
    }

    // Сохранение тестового поста
    savedPost, err := repo.SavePost(context.Background(), post)
    if err != nil {
        t.Fatalf("Failed to save post: %v", err)
    }

    // Получение поста из базы данных по его ID
    foundPost, err := repo.GetPost(context.Background(), savedPost.ID)
    if err != nil {
        t.Fatalf("Failed to get post: %v", err)
    }

    // Проверка, что полученный пост соответствует ожидаемому
    if foundPost.ID != savedPost.ID || foundPost.Title != savedPost.Title || foundPost.Content != savedPost.Content || foundPost.AuthorID != savedPost.AuthorID || foundPost.CommentsDisabled != savedPost.CommentsDisabled {
        t.Errorf("Unexpected post data: got %v, want %v", foundPost, savedPost)
    }
}

func testUpdatePostCommentsSettings(t *testing.T, repo repository.Repository) {
    // Создание тестового поста
    post := &model.Post{
        Title:            "Test Post",
        Content:          "This is a test post.",
        AuthorID:         1,
        CommentsDisabled: false,
        CreatedAt:        time.Now(),
    }

    // Сохранение тестового поста
    savedPost, err := repo.SavePost(context.Background(), post)
    if err != nil {
        t.Fatalf("Failed to save post: %v", err)
    }

    // Обновление настроек комментариев поста
    updatedPost1, err := repo.UpdatePostCommentsSettings(context.Background(), savedPost.ID, true)
    if err != nil {
        t.Fatalf("Failed to update post: %v", err)
    }

    // Получение обновленного поста из базы данных
    updatedPost2, err := repo.GetPost(context.Background(), savedPost.ID)
    if err != nil {
        t.Fatalf("Failed to get updated post: %v", err)
    }

    // Проверка, что полученный пост соответствует ожидаемому
    if updatedPost1.CommentsDisabled != updatedPost2.CommentsDisabled {
        t.Errorf("Unexpected CommentsDisabled: got %t, want %t", updatedPost2.CommentsDisabled, true)
    }
}

func testSaveComment(t *testing.T, repo repository.Repository) {
    // Создание тестового поста
    post := &model.Post{
        Title:            "Test Post",
        Content:          "This is a test post.",
        AuthorID:         1,
        CommentsDisabled: false,
        CreatedAt:        time.Now(),
    }

    // Сохранение тестового поста
    savedPost, err := repo.SavePost(context.Background(), post)
    if err != nil {
        t.Fatalf("Failed to save post: %v", err)
    }

    // Создание тестового комментария
    comment := &model.Comment{
        Content:   "This is a test comment.",
        AuthorID:  1,
        PostID:    savedPost.ID, // Используем ID сохраненного поста
        CreatedAt: time.Now(),
    }

    // Сохранение комментария в базу данных
    savedComment, err := repo.SaveComment(context.Background(), comment)
    if err != nil {
        t.Fatalf("Failed to save comment: %v", err)
    }

    // Проверка, что сохраненный комментарий имеет ненулевой ID
    if savedComment.ID == 0 {
        t.Error("Expected saved comment to have a non-zero ID")
    }

    // Проверка, что данные сохраненного комментария соответствуют ожидаемым
    if savedComment.Content != comment.Content || savedComment.AuthorID != comment.AuthorID || savedComment.PostID != comment.PostID {
        t.Errorf("Unexpected comment data: got %v, want %v", savedComment, comment)
    }
}

func testGetCommentsByID(t *testing.T, repo repository.Repository) {
    // Создание тестового поста
    post := &model.Post{
        Title:            "Test Post",
        Content:          "This is a test post.",
        AuthorID:         1,
        CommentsDisabled: false,
        CreatedAt:        time.Now(),
    }

    // Сохранение тестового поста
    savedPost, err := repo.SavePost(context.Background(), post)
    if err != nil {
        t.Fatalf("Failed to save post: %v", err)
    }

    // Создание тестового комментария с использованием ID сохраненного поста
    comment := &model.Comment{
        Content:   "This is a test comment.",
        AuthorID:  1,
        PostID:    savedPost.ID,
        CreatedAt: time.Now(),
    }

    // Сохранение комментария в базу данных
    _, err = repo.SaveComment(context.Background(), comment)
    if err != nil {
        t.Fatalf("Failed to save comment: %v", err)
    }

    // Получение комментариев по ID поста
    comments, err := repo.GetCommentsByID(context.Background(), savedPost.ID, nil, nil)
    if err != nil {
        t.Fatalf("Failed to get comments by ID: %v", err)
    }

    // Проверка, что количество полученных комментариев больше 0
    if len(comments) == 0 {
        t.Error("Expected to get comments by ID, but got none")
    }
}

func testGetComment(t *testing.T, repo repository.Repository) {
    // Создание тестового поста
    post := &model.Post{
        Title:            "Test Post",
        Content:          "This is a test post.",
        AuthorID:         1,
        CommentsDisabled: false,
        CreatedAt:        time.Now(),
    }

    // Сохранение поста в базу данных
    savedPost, err := repo.SavePost(context.Background(), post)
    if err != nil {
        t.Fatalf("Failed to save post: %v", err)
    }

    // Создание тестового комментария
    comment := &model.Comment{
        Content:   "This is a test comment.",
        AuthorID:  1,
        PostID:    savedPost.ID, // Используем ID сохраненного поста
        CreatedAt: time.Now(),
    }

    // Сохранение комментария в базу данных
    savedComment, err := repo.SaveComment(context.Background(), comment)
    if err != nil {
        t.Fatalf("Failed to save comment: %v", err)
    }

    // Получение комментария по его ID
    foundComment, err := repo.GetComment(context.Background(), savedComment.ID)
    if err != nil {
        t.Fatalf("Failed to get comment: %v", err)
    }

    // Проверка, что полученный комментарий соответствует ожидаемому
    if foundComment.ID != savedComment.ID || foundComment.Content != savedComment.Content || foundComment.AuthorID != savedComment.AuthorID || foundComment.PostID != savedComment.PostID {
        t.Errorf("Unexpected comment data: got %v, want %v", foundComment, savedComment)
    }
}

