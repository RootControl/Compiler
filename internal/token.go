package internal

import (
	"fmt"
)

type Token struct {
	TypeValue int32
	Character string
}

func NewToken(typeValue int32, character string) (Token, error) {
	token := Token{}

	if character == "" {
		return token, fmt.Errorf("character not defined")
	}

	if typeValue < -1 {
		return token, fmt.Errorf("Token type is invalid")
	}

	return Token {
		TypeValue: typeValue,
		Character: character,
	},
	nil
}
