package main

import (
	"os"
	"io"
	"bytes"
	"fmt"
)

type Lexer struct {
	file *os.File
	currChar byte
	isEOF bool
}

func (this *Lexer) advance() {
	var b [1]byte
	_, err := this.file.Read(b[:])
	
	if err != nil {
		if err == io.EOF {
			this.isEOF = true
			return
		} else {
			panic(err)
		}
	}

	this.currChar = b[0]
}

func (this *Lexer) buildIdentifier() Token {
    var buffer bytes.Buffer

	for IsAsciiLetter(this.currChar) || IsAsciiDigit(this.currChar) {
		buffer.WriteByte(this.currChar)
		this.advance()
	}

	return Token{TokenType_IDENTIFIER, buffer.String()}
}

func (this *Lexer) buildString() Token {
	var buffer bytes.Buffer
	this.advance()

	for this.currChar != byte('"') {
		buffer.WriteByte(this.currChar)
		this.advance()
	}

	this.advance()
	return Token{TokenType_STRING, buffer.String()}
}

func (this *Lexer) GetNextToken() Token {
	for {
		if this.isEOF {
			return Token{TokenType_EOF, ""}
		}

		if this.currChar == byte(' ') || this.currChar == byte('\n') || this.currChar == byte('\t') {
			this.advance()
			continue
		}

		if this.currChar == byte('(') {
			this.advance()
			return Token{TokenType_LPAREN, ""}
		}

		if this.currChar == byte(')') {
			this.advance()
			return Token{TokenType_RPAREN, ""}
		}

		if this.currChar == byte(';') {
			this.advance()
			return Token{TokenType_SEMI, ""}
		}

		if this.currChar == byte('"') {
			return this.buildString()
		}

		if IsAsciiLetter(this.currChar) {
			return this.buildIdentifier()
		}

		panic(fmt.Sprintf("Illegal character: '%s' %d(%q)", string(this.currChar), this.currChar, this.currChar))
	}
}

func NewLexer(fileName string) Lexer {
	file, err := os.Open("./main.es")
	check(err)

	this := Lexer{file, 0, false}
	this.advance()

	return this
}