package token

// String alias to define type of token
type TokenType string

// Lexical token with its type and literal value
type Token struct {
	Type    TokenType
	Literal string // Actual value/text
}

const (
	ILLEGAL = "ILLEGAL" // Token or character we don't know about
	EOF     = "EOF"     // End of file

	IDENT = "IDENT" // Identifiers like variable and function names
	INT   = "INT"   // Integer literals

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(identifier string) TokenType {
	if tokenType, ok := keywords[identifier]; ok {
		return tokenType
	}

	return IDENT
}
