package command

import (
	"context"

	"github.com/payfazz/fazzlearning-api/config"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/data"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/repository"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// NewTodoCommand is a function to create new Command Instance
func NewTodoCommand() TodoCommandInterface {
	db := config.GetDb()
	q := fazzdb.QueryDb(db, config.GetIfQueryConfig(config.I_QUERY_CONFIG))

	return &todoCommand{
		Repository: repository.NewTodoRepository(q),
	}
}

func (c *todoCommand) Create(ctx context.Context, payload data.PayloadCreateTodo) (*int64, error) {
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

func (c *todoCommand) Update(ctx context.Context, payload data.PayloadUpdateTodo, id int64) (bool, error) {
	t := model.TodoModel()

	t.ID = id
	t.Activity = payload.Activity
	t.Description = payload.Description
	t.StartDate = payload.StartDate
	t.DueDate = payload.DueDate

	res, err := c.Repository.Update(ctx, t)

	if !res {
		// You can replace any log logic here
		return false, err
	}

	return true, err
}

func (c *todoCommand) Delete(ctx context.Context, id int64) (bool, error) {
	t := model.TodoModel()

	t.ID = id

	res, err := c.Repository.Delete(ctx, t)

	if !res {
		// You can replace any log logic here
		return false, err
	}

	return true, err
}
