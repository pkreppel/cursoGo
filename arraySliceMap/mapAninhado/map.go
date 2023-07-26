package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 234234.345,
			"Guga":     35435.345,
		},
		"J": {
			"Joao": 23243.2,
			"Jose": 234233.3,
		},
		"P": {
			"Pedro": 234234,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
