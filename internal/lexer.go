package internal

import (
	"fmt"
)

type Lexer struct {
	SourceCode string
	SourceCodeLenght uint
	CurrentPosition uint 
	CurrentChar byte
}

func NewLexer(sourceCode string) (Lexer, error) {
	if sourceCode == "" {
		return Lexer{}, fmt.Errorf("No source code was provided.")
	}

	return Lexer {
		SourceCode: sourceCode + "\n",
		SourceCodeLenght: uint(len(sourceCode)),
		CurrentPosition: 0,
		CurrentChar: sourceCode[0],
	},
	nil
}

func (l *Lexer) NextChar() {
	l.CurrentPosition++

	if l.CurrentPosition >= l.SourceCodeLenght {
		l.CurrentChar = '\n'
	} else {
		l.CurrentChar = l.SourceCode[l.CurrentPosition]
	}
}

func (l Lexer) PeekNext() byte {
	if l.CurrentPosition + 1 >= l.SourceCodeLenght {
		return '\n'
	}

	return l.SourceCode[l.CurrentPosition+1]
}
