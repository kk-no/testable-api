package mysql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/kk-no/testable-api/database"
)

func init() {
	conn, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("Failed to open connection for database: %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to communicate with the database: %v", err)
	}

	database.Conn = conn
}
