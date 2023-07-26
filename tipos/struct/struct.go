package main

import "fmt"

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Caderno", 3.90, 0.06}
	fmt.Println(produto2, produto2.precoComDesconto())

}

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Metod: funcao com receiver(receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}
