package main

import "fmt"

func main() {
	i := 1
	//Go nao tem aritmetica de ponteiros
	//p++
	var p *int = nil
	p = &i //pega o endereco da variavel e atribui ao ponteiro
	*p++

	fmt.Println(p, &i, *p, i)
}
