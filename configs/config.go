package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	user     = "postgres.xmiebvtzreuumwzizwbv"
	password = "pass@Supabase0"
	host     = "aws-0-ap-south-1.pooler.supabase.com"
	port     = "5432"
	dbname   = "postgres"
)

func ConnectToDB() (*sql.DB, error) {

	fmt.Println("Connecting to DB...")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	fmt.Println("Successfully connected to database")
	return db, nil
}
