package route

import (
	"net/http"

	"github.com/payfazz/fazzlearning-api/http/app"
	"github.com/payfazz/fazzlearning-api/http/middleware"
	"github.com/payfazz/fazzlearning-api/http/route/collection"
	"github.com/payfazz/go-apt/pkg/fazzrouter"
	"github.com/payfazz/go-middleware/common/kv"
)

// Compile is a function to compile the data
func Compile() http.HandlerFunc {
	r := fazzrouter.BaseRoute()
	r.Use(
		kv.New(),
		app.New(),
		middleware.Cors(),
	)

	collection.BaseRoute(r)
	collection.RouteVersion1(r)

	return r.Compile()
}
