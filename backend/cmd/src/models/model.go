package models

import "github.com/jinzhu/gorm"
import "time"

type Post struct {
	ID int
	gorm.Model
	Title    string
	Text     string
	PostedOn time.Time
	Section  []Section
}

type Section struct {
	ID     int
	PostID uint
	gorm.Model
	Name string
}
