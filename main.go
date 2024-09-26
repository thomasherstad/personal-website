package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	const filePathRoot = "./templates/"
	const port = "8080"

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(filePathRoot)))

	log.Println("Personal website server running...")
	log.Fatal(srv.ListenAndServe())
}
