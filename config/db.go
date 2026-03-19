package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "2xkwvxorTfpY932.root:NlbBlPrB7YD8co7u@tcp(gateway01.ap-southeast-1.prod.aws.tidbcloud.com:4000)/GO?tls=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}