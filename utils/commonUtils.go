package utils

import (
	"cTube/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func FetchYouTubeVideos(apiKey, query string) ([]models.Video, error) {
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&key=%s&type=video&order=date", query, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	items, ok := result["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format: items not found or invalid")
	}

	var videos []models.Video
	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		snippet, ok := itemMap["snippet"].(map[string]interface{})
		if !ok {
			continue
		}

		title, ok := snippet["title"].(string)
		if !ok {
			title = "Untitled"
		}

		description, ok := snippet["description"].(string)
		if !ok {
			description = "No description available"
		}

		publishedAtStr, ok := snippet["publishedAt"].(string)
		if !ok {
			continue
		}

		publishedAt, err := parseTime(publishedAtStr)
		if err != nil {
			continue
		}

		thumbnailURL := ""
		if thumbnails, ok := snippet["thumbnails"].(map[string]interface{}); ok {
			if defaultThumb, ok := thumbnails["default"].(map[string]interface{}); ok {
				if url, ok := defaultThumb["url"].(string); ok {
					thumbnailURL = url
				}
			}
		}

		video := models.Video{
			Title:       title,
			Description: description,
			PublishedAt: publishedAt,
			Thumbnail:   thumbnailURL,
		}
		videos = append(videos, video)
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
