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
		nextToken = lexer.createTwoCharactersToken(token.EQ, '=')
		if nextToken.Type == token.ILLEGAL {
			nextToken = lexer.createToken(token.ASSIGN)
		}
	case '+':
		nextToken = lexer.createToken(token.PLUS)
	case '-':
		nextToken = lexer.createToken(token.MINUS)
	case '!':
		nextToken = lexer.createTwoCharactersToken(token.NOT_EQ, '=')
		if nextToken.Type == token.ILLEGAL {
			nextToken = lexer.createToken(token.BANG)
		}
	case '/':
		nextToken = lexer.createToken(token.SLASH)
	case '*':
		nextToken = lexer.createToken(token.ASTERISK)
	case '<':
		nextToken = lexer.createToken(token.LT)
	case '>':
		nextToken = lexer.createToken(token.GT)
	case ';':
		nextToken = lexer.createToken(token.SEMICOLON)
	case ',':
		nextToken = lexer.createToken(token.COMMA)
	case '(':
		nextToken = lexer.createToken(token.LPAREN)
	case ')':
		nextToken = lexer.createToken(token.RPAREN)
	case '{':
		nextToken = lexer.createToken(token.LBRACE)
	case '}':
		nextToken = lexer.createToken(token.RBRACE)
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
			nextToken = lexer.createToken(token.ILLEGAL)
		}
	}

	lexer.readNextChar()
	return nextToken
}

func (lexer *Lexer) createToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(lexer.currentChar),
	}
}

func (lexer *Lexer) createTwoCharactersToken(tokenType token.TokenType, expectedNextChar byte) token.Token {
	currentChar := lexer.currentChar

	if lexer.peekNextChar() == expectedNextChar {
		lexer.readNextChar()
		literal := string(currentChar) + string(lexer.currentChar)
		return token.Token{
			Type:    tokenType,
			Literal: literal,
		}
	}

	return lexer.createToken(token.ILLEGAL)
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

func (lexer *Lexer) peekNextChar() byte {
	if lexer.nextReadPosition >= len(lexer.sourceCode) {
		return 0
	} else {
		return lexer.sourceCode[lexer.nextReadPosition]
	}
}
