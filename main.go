package main

import (
	"cTube/configs"
	"cTube/constants"
	"cTube/handlers"
	"cTube/repository"
	"cTube/router"
	"cTube/services"
	"cTube/utils"
	"fmt"
	"time"
)

func main() {
	db, err := configs.ConnectToDB()
	if err != nil {
		fmt.Println("error->>>>", err)
	}

	//defer db.Close()
	videoRepo := repository.NewVideoRepository(db)
	videoService := services.NewVideoService(videoRepo)
	videoHandler := handlers.NewVideoHandler(videoService)

	go func() {
		apiKey := constants.YoutubeDataApiKey
		for {
			videos, err := utils.FetchYouTubeVideos(apiKey, "hiphop")
			if err != nil {
				fmt.Errorf("error fetching videos: %v", err)
				return
			}

			for _, video := range videos {
				fmt.Println("title ->", video.Title)
				fmt.Println("description ->", video.Description)
				fmt.Println("published at ->", video.PublishedAt)
				fmt.Println("thumbnail ->", video.Thumbnail)
				videoService.SaveVideo(video)
			}

			time.Sleep(10 * time.Second)
		}
	}()

	router := router.SetUpRouter(videoHandler)
	router.Run(":8080")
}
