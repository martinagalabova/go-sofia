package diagnostics

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewDiagnostics() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthz)
	router.HandleFunc("/ready", ready)
	return router
}

func healthz(w http.ResponseWriter, r *http.Request) {
	log.Print("/healtz endpoint was called")
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}

func ready(w http.ResponseWriter, r *http.Request) {
	log.Print("/ready endpoint was called")
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
