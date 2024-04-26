package main

import (
	"fmt"
	"time"
)

type CheckingAccount struct {
	Id               Id
	Name             string
	PersonalDocument string
}

type Id struct {
	Id       int
	CreateAt time.Time
}

var id = 10000

func main() {
	accoun1 := createAccount("João do Sol", "1234567899")
	fmt.Println("Nova Conta Corrente Criada: ", accoun1)

	accoun2 := createAccount("Maria da Neves", "987654321")
	fmt.Println("Nova Conta Corrente Criada: ", accoun2)
}

func createAccount(name string, personalDocument string) CheckingAccount {
	return CheckingAccount{
		createId(),
		name,
		personalDocument,
	}
}

func createId() Id {
	id++
	return Id{
		id,
		time.Now(),
	}
}
