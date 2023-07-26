package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	dataSourceName := fmt.Sprintf("host=127.0.0.1 port=5432 user=usuario " +
		"password=senha dbname=dbGo sslmode=disable")

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//exec(db, "create database if not exists cursoGoTeste")
	//exec(db, "CONNECT dbGo; USE dbGo;")
	//exec(db, "drop table if exists usuario")
	exec(db, `create table if not exists usuarios (
		id SERIAL ,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
