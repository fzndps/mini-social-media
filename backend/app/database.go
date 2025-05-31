package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/fzndps/mini-social-media/backend/helper"
)

func NewDB() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, port, dbName)

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
