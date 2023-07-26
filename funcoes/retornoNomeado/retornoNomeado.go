package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {
	p1, p2 := troca(1, 2)
	fmt.Println(p1, p2)
}
