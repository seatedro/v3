package main

import (
	"log"
	"net/http"

	"github.com/seatedro/v3/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/blog", handlers.BlogHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
