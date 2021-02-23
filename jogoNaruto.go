package main

import (
	"fmt"
)

func main() {
	var opcao int
	var pontos int = 0
    

    fmt.Println("Onde nasceu o Naruto?")
	fmt.Println("1) Aldeia da Folha")
	fmt.Println("2) Aldeia da Nevoa")
	fmt.Println("3) Aldeia da Pedra")
	fmt.Println("4) Aldeia do Som")
	fmt.Scanln(&opcao)
 
	if opcao == 1 {
		fmt.Println("Voce acertou!!!")
		pontos = 1
	} else {
		fmt.Println("Voce errou!")
		fmt.Println("Voce é um BURRO! Dá ZERO pra ele!!!!!")
	}









 	fmt.Println("Fim do Jogo!")
	fmt.Printf("Voce fez %v pontos", pontos)
}
