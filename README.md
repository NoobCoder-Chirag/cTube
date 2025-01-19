# cTube - YouTube Video Fetcher API

cTube is a backend API built using Go and the Gin framework to fetch the latest videos from YouTube for a given search query. It continuously fetches data in the background and stores it in a PostgreSQL database. The API provides paginated responses, search functionality, and efficient handling of YouTube API quota limits.

## Features

- Fetch and store the latest YouTube videos asynchronously.
- Paginated GET API to retrieve video data sorted by publishing date in descending order.
- Search API to search videos by title or description with support for partial matches.
- Dockerized for easy deployment.
- Scalable architecture with proper file structure and database indexing.

## File Structure

```plaintext
.
├── main.go                // Entry point of the application
├── config/
│   ├── config.go          // Application configuration
├── models/
│   ├── video.go           // Database model for videos
├── repositories/
│   ├── video_repository.go // Raw SQL queries for database interactions
├── services/
│   ├── video_service.go   // Business logic for fetching and managing videos
├── handlers/
│   ├── video_handler.go   // API handlers for GET and search APIs
├── router/
│   ├── router.go          // Routes setup
├── utils/
│   ├── time_parser.go     // Utility for parsing time
├── docker/
│   ├── Dockerfile         // Dockerfile for the project
│   ├── docker-compose.yml // Docker Compose configuration
└── README.md              // Documentation
