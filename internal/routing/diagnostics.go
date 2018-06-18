package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewDiagnosticsRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/alive", diagnosticHandler())
	return r
}

func diagnosticHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("alive!"))
	}
}
