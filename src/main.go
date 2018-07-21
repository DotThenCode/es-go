package main

import (
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	lexer := NewLexer("./main.es")

	for tok := lexer.GetNextToken(); tok.Type != TokenType_EOF; tok = lexer.GetNextToken() {
		fmt.Printf("%v\n", tok)
	} 

	fmt.Printf("Finished!")
}