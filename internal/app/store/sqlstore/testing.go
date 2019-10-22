package sqlstore

import (
	"database/sql"
	"fmt"
	"testing"
)

// TestDB ..
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(tables []string)) {
	t.Helper()

	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables []string) {
		for _, table := range tables {
			sql := "TRUNCATE TABLE %s;"

			if _, err := db.Exec(fmt.Sprintf(sql, table)); err != nil {
				t.Fatal(err)
			}
		}

		db.Close()
	}
}
