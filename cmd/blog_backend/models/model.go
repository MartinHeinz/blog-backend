package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title    string
	Text     string
	PostedOn time.Time
	Section  []Section
}

type Section struct {
	PostID uint
	gorm.Model
	Name string
}
