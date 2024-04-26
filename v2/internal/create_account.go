package internal

import "time"

var id = 10000

func CreateAccount(name string, personalDocument string) CheckingAccount {
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
