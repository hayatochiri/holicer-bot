package holicerBot

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func openDBonMemory(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("Error occurred when sql.Open() (%v)", err)
	}

	return db
}

func TestInitializeDB(t *testing.T) {
	db := openDBonMemory(t)
	defer db.Close()

	if err := initializeDB(db); err != nil {
		t.Fatalf("Error occurred when initializeDB() (%v)", err)
	}
}
