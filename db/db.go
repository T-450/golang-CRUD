package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// ConectaComBancoDeDados retorna db
func ConectaComBancoDeDados() *sql.DB {
	connStr := "user=postgres password='123456' dbname=AluraLoja sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db;
}