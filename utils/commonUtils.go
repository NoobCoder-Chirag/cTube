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
	json.Unmarshal(body, &result)

	items := result["items"].([]interface{})
	var videos []models.Video

	for _, item := range items {
		snippet := item.(map[string]interface{})["snippet"].(map[string]interface{})
		video := models.Video{
			Title:       snippet["title"].(string),
			Description: snippet["description"].(string),
			PublishedAt: parseTime(snippet["publishedAt"].(string)),
			Thumbnail:   snippet["thumbnails"].(map[string]interface{})["default"].(map[string]interface{})["url"].(string),
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func parseTime(t string) time.Time {
	parsed, _ := time.Parse(time.RFC3339, t)
	return parsed
}
