package http

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/milanrodriguez/stee/stee"
)

func handleSimpleAdd(core *stee.Core) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		key := ps.ByName("key")
		targetBytes, err := base64.URLEncoding.DecodeString(ps.ByName("base64target"))
		if err != nil {
			targetBytes, err = base64.RawURLEncoding.DecodeString(ps.ByName("base64target"))
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		target := string(targetBytes)
		err = core.AddRedirection(key, target)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		fmt.Fprintf(w, "Added redirection: %s -> %s", key, target)
	}
}

func handleSimpleGet(core *stee.Core) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		key := ps.ByName("key")
		target, err := core.GetRedirection(key)
		if err != true {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Impossible to read key")
			return
		}
		fmt.Fprintf(w, "Key \"%s\" is pointing to \"%s\"", key, target)
	}
}

func handleSimpleDel(core *stee.Core) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		key := ps.ByName("key")
		err := core.DeleteRedirection(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
		fmt.Fprintf(w, "Deleted key \"%s\"", key)
	}
}
