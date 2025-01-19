package models

type Video struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
	Thumbnail   string `json:"thumbnail"`
}

type YouTubeVideo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"publishedAt"`
	Thumbnail   string `json:"thumbnail"`
}

type YouTubeAPIResponse struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			PublishedAt string `json:"publishedAt"`
			Thumbnails  struct {
				Default struct {
					URL string `json:"url"`
				} `json:"default"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}
