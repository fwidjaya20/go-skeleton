package command

import "github.com/payfazz/fazzlearning-api/internal/domains/todo/repository"

type todoCommand struct {
	Repository repository.TodoRepositoryInterface
}
