package dto

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PublishedAt time.Time
	Title       string
	Category    string
	Summary     string
	Body        string
	SeriesID    uint
	Version     int
}
