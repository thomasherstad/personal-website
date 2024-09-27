package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	const filePathRoot = "./templates/"
	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(filePathRoot)))

	log.Println("Personal website server running on port", port)
	log.Fatal(srv.ListenAndServe())
}
