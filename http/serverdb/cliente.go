package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Usuario struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

// Usuario handler
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :()")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	dataSourceName := fmt.Sprintf("host=127.0.0.1 port=5432 user=usuario " +
		"password=senha dbname=dbGo sslmode=disable")

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = $1", id).Scan(&u.Id, &u.Nome)

	json, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	dataSourceName := fmt.Sprintf("host=127.0.0.1 port=5432 user=usuario " +
		"password=senha dbname=dbGo sslmode=disable")

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, nome FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.Id, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}
	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
