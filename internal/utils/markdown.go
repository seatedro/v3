package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/seatedro/v3/internal/models"
	"github.com/yuin/goldmark"
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
		if err := goldmark.Convert(rest, &buf); err != nil {
			return nil, err
		}
		post.Content = buf.String()

		// Generate a slug from the title
		post.Slug = generateSlug(post.Title)

		posts = append(posts, post)
	}

	return posts, nil
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
