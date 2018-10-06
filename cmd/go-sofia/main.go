package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/martinagalabova/go-sofia/internal/diagnostics"
)

type serverConf struct {
	port   string
	router http.Handler
	name   string
}

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
	diagnostics := diagnostics.NewDiagnostics()

	possibleErrors := make(chan error, 2)

	servConf := []serverConf{
		{
			port:   blPort,
			router: router,
			name:   "App server",
		},
		{
			port:   diagPort,
			router: diagnostics,
			name:   "Diagnostics server",
		},
	}

	servers := make([]*http.Server, 2)

	for i, conf := range servConf {
		go func(conf serverConf, i int) {
			log.Printf("The %s server is about to handle connections...", conf.name)
			servers[i] = &http.Server{
				Addr:    ":" + conf.port,
				Handler: conf.router,
			}
			err := servers[i].ListenAndServe()
			if err != nil {
				possibleErrors <- err
			}
		}(conf, i)
	}

	select {
	case err := <-possibleErrors:
		log.Fatal(err)
	}

	// log.Print("The diagnostics server is about to handle connections...")
	// err := http.ListenAndServe(":"+diagPort, diagnostics)
	// if err != nil {
	// 	possibleErrors <- err
	// }
}
