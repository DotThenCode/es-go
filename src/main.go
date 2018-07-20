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
	lexer.GetNextToken()

	fmt.Printf("HELLO")

	/*var b [1]byte

	f, err := os.Open("./main.es")
	check(err)
	
	for {
		_, err := f.Read(b[:])

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		fmt.Printf("%s", string(b[:]))
	}*/
}