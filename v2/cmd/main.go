package main

import (
	"fmt"
	"v2/internal"
)

func main() {
	account1 := internal.CreateAccount("João do Sol", "1234567899")
	fmt.Println("Nova Conta Corrente Criada: ", account1)

	account2 := internal.CreateAccount("Maria da Neves", "987654321")
	fmt.Println("Nova Conta Corrente Criada: ", account2)
}
