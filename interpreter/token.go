package interpreter

import (
	"fmt"
)

type TokenType int

const (
	TOKEN_TYPE_INT TokenType = iota
	TOKEN_TYPE_REAL
	TOKEN_TYPE_INTEGER_CONST
	TOKEN_TYPE_REAL_CONST
	TOKEN_TYPE_ADDITION
	TOKEN_TYPE_SUBTRACTION
	TOKEN_TYPE_MULTIPLICATION
	TOKEN_TYPE_INTEGER_DIV
	TOKEN_TYPE_FLOAT_DIV
	TOKEN_TYPE_LPAREN
	TOKEN_TYPE_RPAREN
	TOKEN_TYPE_ID
	TOKEN_TYPE_ASSIGN
	TOKEN_TYPE_BEGIN
	TOKEN_TYPE_END
	TOKEN_TYPE_SEMICOLON
	TOKEN_TYPE_DOT
	TOKEN_TYPE_PROGRAM
	TOKEN_TYPE_VAR
	TOKEN_TYPE_COLON
	TOKEN_TYPE_COMMA
	TOKEN_TYPE_EOF
)

var RESERVED_WORDS = map[string]TokenType{
	"PROGRAM": TOKEN_TYPE_PROGRAM,
	"VAR":     TOKEN_TYPE_VAR,
	"DIV":     TOKEN_TYPE_INTEGER_DIV,
	"INTEGER": TOKEN_TYPE_INT,
	"REAL":    TOKEN_TYPE_REAL,
	"BEGIN":   TOKEN_TYPE_BEGIN,
	"END":     TOKEN_TYPE_END,
}

var TOKEN_TYPES = map[rune]TokenType{
	'+': TOKEN_TYPE_ADDITION,
	'-': TOKEN_TYPE_SUBTRACTION,
	'*': TOKEN_TYPE_MULTIPLICATION,
	'(': TOKEN_TYPE_LPAREN,
	')': TOKEN_TYPE_RPAREN,
	'.': TOKEN_TYPE_DOT,
	'/': TOKEN_TYPE_FLOAT_DIV,
	';': TOKEN_TYPE_SEMICOLON,
	',': TOKEN_TYPE_COMMA,
	':': TOKEN_TYPE_COLON,
}

func (tt TokenType) String() string {
	switch tt {
	case TOKEN_TYPE_INT:
		return "(int)"
	case TOKEN_TYPE_ADDITION:
		return "Addition"
	case TOKEN_TYPE_SUBTRACTION:
		return "Subtraction"
	case TOKEN_TYPE_MULTIPLICATION:
		return "Product"
	case TOKEN_TYPE_LPAREN:
		return "Lparen"
	case TOKEN_TYPE_RPAREN:
		return "Rparen"
	case TOKEN_TYPE_DOT:
		return "DOT"
	case TOKEN_TYPE_SEMICOLON:
		return "SEMICOLON"
	case TOKEN_TYPE_EOF:
		return "EOF"
	default:
		return ""
	}
}

type Token struct {
	Type  TokenType
	Value interface{}
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(%s, %d)", t.Type, t.Value)
}

func NewToken(t TokenType, v interface{}) *Token {
	return &Token{
		Type:  t,
		Value: v,
	}
}
