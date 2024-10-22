# Overview

My learning experience on Go on the go :)
while working on creating also along an interpreter

# Syntax

A file can begin with a package name making it exported essentially under that name

```go
package mypackage
```

Later imported to other files with

```go
import "<project>/mypackage"
```

Defining functions with lowercase vs uppercase is different

```go
func MyFunc() {} // Uppercase means the function is exported
func myFunc() {} // Lowercase means the function is private and scoped to only within its file
```

Current assumption of how this works:
assigns to 'l' a pointer to the Lexer struct and initializes Lexer's input prop essentially to the input argument with the remaining props being nil?

```go
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // byte is an alias for type uint8
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
```

I guess golang has 2 ways to declare variables
TODO: check to confirm if that is the case

```go
foo := 3 // this is const
var bar // this is mutable?
```

So when returning struct it requires to do it as token.Token{...} i dont get it
probably because the struct is a pointer this is the only way to assign the properties in that pointer space??

```go
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

```

TODO: how does that works
how Fatalf is exists as a method there idk
and the fmt strings dig how they work

```go
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
```

# Good to knows

Starting a project is easy
`go mod init <project>`
