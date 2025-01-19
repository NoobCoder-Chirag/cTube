package handlers

import (
	"cTube/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoHandler struct {
	Service *services.VideoService
}

func NewVideoHandler(service *services.VideoService) *VideoHandler {
	return &VideoHandler{Service: service}
}

func (h *VideoHandler) GetVideos(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sortOrder := c.Query("sortOrder")
	offset := (page - 1) * size
	videos, err := h.Service.GetVideos(offset, size, sortOrder)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}

func (h *VideoHandler) SearchVideos(c *gin.Context) {
	keyword := c.Query("keyword")
	videos, err := h.Service.SearchVideos(keyword)
	if err != nil {
		fmt.Errorf("error getting searched videos %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, videos)
}
