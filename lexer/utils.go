package lexer

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}
