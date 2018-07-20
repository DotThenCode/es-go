func IsAsciiLetter(byte char) {
	return (char >= 65 && char <= 90) || (char >= 97 && char <= 122)
}

func IsAsciiDigit(byte char) {
	return char >= 48 && char <= 57
}