package interpreter

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

func (p *Parser) consume(typ TokenType) {
	if p.CurToken.Type == typ {
		p.CurToken = p.Lexer.getNextToken()
		return
	}
	log.Println("consume fail!!", p.CurToken.String())
	p.Error()
}

func (p *Parser) factor() AST {
	currentToken := p.CurToken
	log.Println("Factor: ", currentToken.String())
	switch currentToken.Type {
	case TOKEN_TYPE_ADDITION,
		TOKEN_TYPE_SUBTRACTION:
		p.consume(currentToken.Type)
		return NewUnaryOp(currentToken, p.factor())
	case TOKEN_TYPE_INT:
		p.consume(currentToken.Type)
		return NewNum(currentToken)
	case TOKEN_TYPE_LPAREN:
		p.consume(TOKEN_TYPE_LPAREN)
		node := p.expression()
		p.consume(TOKEN_TYPE_RPAREN)
		return node
	default:
		return p.variable()
	}
}

func (p *Parser) term() AST {
	factorNode := p.factor()
	if factorNode == nil {
		return nil
	}
	for p.CurToken.Type == TOKEN_TYPE_MULTIPLICATION ||
		p.CurToken.Type == TOKEN_TYPE_DIVISION {
		tempCurrentToken := p.CurToken
		p.consume(tempCurrentToken.Type)
		otherFactorNode := p.factor()
		if factorNode == nil {
			log.Println("ERROR: node is nil!")
			return nil
		}
		binaryOperator := NewBinaryOp(factorNode, otherFactorNode, tempCurrentToken)
		factorNode = binaryOperator
	}
	return factorNode
}

func (p *Parser) expression() AST {
	log.Println("expression...")
	expressionNode := p.term()
	if expressionNode == nil {
		log.Println("ERROR: node nil!")
		return nil
	}
	for p.CurToken.Type == TOKEN_TYPE_ADDITION ||
		p.CurToken.Type == TOKEN_TYPE_SUBTRACTION {
		t := p.CurToken
		p.consume(t.Type)
		factorNode := p.term()
		if factorNode == nil {
			log.Printf("ERROR: node nil!")
			return nil
		}
		expressionNode = NewBinaryOp(expressionNode, factorNode, t)
	}
	return expressionNode
}

func (p *Parser) variable() AST {
	log.Println("variable...")
	node := NewVar(p.CurToken)
	p.consume(TOKEN_TYPE_ID)
	return node
}

func (p *Parser) empty() AST {
	log.Println("NoOp...")
	return NewNoOp()
}

func (p *Parser) assign() AST {
	log.Println("assign...")
	l := p.variable()
	log.Println("got left var")
	t := p.CurToken
	log.Println("consume assign...")
	p.consume(TOKEN_TYPE_ASSIGN)
	log.Println("consumed assign...")
	r := p.expression()
	node := NewAssign(l, r, t)
	log.Printf("assign: ret node: %v\n", node)
	return node
}

func (p *Parser) statement() AST {
	log.Println("statement...")
	switch p.CurToken.Type {
	case TOKEN_TYPE_BEGIN:
		log.Println("begin...")
		return p.compound()
	case TOKEN_TYPE_ID:
		log.Println("id...")
		return p.assign()
	default:
		log.Println("empty...")
		return p.empty()
	}
}

func (p *Parser) statements() []AST {
	statementNode := p.statement()
	res := make([]AST, 0)
	res = append(res, statementNode)
	for p.CurToken.Type == TOKEN_TYPE_SEMICOLON {
		log.Println("seent semicolon...")
		p.consume(TOKEN_TYPE_SEMICOLON)
		res = append(res, p.statement())
	}
	if p.CurToken.Type == TOKEN_TYPE_ID {
		log.Println("curToken.Type == TOKEN_TYPE_ID")
		p.Error()
	}
	return res
}

func (p *Parser) compound() AST {
	log.Println("expecting begin token")
	p.consume(TOKEN_TYPE_BEGIN)
	log.Println("got begin token... next token: ", p.CurToken.String())
	statements := p.statements()
	log.Println("STATEMENTS: ", statements)
	p.consume(TOKEN_TYPE_END)
	log.Println("Consumed end...")
	c := NewCompound()
	c.children = append(c.children, statements...)
	return c
}

func (p *Parser) program() AST {
	log.Println("getting compound...")
	programNode := p.compound()
	log.Println("Got compound!")
	p.consume(TOKEN_TYPE_DOT)
	return programNode
}

func (p *Parser) Parse() AST {
	log.Println("parsing program...")
	node := p.program()
	log.Println("parsed program...", p.CurToken.String())
	if p.CurToken.Type != TOKEN_TYPE_EOF {
		log.Println("curToken.Type != TOKEN_TYPE_EOF")
		p.Error()
	}
	return node
}

func (p *Parser) Error() {
	log.Fatal("invalid syntax")
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		Lexer:    l,
		CurToken: l.getNextToken(),
	}
	return p
}
