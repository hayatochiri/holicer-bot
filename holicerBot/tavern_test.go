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
	output += "\tDescription\n"
	output += fmt.Sprintf("\t\tactual : \"%s\"\n", actual.Description)
	output += fmt.Sprintf("\t\texpect : \"%s\"\n", expect.Description)
	output += "\tLocate\n"
	output += fmt.Sprintf("\t\tactual : \"%s\"\n", actual.Locate)
	output += fmt.Sprintf("\t\texpect : \"%s\"\n", expect.Locate)
	output += "\tHomepage\n"
	output += fmt.Sprintf("\t\tactual : \"%s\"\n", actual.Homepage)
	output += fmt.Sprintf("\t\texpect : \"%s\"\n", expect.Homepage)
	output += "\tIsRemoved\n"
	output += "\t\tactual : " + TO_BOOL[actual.IsRemoved] + "\n"
	output += "\t\texpect : " + TO_BOOL[expect.IsRemoved] + "\n"

	if actual.Id == expect.Id &&
		actual.NameJA == expect.NameJA &&
		actual.NameEN == expect.NameEN &&
		actual.Description == expect.Description &&
		actual.Locate == expect.Locate &&
		actual.Homepage == expect.Homepage &&
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
	expects := []Tavern{
		{NameJA: `居酒屋(No.1)`, NameEN: `Tavern(No.1)`, Description: `詳細01`, Locate: `場所001`, Homepage: `https://www.test.tavern1.local`},
		{NameJA: `居酒屋(No.2)`},
		{NameEN: `Tavern(No.3)`},
		{NameJA: `居酒屋(No.4)`, NameEN: `Tavern(No.4)`, Description: `詳細04`, Locate: `場所004`, Homepage: `https://www.test.tavern4.local`},
		{NameJA: `居酒屋(No.5)`, NameEN: `Tavern(No.5)`},
		{NameJA: `居酒屋(No.6)`, Description: `詳細06`, Locate: `場所006`, Homepage: `https://www.test.tavern6.local`},
		{NameEN: `Tavern(No.7)`, Description: `詳細07`, Locate: `場所007`, Homepage: `https://www.test.tavern7.local`},
		{NameJA: `居酒屋(No.8)`, NameEN: `Tavern(No.8)`, Description: `詳細08`, Locate: `場所008`, Homepage: `https://www.test.tavern8.local`},
		{NameJA: `居酒屋(No.9)`, NameEN: `Tavern(No.9)`, Description: `詳細09`, Locate: `場所009`, Homepage: `https://www.test.tavern9.local`},
		{NameJA: `居酒屋(No.10)`, NameEN: `Tavern(No.10)`, Description: `詳細10`, Locate: `場所010`, Homepage: `https://www.test.tavern10.local`},
	}

	if err := createDB(); err != nil {
		t.Fatalf("Error occurred when createDB() (%v)", err)
	}

	if err := updateDB(); err != nil {
		t.Fatalf("Error occurred when updateDB() (%v)", err)
	}

	for _, expect := range expects {
		_, err := AddTavern(AddTavernParams{
			NameJA:      expect.NameJA,
			NameEN:      expect.NameEN,
			Description: expect.Description,
			Locate:      expect.Locate,
			Homepage:    expect.Homepage,
		})
		if err != nil {
			t.Fatalf("Error occurred when AddTavern() (%v)", err)
		}
	}

	t.Logf("Inserted records.")
	query := `select * from taverns where id = ?`
	for id, expect := range expects {
		expect.Id = int64(id) + 1
		row := db.QueryRow(query, expect.Id)
		row.Scan(&actual.Id, &actual.NameJA, &actual.NameEN, &actual.Description, &actual.Locate, &actual.Homepage, &actual.IsRemoved)
		expect.IsRemoved = FALSE

		actual.compare(t, expect)
	}
}

func TestRemoveTavern(t *testing.T) {
	openDBonMemory(t)
	defer db.Close()

	if err := initializeDB(); err != nil {
		t.Fatalf("Error occurred when initializeDB() (%v)", err)
	}

	remove_list := []Tavern{
		{NameJA: `居酒屋(No.1)`, NameEN: `Tavern(No.1)`, Description: `詳細01`, Locate: `場所001`, Homepage: `https://www.test.tavern1.local`, IsRemoved: FALSE},
		{NameJA: `居酒屋(No.2)`, NameEN: `Tavern(No.2)`, Description: `詳細02`, Locate: `場所002`, Homepage: `https://www.test.tavern2.local`, IsRemoved: TRUE},
		{NameJA: `居酒屋(No.3)`, NameEN: `Tavern(No.3)`, IsRemoved: TRUE},
		{NameJA: `居酒屋(No.4)`, NameEN: `Tavern(No.4)`, Description: `詳細04`, Locate: `場所004`, Homepage: `https://www.test.tavern4.local`, IsRemoved: TRUE},
		{NameJA: `居酒屋(No.5)`, NameEN: `Tavern(No.5)`, Description: `詳細05`, Locate: `場所005`, Homepage: `https://www.test.tavern5.local`, IsRemoved: FALSE},
		{NameJA: `居酒屋(No.6)`, NameEN: `Tavern(No.6)`, Description: `詳細06`, Locate: `場所006`, Homepage: `https://www.test.tavern6.local`, IsRemoved: TRUE},
		{NameJA: `居酒屋(No.7)`, NameEN: `Tavern(No.7)`, Description: `詳細07`, Locate: `場所007`, Homepage: `https://www.test.tavern7.local`, IsRemoved: FALSE},
		{NameJA: `居酒屋(No.8)`, NameEN: `Tavern(No.8)`, Description: `詳細08`, Locate: `場所008`, Homepage: `https://www.test.tavern8.local`, IsRemoved: FALSE},
		{NameJA: `居酒屋(No.9)`, NameEN: `Tavern(No.9)`, Description: `詳細09`, Locate: `場所009`, Homepage: `https://www.test.tavern9.local`, IsRemoved: FALSE},
		{NameJA: `居酒屋(No.10)`, NameEN: `Tavern(No.10)`, Description: `詳細10`, Locate: `場所010`, Homepage: `https://www.test.tavern10.local`, IsRemoved: TRUE},
	}

	for _, item := range remove_list {
		inserted_id, err := AddTavern(AddTavernParams{
			NameJA:      item.NameJA,
			NameEN:      item.NameEN,
			Description: item.Description,
			Locate:      item.Locate,
			Homepage:    item.Homepage,
		})
		if err != nil {
			t.Fatalf("Error occurred when AddTavern() (%v)", err)
		}

		if item.IsRemoved == FALSE {
			continue
		}

		t.Logf("Remove tavern(%v)", inserted_id)
		is_removed, err := RemoveTavern(inserted_id)
		if err != nil {
			t.Fatalf("Error occurred when RemoveTavern() (%v)", err)
		}
		if !is_removed {
			t.Fatalf("Could not remove the record that could be removed")
		}

		is_removed, err = RemoveTavern(inserted_id)
		if err != nil {
			t.Fatalf("Error occurred when RemoveTavern() (%v)", err)
		}
		if is_removed {
			t.Fatalf("Removed the removed record")
		}
	}

	var actual Tavern
	for index, expect := range remove_list {
		expect.Id = int64(index) + 1
		query := `select * from taverns where id = ?`
		row := db.QueryRow(query, expect.Id)

		row.Scan(&actual.Id, &actual.NameJA, &actual.NameEN, &actual.Description, &actual.Locate, &actual.Homepage, &actual.IsRemoved)
		actual.compare(t, expect)
	}

}
