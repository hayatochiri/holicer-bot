package holicerBot

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	os.Create("./holicer-bot.db")

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
	query := `
		create table if not exists master (
			db_version string
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
