package store

import (
	"fmt"
	"testing"
)

// TestStore ..
func TestStore(t *testing.T, databaseURL string) (*Store, func(tables []string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL

	store := New(config)
	if err := store.Open(); err != nil {
		t.Fatal(err)
	}

	return store, func(tables []string) {
		for _, table := range tables {
			sql := "TRUNCATE TABLE %s;"

			if _, err := store.db.Exec(fmt.Sprintf(sql, table)); err != nil {
				t.Fatal(err)
			}
		}

		store.Close()
	}
}
