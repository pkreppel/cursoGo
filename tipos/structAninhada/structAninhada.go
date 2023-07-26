package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{1, 2, 12.10},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido %.2f", pedido.valorTotal())
}
