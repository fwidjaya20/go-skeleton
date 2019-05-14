package collection

import (
	"github.com/payfazz/fazzlearning-api/http/controller/todo"
	"github.com/payfazz/go-apt/pkg/fazzrouter"
)

// RouteVersion1 is a routes for version 1 app
func RouteVersion1(r *fazzrouter.Route) {
	r.Prefix("v1", func(r *fazzrouter.Route) {
		r.Prefix("todos", func(r *fazzrouter.Route) {
			r.Get("", todo.All())
			r.Post("", todo.Create())
		})
	})
}
