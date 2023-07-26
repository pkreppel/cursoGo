package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i <= 10 { // nao tem while em go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par ", i)
		} else {
			fmt.Println("Impar ", i)
		}
	}

	for {
		fmt.Println("Infinito ...")
		time.Sleep(time.Second)
	}
}
