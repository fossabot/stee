package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/milanrodriguez/stee/stee"
)

func handleAPI(core *stee.Core, APIReservedURLPrefix string) http.Handler {
	router := httprouter.New()

	// General API route(s), not yet implemented
	router.GET(APIReservedURLPrefix+"/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "Error 501: API not implemented in this version. Try to upgrade Stee.")
	})

	// Simple API routes
	router.GET(APIReservedURLPrefix+"/simple/add/:key", handleSimpleAdd(core))
	router.GET(APIReservedURLPrefix+"/simple/add/:key/:base64target", handleSimpleAdd(core))
	router.GET(APIReservedURLPrefix+"/simple/get/:key", handleSimpleGet(core))
	router.GET(APIReservedURLPrefix+"/simple/del/:key", handleSimpleDel(core))

	return router
}
