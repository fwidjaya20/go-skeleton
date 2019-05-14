package todo

import (
	"net/http"

	"github.com/payfazz/fazzlearning-api/http/app"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/data"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/go-apt/pkg/fazzcommon/httpError"
	"github.com/payfazz/go-apt/pkg/fazzcommon/request"
	"github.com/payfazz/go-apt/pkg/fazzcommon/response"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result *model.Todo
		var body data.PayloadCreateTodo
		var err error

		ctx := r.Context()
		appCtx := app.GetApp(r)

		err = request.ParseJson(r, &body)
		if nil != err {
			response.Error(w, httpError.BadRequest(err))
			return
		}

		errTrans := fazzdb.RunDefault(appCtx.QueryDb.Db, func(q *fazzdb.Query) error {
			todoService := todo.NewTodoService(q)
			result, err = todoService.Create(ctx, body)

			if nil != err {
				return err
			}

			return nil
		})

		if nil != errTrans {
			response.Error(w, httpError.InternalServer(err))
			return
		}

		response.Json(w, result, http.StatusCreated)
	}
}
