package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/milanrodriguez/stee/stee"
)

func handleUI(core *stee.Core, UIReservedURLPrefix string) http.Handler {
	router := httprouter.New()
	router.GET(UIReservedURLPrefix + "/*ui", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "Error 501: UI not implemented in this version. Try to upgrade Stee.\n\nSee https://github.com/milanrodriguez/stee")
	})
	return router
}