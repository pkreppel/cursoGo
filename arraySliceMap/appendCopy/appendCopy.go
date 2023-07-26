package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3}
	var slice1 []int

	//array1 := append(array1[:], 1, 2, 3) append nao funciona com array
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) //nao expande o tamanho do slice. nao pode ser um array
	fmt.Println(slice2)

}
