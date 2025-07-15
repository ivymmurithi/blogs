package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	initDB()

	r := chi.NewRouter()
	r.Get("/api/blogs", getAllBlogs)
	r.Get("/api/blog/{slug}", getBlogBySlug)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
