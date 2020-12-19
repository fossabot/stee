package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/milanrodriguez/stee/internal/stee"
)

func handleMain(core *stee.Core) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path

		key = strings.TrimPrefix(key, "/")
		key = strings.TrimSuffix(key, "/")

		target, err := core.GetRedirection(key)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Error 404: No redirection found for key \"%s\"", key)
			return
		}
		http.Redirect(w, r, target, http.StatusFound)
	})
}
