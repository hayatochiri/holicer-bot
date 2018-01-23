package holicerBot

import (
	_ "github.com/mattn/go-sqlite3"
)

type Tavern struct {
	Id        int64
	NameJA    string
	NameEN    string
	IsRemoved int
}

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
