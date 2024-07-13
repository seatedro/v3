package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/seatedro/v3/internal/models"
	"github.com/seatedro/v3/internal/templates"
	"github.com/seatedro/v3/internal/utils"
)

func main() {
	// Read user data
	userData, err := os.ReadFile("content/user.json")
	if err != nil {
		fmt.Println("Error reading user data:", err)
		return
	}

	var user models.User
	err = json.Unmarshal(userData, &user)
	if err != nil {
		fmt.Println("Error parsing user data:", err)
		return
	}

	// Get all posts
	posts, err := utils.GetAllPosts()
	if err != nil {
		fmt.Println("Error getting posts:", err)
		return
	}

	// Create output directory
	err = os.MkdirAll("public", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	// Generate home page
	homeHTML := &bytes.Buffer{}
	err = templates.HomeComponent(user, posts[:3]).Render(context.Background(), homeHTML)
	if err != nil {
		fmt.Println("Error rendering home page:", err)
		return
	}
	err = os.WriteFile("public/index.html", homeHTML.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing home page:", err)
		return
	}

	// Generate blog list page
	/* blogHTML := &bytes.Buffer{}
	err = templates.Blog(user, posts).Render(context.Background(), blogHTML)
	if err != nil {
		fmt.Println("Error rendering blog page:", err)
		return
	}
	err = os.MkdirAll("public/blog", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating blog directory:", err)
		return
	}
	err = ioutil.WriteFile("public/blog/index.html", blogHTML.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing blog page:", err)
		return
	}

	// Generate individual blog post pages
	for _, post := range posts {
		postHTML := &bytes.Buffer{}
		err = templates.BlogPost(user, post).Render(context.Background(), postHTML)
		if err != nil {
			fmt.Printf("Error rendering blog post %s: %v\n", post.Slug, err)
			continue
		}
		err = ioutil.WriteFile(filepath.Join("public/blog", post.Slug+".html"), postHTML.Bytes(), 0644)
		if err != nil {
			fmt.Printf("Error writing blog post %s: %v\n", post.Slug, err)
		}
	} */

	fmt.Println("Static site generated successfully!")
}
