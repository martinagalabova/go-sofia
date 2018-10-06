package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Hello Go Sofia!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
