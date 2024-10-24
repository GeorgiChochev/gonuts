package lexer

import "exrun/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // byte is an alias for type uint8
}

// Conctructor function for creating a new Lexer instance
// and returns a pointer to *Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input} // Creates a new instance of Lexer using "composite literal??" to create a new struct
	// l := &Lexer{input: input}: This line creates a new instance of Lexer.
	// &Lexer{input: input} uses a syntax called a composite literal.
	// input: input initializes the input field of the Lexer with the string passed to the New function.
	// The & operator returns a pointer to this new Lexer instance, which is stored in the variable l.

	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// Utility methods

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // what is this
}

func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' <= ch || 'A' <= ch && 'Z' <= ch || '_' == ch
}

func (l *Lexer) readChar() {
	// Method Receiver: (l *Lexer) specifies that readChar is a method that operates on a pointer to a Lexer struct.
	// This is similar to a method in JavaScript that operates on an object instance.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
