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
		t.Errorf("Test failed: %v", err)
		return
	}

	if lexer.SourceCodeLenght == uint(len(sourceCode)) &&
		lexer.CurrentPosition == 0 &&
		lexer.CurrentChar == sourceCode[0] {
			fmt.Printf("Test Passed: lexer created with the default values")
	} else {
		t.Errorf("Test failed: one or more values of the Lexer where not properly inicialized. \nValue: %v", lexer)
	}
}

func TestLexerNextChar(t *testing.T) {
	sourceCode := "testing source code"
	lexer, err := internal.NewLexer(sourceCode)

	if err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}

	lexer.NextChar()

	if lexer.CurrentPosition == 1 &&
		lexer.CurrentChar == sourceCode[1] {
			fmt.Printf("Test Passed: lexer moved to the next char")
	} else {
		t.Errorf("Test failed: lexer did not move to the next char. \nValue: %v", lexer)
	}
}

func TestLexerPeekNext(t *testing.T) {
	sourceCode := "testing source code"
	lexer, err := internal.NewLexer(sourceCode)

	if err != nil {
		t.Errorf("Test failed: %v", err)
		return
	}

	if lexer.PeekNext() == sourceCode[1] {
		fmt.Printf("Test Passed: lexer peeked the next char")
	} else {
		t.Errorf("Test failed: lexer did not peek the next char. \nValue: %v", lexer)
	}
}
