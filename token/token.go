package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // for simplicity type is string but for performance it would be an improvement to use specific types like int or byte
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Indentifiers + Literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1234...

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "function"
	LET      = "Let"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
