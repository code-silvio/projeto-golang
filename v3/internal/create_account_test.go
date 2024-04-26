package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {

	//given
	name := "João do Sol"
	document := "98765434567"

	//then
	newAccount := CreateAccount(name, document)

	//verify
	assert.Equal(t, name, newAccount.Name, "The two words should be the same.")
	assert.Equal(t, document, newAccount.PersonalDocument, "The two words should be the same.")
}
