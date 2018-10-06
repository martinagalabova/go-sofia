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
		log.Print("The hello handler was called.")
		fmt.Fprintf(w, http.StatusText(http.StatusOK))
	})

	possibleErrors := make(chan error, 2)

	go func() {
		log.Print("The application server is about to handle connections...")
		server := &http.Server{
			Addr:    ":" + blPort,
			Handler: router,
		}
		err := server.ListenAndServe()
		if err != nil {
			possibleErrors <- err
		}
	}()

	select {
	case err := <-possibleErrors:
		log.Fatal(err)
	}

	diagnostics := diagnostics.NewDiagnostics()
	log.Print("The diagnostics server is about to handle connections...")
	err := http.ListenAndServe(":"+diagPort, diagnostics)
	if err != nil {
		possibleErrors <- err
	}
}
