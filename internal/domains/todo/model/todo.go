package model

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// Todo ...
type Todo struct {
	fazzdb.Model `json:"-"`
	ID           int64  `db:"id" json:"id"`
	Activity     string `db:"activity" json:"activity"`
	Description  string `db:"description" json:"description"`
	StartDate    string `db:"start_date" json:"start_date"`
	DueDate      string `db:"due_date" json:"due_date"`
	Status       string `db:"status" json:"status"`
}

// Payload is a function to read the payload data
func (m *Todo) Payload() map[string]interface{} {
	return m.MapPayload(m)
}

// TodoModel is function to get Todo Model
func TodoModel() *Todo {
	return &Todo{
		Model: fazzdb.AutoIncrementModel("todos",
			[]fazzdb.Column{
				fazzdb.Col("id"),
				fazzdb.Col("activity"),
				fazzdb.Col("description"),
				fazzdb.Col("start_date"),
				fazzdb.Col("due_date"),
				fazzdb.Col("status"),
			},
			"id",
			true,
			true,
		),
	}
}
