package testutils

import "database/sql"

var db *sql.DB

func MysqlDbConnString() string {
	return "root:secret@tcp(127.0.0.1:3306)/Academia"
}
func MysqlDBConnection() *sql.DB {
	return db
}
