package main

import "fmt"

func main() {
	f1()
	f2("P1", "P2")

	r3, r4 := f3(), f4("P1", "P2")
	fmt.Println(r3, r4)

	p1, p2 := f5()
	fmt.Println(p1, p2)

}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("%s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno1", "Retorno2"
}
