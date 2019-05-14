package app

import (
	"net/http"

	"github.com/payfazz/fazzlearning-api/config"
	"github.com/payfazz/go-apt/pkg/fazzdb"
	"github.com/payfazz/go-middleware/common/kv"
)

type keyType struct{}

var key keyType

// App is a struct that will be send in the context
type App struct {
	QueryDb *fazzdb.Query
}

// New is a function that as a http handler
func New() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			db := config.GetDb()

			kv.Set(r, key, &App{
				QueryDb: fazzdb.QueryDb(db, config.GetIfQueryConfig(config.I_QUERY_CONFIG)),
			})

			next(w, r)
		}
	}
}

// GetApp is a function that used to get the app context
func GetApp(r *http.Request) *App {
	return kv.MustGet(r, key).(*App)
}
