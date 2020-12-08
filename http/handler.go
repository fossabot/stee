package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/milanrodriguez/stee/stee"
)

type HandlersSet struct {
	Main http.Handler // For browser redirection, "normal usage"
	API http.Handler // For handling api requests
	UI http.Handler // For handling the web app UI
}

func RootHandler(core *stee.Core, APIReservedURLPrefix string, UIReservedURLPrefix string) http.Handler {
    if core == nil {
        panic("no core!")
	}

	handlers := HandlersSet{
		Main: handleMain(core),
		API: handleAPI(core, APIReservedURLPrefix),
		UI: handleUI(core, UIReservedURLPrefix),
	}

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		ctx := context.Background()
		r = r.WithContext(ctx)
        switch {
		case strings.HasPrefix(r.URL.Path, APIReservedURLPrefix):
			handlers.API.ServeHTTP(w,r)
		case strings.HasPrefix(r.URL.Path, UIReservedURLPrefix):
			handlers.UI.ServeHTTP(w,r)
		default:
			handlers.Main.ServeHTTP(w,r)
		}
    })
}