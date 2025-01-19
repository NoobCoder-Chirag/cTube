package repository

import (
	"cTube/constants"
	"cTube/models"
	"database/sql"
	"fmt"
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

func (repo *VideoRepository) InsertVideo(video models.Video) error {
	fmt.Println("hii")
	query := constants.InsertVideo
	_, err := repo.DB.Exec(query, video.Title, video.Description, video.PublishedAt, video.Thumbnail)
	if err != nil {
		fmt.Errorf("error inserting video %v", err)
		return err
	}

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
