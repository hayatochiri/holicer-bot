package holicerBot

import (
	_ "github.com/mattn/go-sqlite3"
)

func AddTavern(name_ja, name_en string) (int64, error) {
	var err error

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	query := `
		insert into taverns (name_jp, name_en, is_removed)
		values (?, ?, ?);
	`
	result, err := tx.Exec(query, name_ja, name_en, FALSE)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	inserted_id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return inserted_id, nil
}
