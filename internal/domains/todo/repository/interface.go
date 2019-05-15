package repository

import (
	"context"

	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// TodoRepositoryInterface ...
type TodoRepositoryInterface interface {
	Find(ctx context.Context, id int64) (*model.Todo, error)
	All(ctx context.Context, conditions []fazzdb.SliceCondition, orders []fazzdb.Order, limit int, offset int) ([]*model.Todo, error)
	Create(ctx context.Context, m *model.Todo) (*int64, error)
}
