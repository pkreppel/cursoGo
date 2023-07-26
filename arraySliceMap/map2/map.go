package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao": 11235.45,
		"Gabi":      15232.245,
		"PEdro":     23423.45,
	}

	funcsESalarios["Rafael"] = 23424.5

	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "inexistente") // nao encontra e nao apaga

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
