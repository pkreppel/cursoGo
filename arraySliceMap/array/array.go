package main

import "fmt"

func main() {
	//homegenea (mesmo tipo) e estática (mesmo tamanh)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 5.6, 4.0
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Media %.2f\n ", media)
}
