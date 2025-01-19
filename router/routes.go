package router

import (
	"cTube/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(videoHandler *handlers.VideoHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/videos", videoHandler.GetVideos)
	router.GET("/search", videoHandler.SearchVideos)

	return router
}
