package main

import "fmt"

//Nao tem operador ternario
func main() {
	fmt.Println(obterResultado(6.2))
}

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	//return (nota >= 6 ? "Aprovado", "Reprovado")
}
