package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	dataSourceName := fmt.Sprintf("host=127.0.0.1 port=5432 user=usuario " +
		"password=senha dbname=dbGo sslmode=disable")

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuario where id > $1)", 5)
	defer rows.Close()
	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

}
