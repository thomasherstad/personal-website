package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	// const filePathRoot = "."
	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/projects", handlerProjects)
	mux.HandleFunc("/about", handlerAbout)
	mux.HandleFunc("/contact", handlerContact)

	log.Println("Personal website server running on port", port)
	log.Fatal(srv.ListenAndServe())
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func handlerProjects(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/projects.html")
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/about.html")
}

func handlerContact(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/contact.html")
}
