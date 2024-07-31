package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	db "github.com/seatedro/v3/internal/metrics-db"
)

type Server struct {
	db        *sql.DB
	dbQueries *db.Queries
}

func NewServer() *Server {
	if os.Getenv("GO_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	sqlDb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer sqlDb.Close()

	dbQueries := db.New(sqlDb)

	newServer := &Server{
		db:        sqlDb,
		dbQueries: dbQueries,
	}

	http.HandleFunc("/", newServer.HomeHandler)
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("robots.txt")
		if err != nil {
			http.Error(w, "Could not read robots.txt", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(data)
	})
	http.HandleFunc("/blog", newServer.BlogHandler)
	http.HandleFunc("/blog/", newServer.BlogPostHandler)
	http.HandleFunc("/api/metrics/stream", newServer.MetricsStreamHandler)
	http.HandleFunc("/api/blog-posts", newServer.BlogPostsHandler)
	http.HandleFunc("/api/blog-posts/search", newServer.SearchBlogPostsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	return newServer

}
