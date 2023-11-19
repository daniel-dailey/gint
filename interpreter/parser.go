package interpreter

import (
	"log"
)

type TreeNodeType = int

const (
	TreeNodeTypeBinaryOp TreeNodeType = iota
	TreeNodeTypeUnaryOp
	TreeNodeTypeNoOp
	TreeNodeTypeProgram
	TreeNodeTypeBlock
	TreeNodeTypeVariableDeclaration
	TreeNodeTypeType
	TreeNodeTypeVar
	TreeNodeTypeCompound
	TreeNodeTypeAssign
	TreeNodeTypeNum
	TreeNodeTypeProcDec
)

type TreeNode interface {
	getType() TreeNodeType
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

func (p *Parser) program() TreeNode {
	p.consume(TOKEN_TYPE_PROGRAM)
	varNode := p.variable()
	progName := varNode.(*Var).val
	p.consume(TOKEN_TYPE_SEMICOLON)
	blockNode := p.block()
	programNode := InitProgram(progName, blockNode)
	p.consume(TOKEN_TYPE_DOT)
	return programNode
}

func (p *Parser) block() TreeNode {
	declarationNodes := p.declarations()
	compoundStatementNode := p.compound()
	node := InitBlock(declarationNodes, compoundStatementNode)
	return node
}

func (p *Parser) declarations() []TreeNode {
	dec := make([]TreeNode, 0)
	if p.CurToken.Type == TOKEN_TYPE_VAR {
		p.consume(TOKEN_TYPE_VAR)
		for p.CurToken.Type == TOKEN_TYPE_ID {
			varDec := p.variableDeclaration()
			dec = append(dec, varDec...)
			p.consume(TOKEN_TYPE_SEMICOLON)
		}
	}
	for p.CurToken.Type == TOKEN_TYPE_PROCEDURE {
		p.consume(TOKEN_TYPE_PROCEDURE)
		procName := p.CurToken.Value.(string)
		p.consume(TOKEN_TYPE_ID)
		p.consume(TOKEN_TYPE_SEMICOLON)
		blockNode := p.block()
		procDec := InitProcedureDeclaration(procName, blockNode)
		dec = append(dec, procDec)
		p.consume(TOKEN_TYPE_SEMICOLON)
	}
	return dec
}

func (p *Parser) variableDeclaration() []TreeNode {
	varNodes := make([]*Var, 0)
	varNodes = append(varNodes, NewVar(p.CurToken))
	p.consume(TOKEN_TYPE_ID)
	for p.CurToken.Type == TOKEN_TYPE_COMMA {
		p.consume(TOKEN_TYPE_COMMA)
		varNodes = append(varNodes, NewVar(p.CurToken))
		p.consume(TOKEN_TYPE_ID)
	}
	p.consume(TOKEN_TYPE_COLON)
	typeNode := p.typeSpec()
	varDeclarations := make([]TreeNode, 0)
	for _, vn := range varNodes {
		varDeclarations = append(varDeclarations, InitVariableDeclaration(vn, typeNode))
	}
	return varDeclarations
}

func (p *Parser) typeSpec() TreeNode {
	token := p.CurToken
	if token.Type == TOKEN_TYPE_INT {
		p.consume(TOKEN_TYPE_INT)
	} else {
		p.consume(TOKEN_TYPE_REAL)
	}
	return InitType(token)
}

func (p *Parser) compound() TreeNode {
	p.consume(TOKEN_TYPE_BEGIN)
	statements := p.statements()
	p.consume(TOKEN_TYPE_END)
	c := NewCompound()
	c.children = append(c.children, statements...)
	return c
}

func (p *Parser) statements() []TreeNode {
	statementNode := p.statement()
	res := make([]TreeNode, 0)
	res = append(res, statementNode)
	for p.CurToken.Type == TOKEN_TYPE_SEMICOLON {
		p.consume(TOKEN_TYPE_SEMICOLON)
		res = append(res, p.statement())
	}
	return res
}

func (p *Parser) statement() TreeNode {
	switch p.CurToken.Type {
	case TOKEN_TYPE_BEGIN:
		return p.compound()
	case TOKEN_TYPE_ID:
		return p.assign()
	default:
		return p.empty()
	}
}

func (p *Parser) assign() TreeNode {
	l := p.variable()
	t := p.CurToken
	p.consume(TOKEN_TYPE_ASSIGN)
	r := p.expression()
	node := NewAssign(l, r, t)
	return node
}

func (p *Parser) variable() TreeNode {
	log.Println("var: ", p.CurToken)
	node := NewVar(p.CurToken)
	p.consume(TOKEN_TYPE_ID)
	return node
}

func (p *Parser) empty() TreeNode {
	return NewNoOp()
}

func (p *Parser) expression() TreeNode {
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

func (p *Parser) term() TreeNode {
	factorNode := p.factor()
	if factorNode == nil {
		return nil
	}
	for p.CurToken.Type == TOKEN_TYPE_MULTIPLICATION ||
		p.CurToken.Type == TOKEN_TYPE_INTEGER_DIV ||
		p.CurToken.Type == TOKEN_TYPE_FLOAT_DIV {
		t := p.CurToken
		p.consume(t.Type)
		otherFactorNode := p.factor()
		if factorNode == nil {
			log.Println("ERROR: node is nil!")
			return nil
		}
		binaryOperator := NewBinaryOp(factorNode, otherFactorNode, t)
		factorNode = binaryOperator
	}
	return factorNode
}

func (p *Parser) factor() TreeNode {
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

func (p *Parser) Parse() TreeNode {
	node := p.program()
	if p.CurToken.Type != TOKEN_TYPE_EOF {
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
