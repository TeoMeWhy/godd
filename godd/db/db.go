package db

import (
	"database/sql"
	"io/ioutil"
	"log"

	// Load sqlite driver onto the sql registry.
	_ "github.com/mattn/go-sqlite3"
)

// OpenSQLite creates a connection to SQLite engine by path.
func OpenSQLite(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Println(err)
		panic("Não foi possível conectar ao banco")
	}
	return db
}

// ImportQuery imports a query from a .txt file by path.
func ImportQuery(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Não foi possível abrir o arquivo da query")
	}
	query := string(data)
	return query
}
