package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/seatedro/v3/internal/models"
	"github.com/seatedro/v3/internal/templates"
	"github.com/seatedro/v3/internal/utils"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	userData, err := os.ReadFile("content/user.json")
	if err != nil {
		http.Error(w, "Error reading user data", http.StatusInternalServerError)
		return
	}

	var user models.User
	err = json.Unmarshal(userData, &user)
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
