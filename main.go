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
	fmt.Println("hii")
	db, err := configs.ConnectToDB()
	fmt.Println("hii again")
	if err != nil {
		fmt.Errorf("error connecting to database: %v", err)
	}

	defer db.Close()
	videoRepo := repository.NewVideoRepository(db)
	videoService := services.NewVideoService(videoRepo)
	videoHandler := handlers.NewVideoHandler(videoService)

	go func() {
		apiKey := constants.YoutubeDataApiKey
		for {
			videos, err := utils.FetchYouTubeVideos(apiKey, "hiphop")
			if err == nil {
				for _, video := range videos {
					fmt.Println("videos are being saved")
					videoService.SaveVideo(video)
				}
			}
			time.Sleep(10 * time.Second)
		}
	}()

	router := router.SetUpRouter(videoHandler)
	router.Run(":8080")
}
