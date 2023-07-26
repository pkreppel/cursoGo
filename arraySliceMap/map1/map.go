package main

import "fmt"

func main() {
	//Maps devem ser inicializados
	//var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678] = "Pedro"
	aprovados[12345] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345])
	fmt.Println(aprovados[123])

	delete(aprovados, 12345)
	fmt.Println(aprovados)
}
