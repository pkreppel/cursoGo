package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	multi := exec(multiplicacao, 1, 2)
	//map, reduce, filter

	fmt.Println(multi)
}
