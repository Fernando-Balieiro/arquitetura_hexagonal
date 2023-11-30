package db

import "database/sql"

var db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
}
