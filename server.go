package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SirodgevAlex/comments-service/graph"
	"github.com/SirodgevAlex/comments-service/repository"
)

const defaultPort = "8080"

func main() {
	useDatabase := false
	if len(os.Args) > 1 && os.Args[1] == "db" {
		useDatabase = true
	}

	var repo repository.Repository
	if useDatabase {
		connStr := "postgres://postgres:1234@host.docker.internal:5432/comments-system?sslmode=disable"
		var err error
		repo, err = repository.NewPostgresRepository(connStr)
		if err != nil {
			log.Fatalf("Failed to create PostgreSQL repository: %v", err)
		}
	} else {
		repo = repository.NewInMemoryRepository()
	}

	service := graph.NewService(repo)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: service}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := defaultPort
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	if useDatabase {
		repository.ClosePostgresRepository(repo.(*repository.PostgresRepository).DB)
	}

	log.Println("Server gracefully stopped")
}
