package main

import (
	"log"
)

type NodeType int

const (
	NodeTypeBinaryOp NodeType = iota
	NodeTypeNumber
)

type AST interface {
	visit() int
}

type Parser struct {
	Lexer    *Lexer
	CurToken *Token
}

func (i *Parser) consume(typ TokenType) {
	if i.CurToken.Type == typ {
		i.CurToken = i.Lexer.getNextToken()
		return
	}
	i.Lexer.Error()
}

func (i *Parser) factor() AST {
	currentToken := i.CurToken
	switch currentToken.Type {
	case TOKEN_TYPE_ADDITION,
		TOKEN_TYPE_SUBTRACTION:
		i.consume(currentToken.Type)
		return NewUnaryOp(currentToken, i.factor())
	case TOKEN_TYPE_INT:
		i.consume(currentToken.Type)
		return NewNum(currentToken)
	case TOKEN_TYPE_LPAREN:
		i.consume(TOKEN_TYPE_LPAREN)
		node := i.expression()
		i.consume(TOKEN_TYPE_RPAREN)
		return node
	default:
		return nil
	}
}

func (i *Parser) term() AST {
	topNode := i.factor()
	if topNode == nil {
		return nil
	}
	for i.CurToken.Type == TOKEN_TYPE_MULTIPLICATION ||
		i.CurToken.Type == TOKEN_TYPE_DIVISION {
		tempCurrentToken := i.CurToken
		i.consume(tempCurrentToken.Type)
		factorNode := i.factor()
		if factorNode == nil {
			log.Println("ERROR: node is nil!")
			return nil
		}
		binaryOperator := NewBinaryOp(topNode, factorNode, tempCurrentToken)
		topNode = binaryOperator
	}
	return topNode
}

func (i *Parser) expression() AST {
	topNode := i.term()
	if topNode == nil {
		log.Println("ERROR: node nil!")
		return nil
	}
	for i.CurToken.Type == TOKEN_TYPE_ADDITION ||
		i.CurToken.Type == TOKEN_TYPE_SUBTRACTION {
		t := i.CurToken
		i.consume(t.Type)
		factorNode := i.term()
		if factorNode == nil {
			log.Printf("ERROR: node nil!")
			return nil
		}
		topNode = NewBinaryOp(topNode, factorNode, t)
	}
	return topNode
}

func (i *Parser) Parse() AST {
	return i.expression()
}

func NewParser(l *Lexer) *Parser {
	i := &Parser{
		Lexer:    l,
		CurToken: l.getNextToken(),
	}
	return i
}
