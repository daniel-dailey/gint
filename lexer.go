package main

import (
	"log"
	"strconv"
)

type Lexer struct {
	buf     string
	pos     int
	curRune rune
}

const (
	SPACE_CHAR    = 0x20
	ZERO_CHAR     = 0x30
	NINE_CHAR     = 0x39
	LOWER_A_CHAR  = 0x61
	LOWER_Z_CHAR  = 0x7A
	HIGHER_A_CHAR = 0x41
	HIGHER_Z_CHAR = 0x5A
)

func (l *Lexer) isDigit() bool {
	return l.curRune >= ZERO_CHAR && l.curRune <= NINE_CHAR
}

func (l *Lexer) isLowerCase() bool {
	return l.curRune >= LOWER_A_CHAR && l.curRune <= LOWER_Z_CHAR
}

func (l *Lexer) isUpperCase() bool {
	return l.curRune >= HIGHER_A_CHAR && l.curRune <= HIGHER_Z_CHAR
}

func (l *Lexer) isAlnum() bool {
	return l.isDigit() || l.isLowerCase() || l.isUpperCase()
}

func (l *Lexer) peek() rune {
	if l.pos+1 > len(l.buf)-1 {
		return -1
	}
	return rune(l.buf[l.pos+1])
}

func (l *Lexer) keyword() *Token {
	res := ""
	for l.curRune != -1 && l.isAlnum() {
		res += string(l.curRune)
		l.advance()
	}
	tokenType := RESERVED_WORDS[res]
	return NewToken(tokenType, 0)
}

func (l *Lexer) int() int {
	ret := ""
	for l.curRune != -1 && l.isDigit() {
		ret += string(l.curRune)
		l.advance()
	}
	n, err := strconv.Atoi(ret)
	if err != nil {
		log.Fatal("Failed to convert int")
	}
	return n
}

func (l *Lexer) getNextToken() *Token {
	for l.curRune != -1 {
		for l.curRune == SPACE_CHAR {
			l.advance()
		}
		if l.isAlnum() {
			return l.keyword()
		}
		//Handle special runes
		switch l.curRune {
		case ':':
			if l.peek() == '=' {
				l.advance()
				l.advance()
				return NewToken(TOKEN_TYPE_ASSIGN, -1)
			}
		}
		if l.isDigit() {
			return NewToken(TOKEN_TYPE_INT, l.int())
		}
		if tt, ok := TOKEN_TYPES[l.curRune]; ok {
			l.advance()
			return NewToken(tt, -1)
		}
		l.Error()
	}
	return NewToken(TOKEN_TYPE_EOF, -1)
}

func (l *Lexer) Error() {
	log.Fatal("failed to parse input")
}

func (l *Lexer) advance() {
	l.pos++
	if l.pos > len(l.buf)-1 {
		l.curRune = -1
		return
	}
	l.curRune = rune(l.buf[l.pos])
}

func NewLexer(buf string) *Lexer {
	return &Lexer{
		buf:     buf,
		curRune: rune(buf[0]),
	}
}
