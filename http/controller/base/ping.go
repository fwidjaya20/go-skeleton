package base

import (
	"net/http"

	"github.com/payfazz/go-apt/pkg/fazzcommon/response"
)

// Ping is a function that used to check the service status
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Text(w, "Ready to serve", http.StatusOK)
		return
	}
}
