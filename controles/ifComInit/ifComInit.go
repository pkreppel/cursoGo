package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if i := numeroAleatorio(); i > 5 { // tambem suporta no switch
		fmt.Println("Ganhou")
	} else {
		fmt.Println("PErdeu")
	}
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
