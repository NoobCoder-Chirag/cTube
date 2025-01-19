package models

import "time"

type Video struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Thumbnail   string    `json:"thumbnail"`
}
