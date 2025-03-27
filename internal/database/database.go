package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "172.19.0.2"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydb"
)

var (
	dbInstance *sql.DB
	once       sync.Once
	initErr    error
)

func GetDB() (*sql.DB, error) {
	once.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
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

func DatabaseConnectionTest() error {
	_, err := GetDB()
	return err
}
