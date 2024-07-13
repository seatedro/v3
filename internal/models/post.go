package models

import "time"

type Post struct {
	Title   string    `yaml:"title"`
	Date    time.Time `yaml:"date"`
	Summary string    `yaml:"summary"`
	Tags    []string  `yaml:"tags"`
	Content string    `yaml:"-"` // This will be populated after parsing the Markdown content
	Slug    string    `yaml:"-"` // This will be generated from the title
}
