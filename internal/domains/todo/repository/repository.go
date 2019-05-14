package repository

import (
	"context"

	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

type TodoRepositoryInterface interface {
	Find(ctx context.Context, id int64) (*model.Todo, error)
	All(ctx context.Context, conditions []fazzdb.SliceCondition, orders []fazzdb.Order, limit int, offset int) ([]*model.Todo, error)
	Create(ctx context.Context, m *model.Todo) (*int64, error)
}

type todoRepository struct {
	Q    *fazzdb.Query
	Todo *model.Todo
}

// NewTodoRepository is a constructor to construct repo
func NewTodoRepository(q *fazzdb.Query) TodoRepositoryInterface {
	return &todoRepository{
		Q:    q,
		Todo: model.TodoModel(),
	}
}

func (r *todoRepository) Create(ctx context.Context, m *model.Todo) (*int64, error) {
	result, err := r.Q.Use(m).InsertCtx(ctx, false)

	if nil != err {
		return nil, err
	}

	id := result.(int64)

	return &id, nil
}

// Find a function that used to find the data by id
func (r *todoRepository) Find(ctx context.Context, id int64) (*model.Todo, error) {
	rows, err := r.Q.Use(r.Todo).
		Where("id", id).
		WithLimit(1).
		AllCtx(ctx)

	if nil != err {
		return nil, err
	}

	results := rows.([]*model.Todo)
	if len(results) == 0 {
		return nil, nil
	}

	result := results[0]
	if nil != err {
		return nil, err
	}

	return result, nil
}

func (r *todoRepository) All(ctx context.Context, conditions []fazzdb.SliceCondition, orders []fazzdb.Order, limit int, offset int) ([]*model.Todo, error) {
	current := r.Q.Use(r.Todo).
		WhereMany(conditions...).
		OrderByMany(orders...)

	if limit > 0 {
		current.WithLimit(limit)
	}

	if offset > 0 {
		current.WithOffset(offset)
	}

	rows, err := current.AllCtx(ctx)

	if nil != err {
		return nil, err
	}

	result := rows.([]*model.Todo)

	return result, nil
}
