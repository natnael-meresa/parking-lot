package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB establishes a connection to the database
func ConnectDB(url string) (*sql.DB, error) {
	var err error
	db, err := sql.Open("mysql", url) // Replace "mysql" with your database driver
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Ping the database to check the connection and ensure readiness
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, err
}
