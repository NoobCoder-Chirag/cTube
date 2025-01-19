package constants

const (
	GetVideos = `SELECT id, title, description, url, created_at FROM videos ORDER BY created_at $1 limit $2 offset $3`

	InsertVideo = `INSERT INTO videos (title, description, published_at, thumbnail) VALUES ($1, $2, $3, $4)`

	SearchVideos = `SELECT id, title, description, published_at, thumbnail FROM videos WHERE title ILIKE $1 OR description ILIKE $1`
)
