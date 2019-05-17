package todo

import (
	"net/http"

	"github.com/payfazz/fazzlearning-api/internal/contract"

	"github.com/payfazz/fazzlearning-api/http/controller/value"
	"github.com/payfazz/fazzlearning-api/internal/domains/todo"
	"github.com/payfazz/go-apt/pkg/fazzcommon/httpError"
	"github.com/payfazz/go-apt/pkg/fazzcommon/request"
	"github.com/payfazz/go-apt/pkg/fazzcommon/response"
)

// All ...
func All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result *contract.Paginate

		params := map[string]string{
			value.LIMIT: "",
			value.PAGE:  "",
		}

		ctx := r.Context()

		queryParams, err := request.ParseQueryParam(r, params)

		if nil != err {
			response.Error(w, httpError.BadRequest(err))
			return
		}

		todoService := todo.NewTodoService()
		result, err = todoService.All(ctx, queryParams)

		if nil != err {
			response.Error(w, httpError.BadRequest(err))
			return
		}

		response.Json(w, result, http.StatusOK)
	}
}
