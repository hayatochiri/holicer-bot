package holicerBot

import (
	"fmt"
	"testing"
)

func (actual Tavern) compare(t *testing.T, expect Tavern) {
	output := "Compare\n"
	output += "\tID\n"
	output += fmt.Sprintf("\t\tactual : %d\n", actual.Id)
	output += fmt.Sprintf("\t\texpect : %d\n", expect.Id)
	output += "\tNameJA\n"
	output += fmt.Sprintf("\t\tactual : \"%s\"\n", actual.NameJA)
	output += fmt.Sprintf("\t\texpect : \"%s\"\n", expect.NameJA)
	output += "\tNameEN\n"
	output += fmt.Sprintf("\t\tactual : \"%s\"\n", actual.NameEN)
	output += fmt.Sprintf("\t\texpect : \"%s\"\n", expect.NameEN)
	output += "\tIsRemoved\n"
	output += "\t\tactual : " + TO_BOOL[actual.IsRemoved] + "\n"
	output += "\t\texpect : " + TO_BOOL[expect.IsRemoved] + "\n"

	if actual.Id == expect.Id &&
		actual.NameJA == expect.NameJA &&
		actual.NameEN == expect.NameEN &&
		actual.IsRemoved == expect.IsRemoved {
		t.Logf(output)
	} else {
		t.Errorf(output)
	}
}

func TestAddTavern(t *testing.T) {
	openDBonMemory(t)
	defer db.Close()

	if err := initializeDB(); err != nil {
		t.Fatalf("Error occurred when initializeDB() (%v)", err)
	}

	var actual Tavern
	expect := Tavern{
		NameJA:    `居酒屋`,
		NameEN:    `Tavern`,
		IsRemoved: FALSE,
	}

	if err := createDB(); err != nil {
		t.Fatalf("Error occurred when createDB() (%v)", err)
	}

	if err := updateDB(); err != nil {
		t.Fatalf("Error occurred when updateDB() (%v)", err)
	}

	inserted_id, err := AddTavern("居酒屋", "Tavern")
	if err != nil {
		t.Fatalf("Error occurred when AddTavern() (%v)", err)
	}

	query := `select * from taverns where id = ?`
	row := db.QueryRow(query, inserted_id)

	row.Scan(&actual.Id, &actual.NameJA, &actual.NameEN, &actual.IsRemoved)
	expect.Id = inserted_id

	t.Logf("Inserted record.")
	actual.compare(t, expect)
}
