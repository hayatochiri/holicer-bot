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

func openDBonMemory(t *testing.T) {
	var err error

	db, err = sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("Error occurred when sql.Open() (%v)", err)
	}
}

func expectTable(t *testing.T, table_name string, expect_definitions map[string]tableDefinitions) error {
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

	t.Log(`Scan '` + table_name + `' table`)
	query := `PRAGMA table_info(` + table_name + `);`
	rows, err := db.Query(query)
	if err != nil {
		t.Fatalf("Error occurred when db.Query(\"%v\") (%v)", query, err)
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

func expectTablesList(t *testing.T, expect_tables []string) error {
	t.Log("Check if the tables matches the list %v", expect_tables)

	query := `select name from sqlite_master where type='table';`
	rows, err := db.Query(query)
	if err != nil {
		t.Fatalf("Error occurred when db.Query(\"%v\") (%v)", query, err)
	}

	table_existence := make(map[string]bool)
	for _, table_name := range expect_tables {
		table_existence[table_name] = false
	}

	// 'sqlite_sequence' テーブルはSQLite3が内部的に生成するため存在確認をtrueにする
	table_existence[`sqlite_sequence`] = true

	var table_name string
	for rows.Next() {
		rows.Scan(&table_name)

		for _, expect := range expect_tables {
			if expect == table_name {
				table_existence[table_name] = true
			}
		}

		if table_existence[table_name] == false {
			t.Errorf("Unexpected table \"%v\" was found", table_name)
		}
	}

	for table_name, exist := range table_existence {
		if !exist {
			t.Errorf("Table \"%v\" is missing", table_name)
		}
	}

	return nil
}

func TestInitializeDB(t *testing.T) {
	openDBonMemory(t)
	defer db.Close()

	if err := initializeDB(); err != nil {
		t.Fatalf("Error occurred when initializeDB() (%v)", err)
	}
}

func TestCreateDB(t *testing.T) {
	openDBonMemory(t)
	defer db.Close()

	if err := createDB(); err != nil {
		t.Fatalf("Error occurred when createDB() (%v)", err)
	}

	expectTable(
		t, `master`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`db_version`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTablesList(
		t,
		[]string{`master`},
	)

}

func TestUpdateDBv1(t *testing.T) {
	openDBonMemory(t)
	defer db.Close()

	if err := createDB(); err != nil {
		t.Fatalf("Error occurred when createDB() (%v)", err)
	}

	if err := updateDBv1(); err != nil {
		t.Fatalf("Error occurred when updateDB() (%v)", err)
	}

	expectTablesList(
		t,
		[]string{`master`, `taverns`, `groups`, `users`, `menus`, `users_log`, `leave_log`, `orders_log`},
	)

	expectTable(
		t, `master`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`db_version`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `taverns`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`name_jp`:    {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`name_en`:    {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`is_removed`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `groups`,
		map[string]tableDefinitions{
			`id`:           {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`uuid`:         {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`name_jp`:      {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`name_en`:      {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`started_time`: {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`tavern_id`:    {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`total_price`:  {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`tax_rate`:     {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`cleard_time`:  {Type: `text`, Notnull: FALSE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `users`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`twitter_id`: {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`email`:      {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`name`:       {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`avator`:     {Type: `blob`, Notnull: FALSE, Dflt_value: ``},
			`status`:     {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`group_id`:   {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `menus`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`tavern_id`:  {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`name_jp`:    {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`name_en`:    {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`price`:      {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`tax_rate`:   {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`is_removed`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `users_log`,
		map[string]tableDefinitions{
			`id`:        {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`user_id`:   {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`group_id`:  {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`status`:    {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`timestamp`: {Type: `text`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `leave_log`,
		map[string]tableDefinitions{
			`id`:          {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`user_log_id`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`pay`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `orders_log`,
		map[string]tableDefinitions{
			`id`:        {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`user_id`:   {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`group_id`:  {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`menu_id`:   {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`split`:     {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`status`:    {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`timestamp`: {Type: `text`, Notnull: TRUE, Dflt_value: ``},
		},
	)

}

func TestUpdateDB(t *testing.T) {
	openDBonMemory(t)
	defer db.Close()

	if err := createDB(); err != nil {
		t.Fatalf("Error occurred when createDB() (%v)", err)
	}

	if err := updateDB(); err != nil {
		t.Fatalf("Error occurred when updateDB() (%v)", err)
	}

	expectTablesList(
		t,
		[]string{`master`, `taverns`, `groups`, `users`, `menus`, `users_log`, `leave_log`, `orders_log`},
	)

	expectTable(
		t, `master`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`db_version`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `taverns`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`name_jp`:    {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`name_en`:    {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`is_removed`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `groups`,
		map[string]tableDefinitions{
			`id`:           {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`uuid`:         {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`name_jp`:      {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`name_en`:      {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`started_time`: {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`tavern_id`:    {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`total_price`:  {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`tax_rate`:     {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`cleard_time`:  {Type: `text`, Notnull: FALSE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `users`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`twitter_id`: {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`email`:      {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`name`:       {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`avator`:     {Type: `blob`, Notnull: FALSE, Dflt_value: ``},
			`status`:     {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`group_id`:   {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `menus`,
		map[string]tableDefinitions{
			`id`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`tavern_id`:  {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`name_jp`:    {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`name_en`:    {Type: `text`, Notnull: FALSE, Dflt_value: ``},
			`price`:      {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`tax_rate`:   {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`is_removed`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `users_log`,
		map[string]tableDefinitions{
			`id`:        {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`user_id`:   {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`group_id`:  {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`status`:    {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`timestamp`: {Type: `text`, Notnull: TRUE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `leave_log`,
		map[string]tableDefinitions{
			`id`:          {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`user_log_id`: {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`pay`:         {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
		},
	)

	expectTable(
		t, `orders_log`,
		map[string]tableDefinitions{
			`id`:        {Type: `integer`, Notnull: FALSE, Dflt_value: ``},
			`user_id`:   {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`group_id`:  {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`menu_id`:   {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`split`:     {Type: `integer`, Notnull: TRUE, Dflt_value: ``},
			`status`:    {Type: `text`, Notnull: TRUE, Dflt_value: ``},
			`timestamp`: {Type: `text`, Notnull: TRUE, Dflt_value: ``},
		},
	)

}
