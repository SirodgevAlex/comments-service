package graph

import (
	"github.com/SirodgevAlex/comments-service/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Repo repository.Repository
}
