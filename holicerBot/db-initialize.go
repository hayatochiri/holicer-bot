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

	err = updateDB(db)
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

		insert into master (id, db_version)
		select 1, 0 where (select count(*) from master) = 0;
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

func updateDB(db *sql.DB) error {
	var db_version int
	query := `select db_version from master where id = 1`

	for already_updated := false; already_updated == false; {

		if err := db.QueryRow(query).Scan(&db_version); err != nil {
			return err
		}

		switch db_version {
		default:
			already_updated = true
		}

	}

	return nil
}
