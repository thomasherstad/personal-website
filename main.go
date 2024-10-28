package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 30,
	}

	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/projects", handlerProjects)
	mux.HandleFunc("/about", handlerAbout)
	mux.HandleFunc("/contact", handlerContact)

	// Serve static files from the static folder
	statics := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", statics))

	// Serve assets from the assets folder
	assets := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assets))

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
