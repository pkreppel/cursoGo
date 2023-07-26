package main

import (
	"database/sql"
	"fmt"

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

	stmt, _ := tx.Prepare("update usuarios set nome = $1 where id = $2")
	stmt.Exec("Maria 2", 1)
	stmt.Exec("José", 2)

	stmt2, _ := tx.Prepare("delete from usuarios where id = $1")

	stmt2.Exec(200)

	// _, err = stmt.Exec(200, "Pedro 2")
	// if err != nil {
	// 	tx.Rollback()
	// 	log.Fatal(err)
	// 	panic(err)
	// }
	tx.Commit()

}
