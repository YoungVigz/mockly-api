package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/YoungVigz/mockly-api/internal/migrations"
	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

var (
	dbInstance *sql.DB
	once       sync.Once
	initErr    error
)

func GetDB() (*sql.DB, error) {
	once.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			initErr = err
			return
		}

		if err = db.Ping(); err != nil {
			initErr = err
			return
		}

		dbInstance = db
	})

	return dbInstance, initErr
}

func DatabaseInit() error {
	db, err := GetDB()

	if err != nil {
		return err
	}

	err = migrations.InitializeTables(db)

	return err
}
