package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/seatedro/v3/internal/models"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v2"
)

func GetRecentPosts(count int) ([]models.Post, error) {
	posts, err := GetAllPosts()
	if err != nil {
		return nil, err
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	if len(posts) > count {
		posts = posts[:count]
	}

	return posts, nil
}

func GetAllPosts() ([]models.Post, error) {
	files, err := os.ReadDir("content/posts")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var posts []models.Post

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		content, err := os.ReadFile(fmt.Sprintf("content/posts/%s", file.Name()))
		if err != nil {
			return nil, err
		}

		var post models.Post
		rest, err := frontmatter.Parse(bytes.NewReader(content), &post, frontmatter.NewFormat("---", "---", yaml.Unmarshal))
		if err != nil {
			return nil, err
		}

		// Parse the remaining content as Markdown
		var buf bytes.Buffer
		markdown := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("onedark"),
					highlighting.WithFormatOptions(chromahtml.WithLineNumbers(true)),
				),
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithUnsafe(),
			),
		)
		if err := markdown.Convert(rest, &buf); err != nil {
			return nil, err
		}
		post.Content = buf.String()

		// Generate a slug from the title
		fileNameWithoutExtension := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		post.Slug = generateSlug(fileNameWithoutExtension)

		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts, nil
}

func SearchBlogPosts(q string, limit int) ([]models.Post, error) {
	posts, err := GetAllPosts()
	if err != nil {
		return nil, err
	}

	filteredPosts := make([]models.Post, 0)
	for _, post := range posts {
		if fuzzy.Match(strings.ToLower(q), strings.ToLower(post.Title)) {
			filteredPosts = append(filteredPosts, post)
		}
	}

	if len(filteredPosts) > limit {
		filteredPosts = filteredPosts[:limit]
	}

	return filteredPosts, nil
}

func GetPostBySlug(slug string) (models.Post, error) {
	posts, err := GetAllPosts()
	if err != nil {
		return models.Post{}, err
	}

	for _, post := range posts {
		if post.Slug == slug {
			return post, nil
		}
	}

	return models.Post{}, fmt.Errorf("Post not found!")
}

func generateSlug(title string) string {
	// Convert to lowercase and replace spaces with hyphens
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove any characters that aren't alphanumeric or hyphens
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, slug)

	return slug
}

func GetUserData() (models.User, error) {
	userData, err := os.ReadFile("content/user.json")
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = json.Unmarshal(userData, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetKinoTech() ([]models.KinoTech, error) {
	kinoTech, err := os.ReadFile("content/kino.json")
	if err != nil {
		return []models.KinoTech{}, err
	}

	var kino []models.KinoTech
	err = json.Unmarshal(kinoTech, &kino)
	if err != nil {
		return []models.KinoTech{}, err
	}

	return kino, nil
}

func GetDailyLog() ([]models.Log, error) {
	data, err := os.ReadFile("content/log.json")
	if err != nil {
		return []models.Log{}, err
	}

	var log []models.Log
	err = json.Unmarshal(data, &log)
	if err != nil {
		fmt.Println(err)
		return []models.Log{}, err
	}

	sort.Slice(log, func(i, j int) bool {
		date1, err := time.Parse("2006-01-02", log[i].Date)
		if err != nil {
			return false
		}
		date2, err := time.Parse("2006-01-02", log[j].Date)
		if err != nil {
			return false
		}

		return date1.After(date2)
	})

	if len(log) > 10 {
		log = log[:10]
	}

	return log, nil
}
