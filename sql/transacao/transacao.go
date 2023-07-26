package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dataSourceName := fmt.Sprintf("host=127.0.0.1 port=5432 user=usuario " +
		"password=senha dbname=dbGo sslmode=disable")

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into usuarios (id, nome) values($1, $2)")
	stmt.Exec(100, "Maria 2")

	_, err = stmt.Exec(200, "Pedro 2")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		panic(err)
	}
	tx.Commit()

}
