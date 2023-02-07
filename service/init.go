package service

import (
	"database/sql"
	database "github.com/raythx98/go-url-shortener/database/sqlc/output"
	"github.com/raythx98/go-url-shortener/service/receiver"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	db, err := sql.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/go_url_shortener")
	if err != nil {
		panic(err)
	}

	receiver.DbInstance.Queries = database.New(db)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
