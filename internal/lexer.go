package internal

import (
	"fmt"
)

type Lexer struct {
	SourceCode string
	SourceCodeLenght uint
	CurrentPosition uint 
	CurrentChar rune
}

func NewLexer(sourceCode string) (Lexer, error) {
	if sourceCode == "" {
		return nil, fmt.Error("No source code was provided.")
	}

	return Lexer {
		SourceCode: sourceCode + '\0',
		SourceCodeLenght: len(sourceCode)
		CurrentPosition: 0,
		CurrentChar: rune(sourceCode[0]),
	},
	nil
}

func (l Lexer) nextChar() {
	l.CurrentPosition++

	if l.CurrentPosition >= l.SourceCodeLenght {
		l.CurrentChar = '\0'
	} else {
		l.CurrentChar = l.SourceCode[l.CurrentPosition]
	}
}

func (l Lexer) peekNext() rune {
	if l.CurrentPosition + 1 >= l.SourceCodeLenght {
		return '\0'
	}

	return l.SourceCode[l.CurrentPosition+1]
}
