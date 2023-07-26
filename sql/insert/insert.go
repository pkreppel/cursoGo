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

	stmt, _ := db.Prepare("insert into usuarios (nome) values($1)")
	stmt.Exec("Maria")

	res, _ := stmt.Exec("Pedro")
	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)

}
