package testutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/kk-no/testable-api/database"
)

func SetupOptionalFixtures(names []string) {
	sqlDir := path.Clean("../test/sql/")
	for _, name := range names {
		p := filepath.Join(sqlDir, name)
		execSchema(p)
	}
}

func execSchema(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("schema reading error: %v", err)
	}

	queries := strings.Split(string(b), ";")

	for _, query := range queries[:len(queries)-1] {
		_, err = database.Conn.Exec(query)
		if err != nil {
			log.Fatalf("exec schema error: %v, query: %s", err, query)
		}
	}
}

func TruncateTables() {
	rows, err := database.Conn.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("show tables error: %#v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			log.Fatalf("show table error: %#v", err)
		}

		commands := []string{
			"SET FOREIGN_KEY_CHECKS = 0",
			fmt.Sprintf("TRUNCATE %s", tableName),
			"SET FOREIGN_KEY_CHECKS = 1",
		}
		for _, cmd := range commands {
			if _, err := database.Conn.Exec(cmd); err != nil {
				log.Fatalf("truncate error: %#v", err)
			}
		}
	}
}
