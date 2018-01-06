package holicerBot

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Initialize() {
	var db *sql.DB
	var err error

	db, err = sql.Open("sqlite3", "./holicer-bot.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	err = createDB(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	db.Close()
}

func createDB(db *sql.DB) error {
	var err error

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := `
		create table if not exists master (
			id         integer primary key,
			db_version integer not null
		);
	`

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
