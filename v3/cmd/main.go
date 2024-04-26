package main

import (
	"fmt"
	usecase "v3/internal"
)

func main() {
	account1 := usecase.CreateAccount("João do Sol", "1234567899")
	fmt.Println("Nova Conta Corrente Criada: ", account1)

	account2 := usecase.CreateAccount("Maria da Neves", "987654321")
	fmt.Println("Nova Conta Corrente Criada: ", account2)
}
