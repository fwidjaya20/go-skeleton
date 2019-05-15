package repository

import (
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

type todoRepository struct {
	Q    *fazzdb.Query
	Todo *model.Todo
}
