package tests

import (
	"testing"
	"fmt"

	"github.com/RootControl/Compiler/internal"
) 

func TestLexerCreation(t *testing.T) {
	sourceCode := "testing source code"
	lexer, err := internal.NewLexer(sourceCode)

	if err != nil {
		t.Errorf('Test failed: %v', err)
		return
	}

	if lexer.SourceCodeLenght == len(sourceCode) &&
		lexer.CurrentPosition == 0 &&
		lexer.CurrentChar == sourceCode[0] {
			fmt.Printf("Test Passed: lexer created with the default values")
	} else {
		t.Errorf('Test failed: one or more values of the Lexer where not properly inicialized. \nValue: %v', lexer)
	}
}
