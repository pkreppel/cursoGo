package main

import (
	"fmt"
	"time"
)

//Channel - é a forma de comuncação entre go routines
// channel é um tipo

func doisTresQuadroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuadroVezes(2, c)

	a, b := <-c, <-c // recebendo dados do canal

	fmt.Println(a, b)

	fmt.Println(<-c)
}
