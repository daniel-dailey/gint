package interpreter

import (
	"log"
)

type AST interface {
	visit() (interface{}, ReturnType)
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
	log.Printf("p.CurToken.Type(%s) != typ(%s)", p.CurToken.String(), typ.String())
	p.Error()
}

func (p *Parser) program() AST {
	p.consume(TOKEN_TYPE_PROGRAM)
	varNode := p.variable()
	progName := varNode.(*Var).val
	p.consume(TOKEN_TYPE_SEMICOLON)
	blockNode := p.block()
	programNode := InitProgram(progName, blockNode)
	p.consume(TOKEN_TYPE_DOT)
	return programNode
}

func (p *Parser) block() AST {
	declarationNodes := p.declarations()
	compoundStatementNode := p.compound()
	node := InitBlock(declarationNodes, compoundStatementNode)
	return node
}

func (p *Parser) declarations() []AST {
	dec := make([]AST, 0)
	if p.CurToken.Type == TOKEN_TYPE_VAR {
		p.consume(TOKEN_TYPE_VAR)
		for p.CurToken.Type == TOKEN_TYPE_ID {
			varDec := p.variableDeclaration()
			dec = append(dec, varDec...)
			p.consume(TOKEN_TYPE_SEMICOLON)
		}
	}
	return dec
}

func (p *Parser) variableDeclaration() []AST {
	varNodes := make([]*Var, 0)
	log.Println("Attempting make var: ", p.CurToken)
	varNodes = append(varNodes, NewVar(p.CurToken))
	p.consume(TOKEN_TYPE_ID)
	for p.CurToken.Type == TOKEN_TYPE_COMMA {
		p.consume(TOKEN_TYPE_COMMA)
		varNodes = append(varNodes, NewVar(p.CurToken))
		p.consume(TOKEN_TYPE_ID)
	}
	p.consume(TOKEN_TYPE_COLON)
	typeNode := p.typeSpec()
	varDeclarations := make([]AST, 0)
	for _, vn := range varNodes {
		log.Printf("var node name: %s", vn.token.String())
		varDeclarations = append(varDeclarations, InitVariableDeclaration(vn, typeNode))
	}
	return varDeclarations
}

func (p *Parser) typeSpec() AST {
	token := p.CurToken
	if token.Type == TOKEN_TYPE_INT {
		p.consume(TOKEN_TYPE_INT)
	} else {
		p.consume(TOKEN_TYPE_REAL)
	}
	return InitType(token)
}

func (p *Parser) compound() AST {
	p.consume(TOKEN_TYPE_BEGIN)
	statements := p.statements()
	p.consume(TOKEN_TYPE_END)
	c := NewCompound()
	c.children = append(c.children, statements...)
	return c
}

func (p *Parser) statements() []AST {
	statementNode := p.statement()
	res := make([]AST, 0)
	res = append(res, statementNode)
	for p.CurToken.Type == TOKEN_TYPE_SEMICOLON {
		p.consume(TOKEN_TYPE_SEMICOLON)
		res = append(res, p.statement())
	}
	return res
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

func (p *Parser) assign() AST {
	l := p.variable()
	t := p.CurToken
	p.consume(TOKEN_TYPE_ASSIGN)
	r := p.expression()
	node := NewAssign(l, r, t)
	return node
}

func (p *Parser) variable() AST {
	log.Println("var: ", p.CurToken)
	node := NewVar(p.CurToken)
	p.consume(TOKEN_TYPE_ID)
	return node
}

func (p *Parser) empty() AST {
	return NewNoOp()
}

func (p *Parser) expression() AST {
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

func (p *Parser) term() AST {
	factorNode := p.factor()
	if factorNode == nil {
		return nil
	}
	token := p.CurToken
	for token.Type == TOKEN_TYPE_MULTIPLICATION ||
		token.Type == TOKEN_TYPE_INTEGER_DIV ||
		token.Type == TOKEN_TYPE_FLOAT_DIV {
		p.consume(token.Type)
		otherFactorNode := p.factor()
		if factorNode == nil {
			log.Println("ERROR: node is nil!")
			return nil
		}
		binaryOperator := NewBinaryOp(factorNode, otherFactorNode, token)
		factorNode = binaryOperator
	}
	return factorNode
}

func (p *Parser) factor() AST {
	currentToken := p.CurToken
	switch currentToken.Type {
	case TOKEN_TYPE_ADDITION,
		TOKEN_TYPE_SUBTRACTION:
		p.consume(currentToken.Type)
		return NewUnaryOp(currentToken, p.factor())
	case TOKEN_TYPE_INTEGER_CONST,
		TOKEN_TYPE_REAL_CONST:
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
