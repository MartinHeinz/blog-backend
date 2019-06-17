package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Blog Post
type Post struct {
	gorm.Model
	Title    string
	Text     string
	PostedOn time.Time
	Section  []Section
	Tag      []Tag
}

// Section of Blog Post (headings)
type Section struct {
	PostID uint
	gorm.Model
	Name string
}

// Tag of Blog Post (hashtag)
type Tag struct {
	PostID uint
	gorm.Model
	Name string
}
