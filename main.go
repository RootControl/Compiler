package main

import (
	"fmt"

	"github.com/RootControl/Compiler/internal"
	"github.com/RootControl/Compiler/internal/enums"
)

func main() {
	source := "#Testing <> != source code\n<= + - >= <>"
	lexer, _ := internal.NewLexer(source)

	token, err := lexer.GetToken()
	if err != nil {
		return 
	}
	

	for token.TypeValue != enums.EOF {
		fmt.Printf("%v\n", token.Character)
		token, err = lexer.GetToken()
		if err != nil {
			return 
		}

	}
}
