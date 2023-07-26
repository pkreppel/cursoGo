package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(tipo(1))
	fmt.Println(tipo(1.6))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não reconheço"

	}
}
