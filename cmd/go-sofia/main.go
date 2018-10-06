package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/martinagalabova/go-sofia/internal/diagnostics"
)

func main() {
	log.Print("Starting...")

	blPort := os.Getenv("PORT")
	if len(blPort) == 0 {
		log.Print("No PORT provided, using 8080")
		blPort = "8080"
	}

	diagPort := os.Getenv("DIAG_PORT")
	if len(diagPort) == 0 {
		log.Print("No DIAG_PORT provided, using 8585")
		diagPort = "8585"
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, http.StatusText(http.StatusOK))
	})

	go func() {
		err := http.ListenAndServe(":"+blPort, router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	diagnostics := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":"+diagPort, diagnostics)
	if err != nil {
		log.Fatal(err)
	}
}
