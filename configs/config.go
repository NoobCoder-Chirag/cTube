package configs

import (
	"database/sql"
	"fmt"
)

func ConnectToDB() (*sql.DB, error) {
	connStr := "postgresql://postgres:LAxsyKTnIbUZjuXUWXCAfyZhYSjfFIKk@postgres.railway.internal:5432/railway"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Errorf("error connecting to database: %v", err)
	}

	return db, err
}
