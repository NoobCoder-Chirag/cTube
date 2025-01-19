package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	host     = "monorail.proxy.rlwy.net"
	port     = "48730"
	user     = "postgres"
	password = "LAxsyKTnIbUZjuXUWXCAfyZhYSjfFIKk"
	dbname   = "railway"
)

func ConnectToDB() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db, nil
}
