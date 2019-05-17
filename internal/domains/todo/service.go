package todo

import (
	"context"

	"github.com/payfazz/fazzlearning-api/internal/contract"

	"github.com/payfazz/fazzlearning-api/internal/domains/todo/command"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/data"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/query"
)

// ServiceInterface ...
type ServiceInterface interface {
	All(ctx context.Context, queryParams map[string]string) (*contract.Paginate, error)
	Create(ctx context.Context, payload data.PayloadCreateTodo) (*model.Todo, error)
}

type service struct {
	Query   query.TodoQueryInterface
	Command command.TodoCommandInterface
}

// NewTodoService is a constructor for Todo Service
func NewTodoService() ServiceInterface {
	return &service{
		Query:   query.NewTodoQuery(),
		Command: command.NewTodoCommand(),
	}
}

func (s *service) Create(ctx context.Context, payload data.PayloadCreateTodo) (*model.Todo, error) {
	id, err := s.Command.Create(ctx, payload)

	if nil != err {
		return nil, err
	}

	return s.Query.Find(ctx, *id)
}

func (s *service) All(ctx context.Context, queryParams map[string]string) (*contract.Paginate, error) {
	// mestinya bikin struct isi array model dan metadata karena sifat paginate --> butuh total data

	res, err := s.Query.All(ctx, queryParams)

	if nil != err {
		return nil, err
	}

	c, err := s.Query.Count(ctx)

	if nil != err {
		return nil, err
	}

	d := &contract.Paginate{}

	d.Data = res
	d.Metadata = contract.PaginateMetadata{
		Total: int64(*c),
	}

	return d, nil
}
