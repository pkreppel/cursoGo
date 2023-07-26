package main

import "fmt"

func inc1(n int) {
	n++
}

// revisao um ponteiro e representado por *
func inc2(n *int) {
	//revisao: * é usado para desreferenciar, ou seja,
	//ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Println(n)

	inc2(&n) //passa o endereço do ponteiro. Passagem por referencia
	fmt.Println(n)
}
