package query

import (
	"context"

	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
)

// TodoQueryInterface ...
type TodoQueryInterface interface {
	All(ctx context.Context, queryParams map[string]string) ([]*model.Todo, error)
	Find(ctx context.Context, id int64) (*model.Todo, error)
}
