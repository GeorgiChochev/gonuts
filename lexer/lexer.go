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
	}

	l.readChar()
	return tok
}

// Utility methods

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
