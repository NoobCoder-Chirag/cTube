package utils

import (
	"cTube/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func FetchYouTubeVideos(apiKey, query string) ([]models.YouTubeVideo, error) {
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&type=video&key=%s", query, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch YouTube videos: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse models.YouTubeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Map the response to a slice of YouTubeVideo
	videos := []models.YouTubeVideo{}
	for _, item := range apiResponse.Items {
		publishedAtTime, err := parseTime(item.Snippet.PublishedAt)
		if err != nil {
			fmt.Printf("failed to parse published at time: %v", err)
		}
		videos = append(videos, models.YouTubeVideo{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishedAt: publishedAtTime,
			Thumbnail:   item.Snippet.Thumbnails.Default.URL,
		})
	}
	return videos, nil
}

func parseTime(t string) (time.Time, error) {
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		fmt.Errorf("error parsing time: %v", err)
	}
	return parsed, nil
}
