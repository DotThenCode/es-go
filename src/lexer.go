package main

import (
	"os"
	"io"
	"fmt"
	"unicode"
)

type Lexer struct {
	file *os.File
	currChar [1]byte
	isEOF bool
}

func (this Lexer) advance() {
	_, err := this.file.Read(this.currChar[:])
	
	if err != nil {
		if err == io.EOF {
			this.isEOF = true
			return
		} else {
			panic(err)
		}
	}
}

func (this Lexer) buildIdentifier() {
	for unicode.IsLetter(this.currChar) || 
	for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
}

func (this Lexer) GetNextToken() {

	if (this.isEOF) {
		this.advance()
		return Token{TokenType_EOF, ""}
	}

	if (this.currChar == '(') {
		this.advance()
		return Token{TokenType_LPAREN, ""}
	}

	if (this.currChar == ')') {
		this.advance()
		return Token(TokenType_RPAREN, "")
	}

	if (unicode.IsLetter(this.currChar)) {
		return this.buildIdentifier()
	}
}

func NewLexer(fileName string) Lexer {
	file, err := os.Open("./main.es")
	check(err)

	this := {file, [1]byte{0}, false}
	this.advance()

	return this
}