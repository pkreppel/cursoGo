package main

import "fmt"

func main() {
	fmt.Println(obterValorAprovad(6000))
	fmt.Println(obterValorAprovad(4000))
}

func obterValorAprovad(numero int) int {
	defer fmt.Println("fim!") // atrasa para o ultimo momento possivel no metodo
	if numero > 5000 {
		fmt.Println("Valor alto!")
		return 5000
	}
	fmt.Println("Valor baixo!")
	return 0
}
