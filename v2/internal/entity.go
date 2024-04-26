package internal

import "time"

type CheckingAccount struct {
	Id               Id
	Name             string
	PersonalDocument string
}

type Id struct {
	Id       int
	CreateAt time.Time
}
