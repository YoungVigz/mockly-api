package migrations

import (
	"database/sql"
	"fmt"
)

func InitializeTables(db *sql.DB) error {
	userTable := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(50) UNIQUE NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL
        );
    `

	schemasTable := `
        CREATE TABLE IF NOT EXISTS schemas (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
            content JSONB NOT NULL
        )
    `

	tables := []string{userTable, schemasTable}

	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}
	}

	return nil
}
