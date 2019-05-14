package command

import (
	"context"

	"github.com/payfazz/fazzlearning-api/internal/domains/todo/data"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/repository"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// TodoCommandInterface ...
type TodoCommandInterface interface {
	Create(ctx context.Context, payload data.PayloadCreateTodo) (*int64, error)
}

type partnerCommand struct {
	Repository repository.TodoRepositoryInterface
}

// NewTodoCommand is a function to create new Command Instance
func NewTodoCommand(q *fazzdb.Query) TodoCommandInterface {
	return &partnerCommand{
		Repository: repository.NewTodoRepository(q),
	}
}

func (c *partnerCommand) Create(ctx context.Context, payload data.PayloadCreateTodo) (*int64, error) {
	t := model.TodoModel()

	t.Activity = payload.Activity
	t.Description = payload.Description
	t.StartDate = payload.StartDate
	t.DueDate = payload.DueDate
	t.Status = payload.Status

	id, err := c.Repository.Create(ctx, t)

	if nil != err {
		return nil, err
	}

	return id, nil
}
