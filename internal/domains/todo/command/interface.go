package command

import (
	"context"

	"github.com/payfazz/fazzlearning-api/internal/domains/todo/data"
)

// TodoCommandInterface ...
type TodoCommandInterface interface {
	Create(ctx context.Context, payload data.PayloadCreateTodo) (*int64, error)
}
