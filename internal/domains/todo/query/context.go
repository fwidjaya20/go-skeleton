package query

import "github.com/payfazz/fazzlearning-api/internal/domains/todo/repository"

type todoQuery struct {
	Repository repository.TodoRepositoryInterface
}
