package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getAllBlogs(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	rows, err := db.Query(`
		SELECT id, title, content, slug, created_at 
		FROM blogs 
		ORDER BY created_at DESC
	`)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []Blog

	for rows.Next() {
		var b Blog
		if err := rows.Scan(&b.ID, &b.Title, &b.Content, &b.Slug, &b.CreatedAt); err == nil {
			blogs = append(blogs, b)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func getBlogBySlug(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	slug := chi.URLParam(r, "slug")
	if slug == "" {
		http.Error(w, "Slug is required", http.StatusBadRequest)
		return
	}

	var b Blog
	err := db.QueryRow(`
		SELECT id, title, content, slug, created_at
		FROM blogs
		WHERE slug = $1
	`, slug).Scan(&b.ID, &b.Title, &b.Content, &b.Slug, &b.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Blog not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}


func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
