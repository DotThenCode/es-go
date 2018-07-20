package main

const (
	TokenType_IDENTIFIER	= iota,
	TokenType_LPAREN		= iota,
	TokenType_RPAREN		= iota,
	TokenType_SEMI			= iota,
	TokenType_EOF			= iota,
)

type Token struct {
	Type int,
	Value string
}