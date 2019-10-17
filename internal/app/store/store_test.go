package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "goapi:secret@tcp(172.21.0.45)/goapi_test"
	}

	os.Exit(m.Run())
}
