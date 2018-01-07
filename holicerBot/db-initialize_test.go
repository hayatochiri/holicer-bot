package holicerBot

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	FALSE int = 0
	TRUE  int = 1
)

var TO_BOOL map[int]string = map[int]string{0: `FALSE`, 1: `TRUE`}

type tableDefinitions struct {
	Type       string
	Notnull    int
	Dflt_value string
}

func openDBonMemory(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("Error occurred when sql.Open() (%v)", err)
	}

	return db
}

func tableExpect(t *testing.T, rows *sql.Rows, expect_definitions map[string]tableDefinitions) error {
	var actual_cid int
	var actual_name string
	var actual_type string
	var actual_notnull int
	var actual_dflt_value string
	var actual_pk int

	field_existence := make(map[string]bool)
	for k := range expect_definitions {
		field_existence[k] = false
	}

	for rows.Next() {
		rows.Scan(&actual_cid, &actual_name, &actual_type, &actual_notnull, &actual_dflt_value, &actual_pk)

		expect := expect_definitions[actual_name]

		if expect.Type == `` {
			t.Errorf("Unexpected field \"%v\" was found", actual_name)
			t.Logf("cid           : %v", actual_cid)
			t.Logf("name          : %v", actual_name)
			t.Logf("type          : %v", actual_type)
			t.Logf("not null      : %v", actual_notnull)
			t.Logf("default value : %v", actual_dflt_value)
			t.Logf("primary key   : %v", actual_pk)
			continue
		}

		field_existence[actual_name] = true

		if actual_type != expect.Type {
			t.Errorf("type of field \"%v\" is incorrect", actual_name)
			t.Logf("actual : %v", actual_type)
			t.Logf("expect : %v", expect.Type)
		}

		if actual_notnull != expect.Notnull {
			t.Errorf("notnull of field \"%v\" is incorrect", actual_name)
			t.Logf("actual : %v", TO_BOOL[actual_notnull])
			t.Logf("expect : %v", TO_BOOL[expect.Notnull])
		}

		if actual_dflt_value != expect.Dflt_value {
			t.Errorf("dflt_value of field \"%v\" is incorrect", actual_name)
			t.Logf("actual : %v", actual_dflt_value)
			t.Logf("expect : %v", expect.Dflt_value)
		}
	}

	for field_name, exist := range field_existence {
		if !exist {
			t.Errorf("Field \"%v\" is missing", field_name)
			t.Logf("name          : %v", field_name)
			t.Logf("type          : %v", expect_definitions[field_name].Type)
			t.Logf("not null      : %v", expect_definitions[field_name].Notnull)
			t.Logf("default value : %v", expect_definitions[field_name].Dflt_value)
		}
	}

	return nil
}

func TestInitializeDB(t *testing.T) {
	db := openDBonMemory(t)
	defer db.Close()

	if err := initializeDB(db); err != nil {
		t.Fatalf("Error occurred when initializeDB() (%v)", err)
	}
}
