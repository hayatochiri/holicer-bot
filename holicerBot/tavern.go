package holicerBot

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type Tavern struct {
	Id          int64
	NameJA      string
	NameEN      string
	Description string
	Locate      string
	Homepage    string
	IsRemoved   int
}

type AddTavernParams struct {
	NameJA      string
	NameEN      string
	Description string
	Locate      string
	Homepage    string
}

func AddTavern(params AddTavernParams) (int64, error) {
	var err error

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	query := `
		insert into taverns (name_jp, name_en, description, locate, homepage, is_removed)
		values (?, ?, ?, ?, ?, ?);
	`

	result, err := tx.Exec(query, params.NameJA, params.NameEN, params.Description, params.Locate, params.Homepage, FALSE)
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

type GetTavernsListParams struct {
	IsRemoved bool
}

func GetTavernsList(params GetTavernsListParams) ([]Tavern, error) {
	var taverns_list []Tavern
	var tavern_row Tavern
	query := `select * from taverns where is_removed = ?`

	rows, err := db.Query(query, map[bool]int{true: TRUE, false: FALSE}[params.IsRemoved])
	if err != nil {
		return []Tavern{}, nil
	}

	for rows.Next() {
		rows.Scan(&tavern_row.Id, &tavern_row.NameJA, &tavern_row.NameEN, &tavern_row.IsRemoved)
		taverns_list = append(taverns_list, tavern_row)
	}

	return taverns_list, nil
}

func RemoveTavern(remove_id int64) (bool, error) {

	query := `update taverns set is_removed = ? where id = ? and is_removed = ?`

	tx, err := db.Begin()
	if err != nil {
		return false, err
	}

	result, err := tx.Exec(query, TRUE, remove_id, FALSE)
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	is_removed, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if is_removed >= 2 {
		return false, errors.New("[Unexpected] More than one tavern has been removed.")
	}

	return map[int64]bool{0: false, 1: true}[is_removed], nil
}
