package services

import "cTube/repository"

type VideoService struct {
	Repo *repository.VideoRepository
}

func NewVideoService(repo *repository.VideoRepository) *VideoService {
	return &VideoService{Repo: repo}
}
