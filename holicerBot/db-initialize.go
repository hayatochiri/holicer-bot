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

	err = initializeDB(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	db.Close()
}

func initializeDB(db *sql.DB) error {
	var err error

	err = createDB(db)
	if err != nil {
		return err
	}

	err = updateDB(db)
	if err != nil {
		return err
	}

	return nil
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

	_, err = tx.Exec(query)
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
	var err error

	var db_version int
	query := `select db_version from master where id = 1`

	for already_updated := false; already_updated == false; {

		err = db.QueryRow(query).Scan(&db_version)
		if err != nil {
			return err
		}

		switch db_version {
		case 0:
			err = updateDBv1(db)
		default:
			already_updated = true
			err = nil
		}
		if err != nil {
			return err
		}

	}

	return nil
}

func updateDBv1(db *sql.DB) error {
	var err error

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := `
		update master set db_version = 1 where id = 1;

		create table taverns (
			id         integer primary key autoincrement,
			name_jp    text    null,
			name_en    text    null,
			is_removed integer not null
		);

		create table groups (
			id           integer primary key autoincrement,
			uuid         text    not null    unique,
			name_jp      text    null,
			name_en      text    null,
			started_time text    null,
			tavern_id    integer not null,
			total_price  integer null,
			tax_rate     integer null,
			cleard_time  text    null
		);

		create table users (
			id         integer primary key autoincrement,
			twitter_id text    null        unique,
			email      text    null        unique,
			name       text    not null,
			avator     blob    null,
			status     text    not null,
			group_id   integer null
		);

		create table menus (
			id         integer primary key autoincrement,
			tavern_id  integer not null,
			name_jp    text    null,
			name_en    text    null,
			price      integer not null,
			tax_rate   integer null,
			is_removed integer not null
		);

		create table users_log (
			id        integer primary key autoincrement,
			user_id   integer not null,
			group_id  integer not null,
			status    text    not null,
			timestamp text    not null
		);

		create table leave_log (
			id          integer primary key autoincrement,
			user_log_id integer not null    unique,
			pay         integer null
		);

		create table orders_log (
			id        integer primary key autoincrement,
			user_id   integer not null,
			group_id  integer not null,
			menu_id   integer not null,
			split     integer not null,
			status    text    not null,
			timestamp text    not null
		);
	`

	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
