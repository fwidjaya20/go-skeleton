package todo

import (
	"net/http"

	"github.com/payfazz/fazzlearning-api/internal/domains/todo"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/data"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo/model"
	"github.com/payfazz/go-apt/pkg/fazzcommon/httpError"
	"github.com/payfazz/go-apt/pkg/fazzcommon/request"
	"github.com/payfazz/go-apt/pkg/fazzcommon/response"
)

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result *model.Todo
		var body data.PayloadCreateTodo
		var err error

		ctx := r.Context()

		err = request.ParseJson(r, &body)
		if nil != err {
			response.Error(w, httpError.BadRequest(err))
			return
		}

		todoService := todo.NewTodoService()
		result, err = todoService.Create(ctx, body)

		if nil != err {
			response.Error(w, httpError.BadRequest(err))
			return

		}

		response.Json(w, result, http.StatusCreated)
	}
}
