package collection

import (
	"github.com/payfazz/fazzlearning-api/http/controller/base"
	"github.com/payfazz/go-apt/pkg/fazzrouter"
)

// BaseRoute is a base routes
func BaseRoute(r *fazzrouter.Route) {
	r.Get("ping", base.Ping())
}
