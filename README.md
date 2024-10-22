# How to create interpreter in go by Thorsten Ball

## Main features
- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data struct

## REPL
- the lexer
- the parser
- the Abstract Syntax Tree (AST)
- the internal object system
- the evaluato

## The basic idea
Take a source code and parse then finally evaluate
Source Code -> Tokens -> AST (Abstract Syntax Tree)

Creating tokens is done via the Lexer which is a functionality that analyzes the source code and tokenize its parts which are then fed to the AST

Here’s an example. This is the input one gives to a lexer:

```go
"let x = 5 + 5;"

//And what comes out of the lexer looks kinda like this:

[
LET,
IDENTIFIER("x"),
EQUAL_SIGN,
INTEGER(5),
PLUS_SIGN,
INTEGER(5),
SEMICOLON
]
```
