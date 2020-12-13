package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/milanrodriguez/stee/stee"
)

type rootHandler struct {
	http.Handler
	core    *stee.Core
	general http.Handler // Note: the "general handler" is the handler for the "normal" requests, the one that handles redirections.
	api     struct {
		enable    bool
		prefix    string
		simpleAPI struct {
			enable bool
		}
		http.Handler
	}
	ui struct {
		enable bool
		prefix string
		http.Handler
	}
}

type handleRootOption func(*rootHandler)

// HandleRoot returns a http.Handler in charge of dispatching requests to the appropriate "sub"-handler
func HandleRoot(options ...handleRootOption) http.Handler {

	rh := &rootHandler{}

	for _, option := range options {
		option(rh)
	}

	rh.general = handleMain(rh.core)
	if rh.api.enable {
		rh.api.Handler = handleAPI(rh.core, rh.api.prefix, rh.api.simpleAPI.enable)
	}
	if rh.ui.enable {
		rh.ui.Handler = handleUI(rh.core, rh.ui.prefix)
	}

	rh.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		r = r.WithContext(ctx)
		switch {
		case strings.HasPrefix(r.URL.Path, rh.api.prefix) && rh.api.enable:
			rh.api.ServeHTTP(w, r)
		case strings.HasPrefix(r.URL.Path, rh.ui.prefix) && rh.ui.enable:
			rh.ui.ServeHTTP(w, r)
		default:
			rh.general.ServeHTTP(w, r)
		}
	})

	return rh
}

func Core(core *stee.Core) handleRootOption {
	return func(rh *rootHandler) {
		rh.core = core
	}
}

func EnableAPI(enable bool, prefix string) handleRootOption {
	return func(rh *rootHandler) {
		rh.api.enable = enable
		if enable {
			rh.api.prefix = cleanPrefix(prefix)
		}
	}
}

func EnableSimpleAPI(enable bool) handleRootOption {
	return func(rh *rootHandler) {
		rh.api.simpleAPI.enable = enable
	}
}

func EnableUI(enable bool, prefix string) handleRootOption {
	return func(rh *rootHandler) {
		rh.ui.enable = enable
		if enable {
			rh.ui.prefix = cleanPrefix(prefix)
		}
	}
}

func cleanPrefix(prefix string) string {
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	if strings.HasSuffix(prefix, "/") {
		prefix = (prefix)[:len(prefix)]
	}
	return prefix
}
