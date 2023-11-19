package token

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
	TOKEN_TYPE_PROCEDURE
	TOKEN_TYPE_EOF
)

var RESERVED_WORDS = map[string]TokenType{
	"PROGRAM":   TOKEN_TYPE_PROGRAM,
	"program":   TOKEN_TYPE_PROGRAM,
	"VAR":       TOKEN_TYPE_VAR,
	"var":       TOKEN_TYPE_VAR,
	"DIV":       TOKEN_TYPE_INTEGER_DIV,
	"div":       TOKEN_TYPE_INTEGER_DIV,
	"INTEGER":   TOKEN_TYPE_INT,
	"integer":   TOKEN_TYPE_INT,
	"REAL":      TOKEN_TYPE_REAL,
	"real":      TOKEN_TYPE_REAL,
	"BEGIN":     TOKEN_TYPE_BEGIN,
	"begin":     TOKEN_TYPE_BEGIN,
	"END":       TOKEN_TYPE_END,
	"end":       TOKEN_TYPE_END,
	"PROCEDURE": TOKEN_TYPE_PROCEDURE,
	"procedure": TOKEN_TYPE_PROCEDURE,
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
	case TOKEN_TYPE_REAL:
		return "(float64)"
	case TOKEN_TYPE_ADDITION:
		return "+"
	case TOKEN_TYPE_SUBTRACTION:
		return "-"
	case TOKEN_TYPE_MULTIPLICATION:
		return "*"
	case TOKEN_TYPE_INTEGER_DIV:
		return "/"
	case TOKEN_TYPE_FLOAT_DIV:
		return "float64(/)"
	case TOKEN_TYPE_LPAREN:
		return "LPAREN"
	case TOKEN_TYPE_RPAREN:
		return "RPAREN"
	case TOKEN_TYPE_DOT:
		return "DOT"
	case TOKEN_TYPE_SEMICOLON:
		return "SEMICOLON"
	case TOKEN_TYPE_EOF:
		return "EOF"
	case TOKEN_TYPE_REAL_CONST:
		return "REAL"
	case TOKEN_TYPE_INTEGER_CONST:
		return "INTEGER"
	case TOKEN_TYPE_VAR:
		return "VAR"
	case TOKEN_TYPE_ASSIGN:
		return "="
	case TOKEN_TYPE_PROGRAM:
		return "PROGRAM"
	case TOKEN_TYPE_BEGIN:
		return "BEGIN"
	case TOKEN_TYPE_END:
		return "END"
	case TOKEN_TYPE_ID:
		return "ID"
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
