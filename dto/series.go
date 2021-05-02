package dto

import (
	"time"

	"gorm.io/gorm"
)

type Series struct {
	gorm.Model
	PublishedAt time.Time
	Title       string
	Category    string
	Summary     string
	Posts       []Post
}
