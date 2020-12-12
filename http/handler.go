package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/milanrodriguez/stee/stee"
)

// Note: the "general handler" is the handler for the "normal" requests, the one that handles redirections.

// HandlersSet is the set of "sub"-handlers for the different categories of requests (API, UI, Main).
type HandlersSet struct {
	General http.Handler // For browser redirection, "normal usage"
	API     http.Handler // For handling api requests
	UI      http.Handler // For handling the web app UI
}

// RootHandler returns a http.Handler in charge of dispatching requests to the appropriate "sub"-handler
func RootHandler(core *stee.Core, APIPrefix string, UIPrefix string) http.Handler {
	if core == nil {
		panic("RootHandler: no core!")
	}

	// Clean the prefixes
	prefixes := []*string{&APIPrefix, &UIPrefix}
	for i := range prefixes {
		if !strings.HasPrefix(*prefixes[i], "/") {
			*prefixes[i] = "/" + *prefixes[i]
		}
		if strings.HasSuffix(*prefixes[i], "/") {
			*prefixes[i] = (*prefixes[i])[:len(*prefixes[i])]
		}
	}

	handlers := HandlersSet{
		General: handleMain(core),
		API:     handleAPI(core, APIPrefix),
		UI:      handleUI(core, UIPrefix),
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		r = r.WithContext(ctx)
		switch {
		case strings.HasPrefix(r.URL.Path, APIPrefix):
			handlers.API.ServeHTTP(w, r)
		case strings.HasPrefix(r.URL.Path, UIPrefix):
			handlers.UI.ServeHTTP(w, r)
		default:
			handlers.General.ServeHTTP(w, r)
		}
	})
}
