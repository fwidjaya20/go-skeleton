package todo

import (
	"net/http"

	"github.com/payfazz/fazzlearning-api/http/app"
	"github.com/payfazz/fazzlearning-api/http/controller/value"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/go-apt/pkg/fazzcommon/httpError"
	"github.com/payfazz/go-apt/pkg/fazzcommon/request"
	"github.com/payfazz/go-apt/pkg/fazzcommon/response"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// All ...
func All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result []*model.Todo

		params := map[string]string{
			value.LIMIT: "",
			value.PAGE:  "",
		}

		ctx := r.Context()
		appCtx := app.GetApp(r)

		queryParams, err := request.ParseQueryParam(r, params)

		if nil != err {
			response.Error(w, httpError.BadRequest(err))
			return
		}

		errTrans := fazzdb.RunDefault(appCtx.QueryDb.Db, func(q *fazzdb.Query) error {
			todoService := todo.NewTodoService(q)
			result, err = todoService.All(ctx, queryParams)

			if nil != err {
				return err
			}

			return nil
		})

		if nil != errTrans {
			response.Error(w, httpError.InternalServer(errTrans))
			return
		}

		response.Json(w, result, http.StatusOK)
	}
}
