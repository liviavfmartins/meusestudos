package main

import (
	"fmt"
)

func main() {
	var opcao int
	var pontos int = 0
        fmt.Println("Onde nasceu o Naruto?");
	fmt.Println("1) Aldeia da Folha");
	fmt.Println("2) Candido Mota");
	fmt.Println("3) Assis");
	fmt.Scanln(&opcao);
 
	if opcao == 1 {
		fmt.Println("Voce acertou!!!");	
		pontos = 1
	} else {
		fmt.Println("Voce errou!");
	}

 	fmt.Println("Fim do Jogo!");
	fmt.Printf("Voce fez %v pontos", pontos)
}
