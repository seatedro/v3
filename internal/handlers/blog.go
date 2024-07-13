package handlers

import (
	"net/http"

	"github.com/seatedro/v3/internal/templates"
	"github.com/seatedro/v3/internal/utils"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserData()
	if err != nil {
		http.Error(w, "Error parsing user data", http.StatusInternalServerError)
		return
	}

	posts, err := utils.GetAllPosts()
	if err != nil {
		http.Error(w, "Error getting posts", http.StatusInternalServerError)
		return
	}

	if err := templates.BlogComponent(user, posts).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func BlogPostHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/blog/"):]

	userData, err := utils.GetUserData()
	if err != nil {
		http.Error(w, "Error reading user data", http.StatusInternalServerError)
		return
	}

	post, err := utils.GetPostBySlug(slug)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	templates.BlogPost(userData, post).Render(r.Context(), w)
}
