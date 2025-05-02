package main

import (
	"fmt"

	"github.com/RootControl/Compiler/internal"
	"github.com/RootControl/Compiler/internal/enums"
)

func main() {
	source := "+- \"This is a string\" # This is a comment!\n */"
	lexer, _ := internal.NewLexer(source)

	token, err := lexer.GetToken()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	

	for token.TypeValue != enums.EOF {
		fmt.Printf(`===========
 Char: %v
 Type: %v
===========
`, token.Character, token.TypeValue)

		token, err = lexer.GetToken()
		if err != nil {
			return 
		}

	}
}
