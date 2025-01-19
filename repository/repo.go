package repository

import (
	"cTube/constants"
	"cTube/models"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type VideoRepository struct {
	DB *sql.DB
}

func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{DB: db}
}

func (repo *VideoRepository) GetVideos(offset, limit int, sortOrder string) ([]models.Video, error) {
	query := constants.GetVideos
	rows, err := repo.DB.Query(query, sortOrder, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var video models.Video
		if err = rows.Scan(&video.ID, &video.Title, &video.Description, &video.Thumbnail, &video.PublishedAt); err != nil {
			fmt.Errorf("error getting videos %v", err)
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (repo *VideoRepository) InsertVideo(video models.YouTubeVideo) error {
	query := constants.InsertVideo
	videoWithId := models.Video{
		ID:          uuid.New().String(),
		Title:       video.Title,
		Description: video.Description,
		PublishedAt: video.PublishedAt,
		Thumbnail:   video.Thumbnail,
	}

	result, err := repo.DB.Exec(query, videoWithId.ID, videoWithId.Title, videoWithId.Description, videoWithId.PublishedAt, videoWithId.Thumbnail)

	fmt.Println("checkpoint 2")
	fmt.Println(result)
	if err != nil {
		fmt.Errorf("error inserting video %v", err)
		return err
	}

	fmt.Println("video saved")
	return nil
}

func (repo *VideoRepository) SearchVideos(keyword string) ([]models.Video, error) {
	searchQuery := `%` + keyword + `%`
	sqlQuery := constants.SearchVideos
	rows, err := repo.DB.Query(sqlQuery, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var video models.Video
		if err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.PublishedAt, &video.Thumbnail); err != nil {
			fmt.Errorf("error getting videos %v", err)
			continue
		}
		videos = append(videos, video)
	}
	return videos, nil
}
