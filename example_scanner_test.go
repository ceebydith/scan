package scan_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/blockloop/scan"
	_ "github.com/mattn/go-sqlite3"
)

func exampleDB() *sql.DB {
	return mustDB(&testing.T{}, `CREATE TABLE persons (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(120) NOT NULL DEFAULT ''
	);
	INSERT INTO PERSONS (name)
	VALUES ('brett'), ('fred');`)
}

func ExampleRow() {
	db := exampleDB()
	rows, err := db.Query("SELECT * FROM persons LIMIT 1")
	if err != nil {
		panic(err)
	}

	var person struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}

	err = scan.Row(&person, rows)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(os.Stdout).Encode(&person)
	// Output:
	// {"ID":1,"Name":"brett"}
}

func ExampleRow_scalar() {
	db := exampleDB()
	rows, err := db.Query("SELECT name FROM persons LIMIT 1")
	if err != nil {
		panic(err)
	}

	var name string

	err = scan.Row(&name, rows)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%q", name)
	// Output:
	// "brett"
}

func ExampleRows() {
	db := exampleDB()
	rows, err := db.Query("SELECT * FROM persons ORDER BY name")
	if err != nil {
		panic(err)
	}

	var persons []struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}

	err = scan.Rows(&persons, rows)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(os.Stdout).Encode(&persons)
	// Output:
	// [{"ID":1,"Name":"brett"},{"ID":2,"Name":"fred"}]
}

func ExampleRows_primitive() {
	db := exampleDB()
	rows, err := db.Query("SELECT name FROM persons ORDER BY name")
	if err != nil {
		panic(err)
	}

	var names []string
	err = scan.Rows(&names, rows)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(os.Stdout).Encode(&names)
	// Output:
	// ["brett","fred"]
}
