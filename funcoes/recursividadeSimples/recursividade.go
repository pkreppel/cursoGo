package main

import "fmt"

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	resultado2 := fatorial(10)
	fmt.Println(resultado2)

}

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}
