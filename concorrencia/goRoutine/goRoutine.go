package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Println("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Porque você não fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc", 1)
	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)

	go fale("Maria", "Entendi!!", 10)
	fale("João", "Parabéns", 5)

	fmt.Println("Fim!")
}
