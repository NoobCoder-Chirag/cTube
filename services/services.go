package services

import (
	"cTube/models"
	"cTube/repository"
)

type VideoService struct {
	Repo *repository.VideoRepository
}

func NewVideoService(repo *repository.VideoRepository) *VideoService {
	return &VideoService{Repo: repo}
}

func (s *VideoService) GetVideos(offset, limit int, sortOrder string) ([]models.Video, error) {
	return s.Repo.GetVideos(offset, limit, sortOrder)
}

func (s *VideoService) SaveVideo(video models.YouTubeVideo) error {
	return s.Repo.InsertVideo(video)
}

func (s *VideoService) SearchVideos(keyword string) ([]models.Video, error) {
	return s.Repo.SearchVideos(keyword)
}
