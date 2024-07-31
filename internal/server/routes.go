package server

import (
	// "context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/seatedro/v3/internal/models"
	"github.com/seatedro/v3/internal/templates"
	"github.com/seatedro/v3/internal/utils"
)

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
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

	recentPosts, err := utils.GetRecentPosts(3)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting recent posts", http.StatusInternalServerError)
		return
	}

	if err := templates.HomeComponent(user, recentPosts).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) BlogHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) BlogPostHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) MetricsStreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Create a channel to signal when the client has disconnected
	clientClosed := r.Context().Done()

	// Create a ticker to send updates to the client every 60 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	streamMetrics := func() {
		metrics, err := s.dbQueries.GetLatestMetrics(r.Context())
		if err != nil {
			fmt.Printf("Coudln't fetch metrics: %v", err)
			fmt.Fprintf(w, "error: %v\n\n", err)
			flusher.Flush()
			return
		}

		data, err := json.Marshal(metrics)
		if err != nil {
			fmt.Fprintf(w, "error: %v\n\n", err)
			flusher.Flush()
			return
		}

		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()
	}

	streamMetrics()

	for {
		select {
		case <-clientClosed:
			// The client has disconnected, so we return from the handler.
			log.Println("Client Disconnected.")
			return
		case <-ticker.C:
			streamMetrics()
		}
	}
}

func (s *Server) BlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "3"
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 3
	}

	posts, err := utils.GetRecentPosts(limitInt)
	if err != nil {
		http.Error(w, "Error getting posts", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, "Error converting posts to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (s *Server) SearchBlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "Query parameter not provided", http.StatusBadRequest)
		return
	}

	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "3"
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 3
	}

	posts, err := utils.SearchBlogPosts(q, limitInt)
	if err != nil {
		http.Error(w, "Error getting posts", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, "Error converting posts to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
