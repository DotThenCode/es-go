package main

func IsAsciiLetter(char byte) bool {
	return (char >= 65 && char <= 90) || (char >= 97 && char <= 122)
}

func IsAsciiDigit(char byte) bool {
	return char >= 48 && char <= 57
}