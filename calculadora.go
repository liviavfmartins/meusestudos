package main

import "fmt"

func main() {

	a := 10
	b := 20

	c := Add(a, b)

	fmt.Println("O resultado dessa soma é: ", c)

}

// Add é um função que recebe dois inteiros e retorna a soma entre eles
func Add(a int, b int) int {
	return a + b
}
