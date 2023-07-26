package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro e", reflect.TypeOf(32000))

	//sem sinal (só positivos)... uint8, uint16, unit32, uint64
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	//com sinal... int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é ", i1)
	fmt.Println("O tipo da variável i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune é do tipo ", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 e ", i2)

	//numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println(" O literal de 49.99", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Ola sou uma string"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string e ", len(s1))

	//string multiplas linhas
	s2 := `Olá
	eu 
	sou uma
	string`
	fmt.Println("O tamanho da string e ", len(s2))

	//char e na verdade int32
	char := 'a'
	fmt.Println("O tipo de char e ", reflect.TypeOf(char))

}
