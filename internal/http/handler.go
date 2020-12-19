package http

import (
	"net/http"
	"strings"

	"github.com/milanrodriguez/stee/internal/stee"
)

type rootHandler struct {
	http.Handler
	core    *stee.Core
	general struct { // Note: the "general handler" is the handler for the "normal" requests, the one that handles user redirections.
		http.Handler
	}
	api struct {
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
func HandleRoot(core *stee.Core, options ...handleRootOption) http.Handler {

	rh := &rootHandler{core: core}

	for _, option := range options {
		option(rh)
	}

	rh.general.Handler = handleMain(rh.core)

	if rh.api.enable {
		rh.api.Handler = handleAPI(rh.core, rh.api.prefix, rh.api.simpleAPI.enable)
	}
	if rh.ui.enable {
		rh.ui.Handler = handleUI(rh.core, rh.ui.prefix)
	}

	rh.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case rh.api.enable && strings.HasPrefix(r.URL.Path, rh.api.prefix):
			rh.api.ServeHTTP(w, r)
		case rh.ui.enable && strings.HasPrefix(r.URL.Path, rh.ui.prefix):
			rh.ui.ServeHTTP(w, r)
		default:
			rh.general.ServeHTTP(w, r)
		}
	})

	return rh
}

// EnableAPI is an option for HandleRoot().
// It enables the API.
//
// It should only be used as an argument for HandleRoot(), hence the unexported return type.
func EnableAPI(enable bool, prefix string) handleRootOption {
	return func(rh *rootHandler) {
		rh.api.enable = enable
		if enable {
			rh.api.prefix = cleanPathPrefix(prefix)
		}
	}
}

// EnableSimpleAPI is an option for HandleRoot().
// It enables the simple API.
//
// It should only be used as an argument for HandleRoot(), hence the unexported return type.
func EnableSimpleAPI(enable bool) handleRootOption {
	return func(rh *rootHandler) {
		rh.api.simpleAPI.enable = enable
	}
}

// EnableUI is an option for HandleRoot().
// It enables the UI.
//
// It should only be used as an argument for HandleRoot(), hence the unexported return type.
func EnableUI(enable bool, prefix string) handleRootOption {
	return func(rh *rootHandler) {
		rh.ui.enable = enable
		if enable {
			rh.ui.prefix = cleanPathPrefix(prefix)
		}
	}
}

// cleanPathPrefixForHandlers returns the given prefix with a leadind slash and without a trailing slash.
func cleanPathPrefix(prefix string) string {
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	prefix = strings.TrimSuffix(prefix, "/")
	return prefix
}
