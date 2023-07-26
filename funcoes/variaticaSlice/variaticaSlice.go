package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "PEdro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...) //nao funciona com array, somente com slice
}
