package lexer

import "monkeyplus-go/token"

type Lexer struct {
	sourceCode       string // Input  string to tokenize
	currentPosition  int    // Index of the current character being examined (index of currentChar)
	nextReadPosition int    // Index of the next character to read
	currentChar      byte   // Current character under examination
}

// Initializes a new lexer with given input string
func NewLexer(sourceCode string) *Lexer {
	lexer := &Lexer{sourceCode: sourceCode}
	lexer.readNextChar() // Load first character
	return lexer
}

func (lexer *Lexer) readNextChar() {
	if lexer.nextReadPosition >= len(lexer.sourceCode) {
		lexer.currentChar = 0 // Set currentChar to 0 (null character)
	} else {
		lexer.currentChar = lexer.sourceCode[lexer.nextReadPosition]
	}

	lexer.currentPosition = lexer.nextReadPosition
	lexer.nextReadPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var nextToken token.Token

	lexer.skipWhitespace()

	switch lexer.currentChar {
	case '=':
		nextToken = createToken(token.ASSIGN, lexer.currentChar)
	case '+':
		nextToken = createToken(token.PLUS, lexer.currentChar)
	case '-':
		nextToken = createToken(token.MINUS, lexer.currentChar)
	case '!':
		nextToken = createToken(token.BANG, lexer.currentChar)
	case '/':
		nextToken = createToken(token.SLASH, lexer.currentChar)
	case '*':
		nextToken = createToken(token.ASTERISK, lexer.currentChar)
	case '<':
		nextToken = createToken(token.LT, lexer.currentChar)
	case '>':
		nextToken = createToken(token.GT, lexer.currentChar)
	case ';':
		nextToken = createToken(token.SEMICOLON, lexer.currentChar)
	case ',':
		nextToken = createToken(token.COMMA, lexer.currentChar)
	case '(':
		nextToken = createToken(token.LPAREN, lexer.currentChar)
	case ')':
		nextToken = createToken(token.RPAREN, lexer.currentChar)
	case '{':
		nextToken = createToken(token.LBRACE, lexer.currentChar)
	case '}':
		nextToken = createToken(token.RBRACE, lexer.currentChar)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		if isLetter(lexer.currentChar) {
			nextToken.Literal = lexer.readIdentifier()
			nextToken.Type = token.LookupIdentifier(nextToken.Literal)

			return nextToken
		} else if isDigit(lexer.currentChar) {
			nextToken.Type = token.INT
			nextToken.Literal = lexer.readNumber()

			return nextToken
		} else {
			nextToken = createToken(token.ILLEGAL, lexer.currentChar)
		}
	}

	lexer.readNextChar()
	return nextToken
}

func createToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(character),
	}
}

func (lexer *Lexer) readIdentifier() string {
	currentPosition := lexer.currentPosition

	for isLetter(lexer.currentChar) {
		lexer.readNextChar()
	}

	return lexer.sourceCode[currentPosition:lexer.currentPosition]
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.currentChar == ' ' || lexer.currentChar == '\t' || lexer.currentChar == '\n' || lexer.currentChar == '\r' {
		lexer.readNextChar()
	}
}

func (lexer *Lexer) readNumber() string {
	currentPosition := lexer.currentPosition

	for isDigit(lexer.currentChar) {
		lexer.readNextChar()
	}

	return lexer.sourceCode[currentPosition:lexer.currentPosition]
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
