package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SirodgevAlex/comments-service/graph"
	"github.com/SirodgevAlex/comments-service/graph/model"
	"github.com/SirodgevAlex/comments-service/repository"
)

const defaultPort = "8080"

func main() {
	// Проверяем наличие аргумента командной строки для базы данных
	useDatabase := false
	if len(os.Args) > 1 && os.Args[1] == "db" {
		useDatabase = true
	}

	// Создаем репозиторий в зависимости от выбора
	var repo repository.Repository
	if useDatabase {
		// Инициализируйте репозиторий базы данных здесь
	} else {
		repo = &repository.InMemoryRepository{
			Posts:     make(map[int]*model.Post),
			Comments:  make(map[int][]*model.Comment),
			IDCounter: 0,
		}
	}

	// Используем репозиторий в сервисе
	service := graph.NewService(repo)

	// Инициализируем сервер GraphQL
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: service}))

	// Настройка обработчиков маршрутов
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Запуск сервера на порту 8080
	port := defaultPort
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
