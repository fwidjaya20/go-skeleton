package migration

import (
	"github.com/payfazz/fazzlearning-api/internal/values/todo"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

var VersionExample = fazzdb.MigrationVersion{
	Tables: []*fazzdb.MigrationTable{
		fazzdb.CreateTable("todos", func(table *fazzdb.MigrationTable) {
			table.Field(fazzdb.CreateSerial("id").Primary())
			table.Field(fazzdb.CreateString("activity").Unique())
			table.Field(fazzdb.CreateText("description").Nullable())
			table.Field(fazzdb.CreateString("start_date").Nullable())
			table.Field(fazzdb.CreateString("due_date").Nullable())
			table.Field(fazzdb.CreateString("status").Nullable().Default(todo.ACTIVE))
			table.TimestampsTz(0)
			table.SoftDeleteTz(0)
			table.Indexes("activity")
		}),
	},
}
