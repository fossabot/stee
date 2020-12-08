package http

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/milanrodriguez/stee/stee"
)

func handleAPI(core *stee.Core, APIReservedURLPrefix string) http.Handler {
	router := httprouter.New()
	router.GET(APIReservedURLPrefix + "", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "Error 501: API not implemented in this version. Try to upgrade Stee.")
	})
	router.GET(APIReservedURLPrefix + "simple/add/:key/:base64target", handleSimpleAdd(core))
	router.GET(APIReservedURLPrefix + "simple/del/:key", handleSimpleDel(core))
	return router
}

func handleSimpleAdd (core *stee.Core) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		key := ps.ByName("key")
		targetBytes, err := base64.URLEncoding.DecodeString(ps.ByName("base64target"))
		if err != nil { targetBytes, err = base64.RawURLEncoding.DecodeString(ps.ByName("base64target")) }
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		target := string(targetBytes)
		err = core.SetRedirectionTarget(key, target)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		fmt.Fprintf(w, "Added redirection: %s -> %s", key, target)
	}
}

func handleSimpleDel (core *stee.Core) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		key := ps.ByName("key")
		err := core.DelRedirectionTarget(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		fmt.Fprintf(w, "Deleted key \"%s\"", key)
	}
}