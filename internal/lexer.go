package internal

import (
	"fmt"
	"github.com/RootControl/Compiler/internal/enums"
)

type Lexer struct {
	SourceCode string
	SourceCodeLenght uint32
	CurrentPosition uint32 
	CurrentChar string
}

func NewLexer(sourceCode string) (*Lexer, error) {
	if sourceCode == "" {
		return &Lexer{}, fmt.Errorf("no source code was provided")
	}

	return &Lexer {
		SourceCode: sourceCode + "\\0",
		SourceCodeLenght: uint32(len(sourceCode)),
		CurrentPosition: 0,
		CurrentChar: fmt.Sprintf("%c", sourceCode[0]),
	},
	nil
}

func (l *Lexer) nextChar() {
	l.CurrentPosition++

	if l.CurrentPosition >= l.SourceCodeLenght {
		l.CurrentChar = "\\0"
	} else {
		l.CurrentChar = fmt.Sprintf("%c", l.SourceCode[l.CurrentPosition])
		// fmt.Println("----------")
		// fmt.Printf("Current Char: %v\n", l.CurrentChar)
		// fmt.Println("----------")
	}
}

func (l *Lexer) peekNext() string {
	if l.CurrentPosition + 1 >= l.SourceCodeLenght {
		return "\n"
	}

	return fmt.Sprintf("%c", l.SourceCode[l.CurrentPosition + 1])
}

func (l *Lexer) skipWhitespace() {
	for l.CurrentChar == " " ||
		l.CurrentChar == "\t" ||
		l.CurrentChar == "\r" {
			l.nextChar()
		}
}

func (l *Lexer) skipComment() {
	if l.CurrentChar == "#" {
		for l.CurrentChar != "\n" {
			l.nextChar()
		}
		l.nextChar()
	}
}

func (l *Lexer) GetToken() (Token, error) {
	l.skipWhitespace()
	l.skipComment()

	var enumType int32

	switch l.CurrentChar {
	case "+":
		enumType = enums.Plus
	case "-":
		enumType = enums.Minus
	case "*":
		enumType = enums.Asterisk
	case "/":
		enumType = enums.Slash
	case "\n":
		enumType = enums.NewLine
	case "=":
		if l.peekNext() == "=" {
			enumType = enums.SameAs
			lastChar := l.CurrentChar
			l.nextChar()
			l.CurrentChar = lastChar	 + l.CurrentChar
		} else {
			enumType = enums.Equal
		}
	case ">":
		if l.peekNext() == "=" {
			enumType = enums.GreaterThanOrEqual
			lastChar := l.CurrentChar
			l.nextChar()
			l.CurrentChar = lastChar + l.CurrentChar
		} else {
			enumType = enums.GreaterThan
		}
	case "<":
		if l.peekNext() == "=" {
			enumType = enums.LessThanOrEqual
			lastChar := l.CurrentChar
			l.nextChar()
			l.CurrentChar = lastChar + l.CurrentChar
		} else {
			enumType = enums.LessThan
		}
	case "!":
		if l.peekNext() == "=" {
			enumType = enums.LessThanOrEqual
			lastChar := l.CurrentChar
			l.nextChar()
			l.CurrentChar = lastChar + l.CurrentChar
		} else {
			enumType = enums.ERROR
		}
	case "\\0":
		enumType = enums.EOF
	default:
		enumType = enums.ERROR
	}

	token, err := NewToken(enumType, l.CurrentChar)
	if err != nil {
		return token, fmt.Errorf("unable to create a token, because: %v", err)
	}
	
	l.nextChar()
	return token, nil
}
