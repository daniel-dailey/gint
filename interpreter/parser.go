package interpreter

import (
	"log"

	"github.com/daniel-dailey/gint/interpreter/nodes"
	"github.com/daniel-dailey/gint/interpreter/token"
)

type Parser struct {
	Lexer    *Lexer
	CurToken *token.Token
}

func (p *Parser) consume(typ token.TokenType) {
	if p.CurToken.Type == typ {
		p.CurToken = p.Lexer.getNextToken()
		return
	}
	p.Error()
}

func (p *Parser) program() nodes.TreeNode {
	p.consume(token.TOKEN_TYPE_PROGRAM)
	varNode := p.variable()
	progName := varNode.(*nodes.Var).GetVal()
	p.consume(token.TOKEN_TYPE_SEMICOLON)
	blockNode := p.block()
	programNode := nodes.InitProgram(progName, blockNode)
	p.consume(token.TOKEN_TYPE_DOT)
	return programNode
}

func (p *Parser) block() nodes.TreeNode {
	declarationNodes := p.declarations()
	compoundStatementNode := p.compound()
	node := nodes.InitBlock(declarationNodes, compoundStatementNode)
	return node
}

func (p *Parser) declarations() []nodes.TreeNode {
	dec := make([]nodes.TreeNode, 0)
	if p.CurToken.Type == token.TOKEN_TYPE_VAR {
		p.consume(token.TOKEN_TYPE_VAR)
		for p.CurToken.Type == token.TOKEN_TYPE_ID {
			varDec := p.variableDeclaration()
			dec = append(dec, varDec...)
			p.consume(token.TOKEN_TYPE_SEMICOLON)
		}
	}
	for p.CurToken.Type == token.TOKEN_TYPE_PROCEDURE {
		p.consume(token.TOKEN_TYPE_PROCEDURE)
		procName := p.CurToken.Value.(string)
		p.consume(token.TOKEN_TYPE_ID)
		p.consume(token.TOKEN_TYPE_SEMICOLON)
		blockNode := p.block()
		procDec := nodes.InitProcedureDeclaration(procName, blockNode)
		dec = append(dec, procDec)
		p.consume(token.TOKEN_TYPE_SEMICOLON)
	}
	return dec
}

func (p *Parser) variableDeclaration() []nodes.TreeNode {
	varNodes := make([]*nodes.Var, 0)
	varNodes = append(varNodes, nodes.NewVar(p.CurToken))
	p.consume(token.TOKEN_TYPE_ID)
	for p.CurToken.Type == token.TOKEN_TYPE_COMMA {
		p.consume(token.TOKEN_TYPE_COMMA)
		varNodes = append(varNodes, nodes.NewVar(p.CurToken))
		p.consume(token.TOKEN_TYPE_ID)
	}
	p.consume(token.TOKEN_TYPE_COLON)
	typeNode := p.typeSpec()
	varDeclarations := make([]nodes.TreeNode, 0)
	for _, vn := range varNodes {
		varDeclarations = append(varDeclarations, nodes.InitVariableDeclaration(vn, typeNode))
	}
	return varDeclarations
}

func (p *Parser) typeSpec() nodes.TreeNode {
	t := p.CurToken
	if t.Type == token.TOKEN_TYPE_INT {
		p.consume(token.TOKEN_TYPE_INT)
	} else {
		p.consume(token.TOKEN_TYPE_REAL)
	}
	return nodes.InitType(t)
}

func (p *Parser) compound() nodes.TreeNode {
	p.consume(token.TOKEN_TYPE_BEGIN)
	statements := p.statements()
	p.consume(token.TOKEN_TYPE_END)
	c := nodes.NewCompound()
	children := c.GetChildren()
	children = append(children, statements...)
	c.SetChildren(children)
	return c
}

func (p *Parser) statements() []nodes.TreeNode {
	statementNode := p.statement()
	res := make([]nodes.TreeNode, 0)
	res = append(res, statementNode)
	for p.CurToken.Type == token.TOKEN_TYPE_SEMICOLON {
		p.consume(token.TOKEN_TYPE_SEMICOLON)
		res = append(res, p.statement())
	}
	return res
}

func (p *Parser) statement() nodes.TreeNode {
	switch p.CurToken.Type {
	case token.TOKEN_TYPE_BEGIN:
		return p.compound()
	case token.TOKEN_TYPE_ID:
		return p.assign()
	default:
		return p.empty()
	}
}

func (p *Parser) assign() nodes.TreeNode {
	l := p.variable()
	t := p.CurToken
	p.consume(token.TOKEN_TYPE_ASSIGN)
	r := p.expression()
	node := nodes.NewAssign(l, r, t)
	return node
}

func (p *Parser) variable() nodes.TreeNode {
	node := nodes.NewVar(p.CurToken)
	p.consume(token.TOKEN_TYPE_ID)
	return node
}

func (p *Parser) empty() nodes.TreeNode {
	return nodes.NewNoOp()
}

func (p *Parser) expression() nodes.TreeNode {
	expressionNode := p.term()
	if expressionNode == nil {
		log.Println("ERROR: node nil!")
		return nil
	}
	for p.CurToken.Type == token.TOKEN_TYPE_ADDITION ||
		p.CurToken.Type == token.TOKEN_TYPE_SUBTRACTION {
		t := p.CurToken
		p.consume(t.Type)
		factorNode := p.term()
		if factorNode == nil {
			log.Printf("ERROR: node nil!")
			return nil
		}
		expressionNode = nodes.NewBinaryOp(expressionNode, factorNode, t)
	}
	return expressionNode
}

func (p *Parser) term() nodes.TreeNode {
	factorNode := p.factor()
	if factorNode == nil {
		return nil
	}
	for p.CurToken.Type == token.TOKEN_TYPE_MULTIPLICATION ||
		p.CurToken.Type == token.TOKEN_TYPE_INTEGER_DIV ||
		p.CurToken.Type == token.TOKEN_TYPE_FLOAT_DIV {
		t := p.CurToken
		p.consume(t.Type)
		otherFactorNode := p.factor()
		if factorNode == nil {
			log.Println("ERROR: node is nil!")
			return nil
		}
		binaryOperator := nodes.NewBinaryOp(factorNode, otherFactorNode, t)
		factorNode = binaryOperator
	}
	return factorNode
}

func (p *Parser) factor() nodes.TreeNode {
	currentToken := p.CurToken
	switch currentToken.Type {
	case token.TOKEN_TYPE_ADDITION,
		token.TOKEN_TYPE_SUBTRACTION:
		p.consume(currentToken.Type)
		return nodes.NewUnaryOp(currentToken, p.factor())
	case token.TOKEN_TYPE_INTEGER_CONST,
		token.TOKEN_TYPE_REAL_CONST:
		p.consume(currentToken.Type)
		return nodes.NewNum(currentToken)
	case token.TOKEN_TYPE_LPAREN:
		p.consume(token.TOKEN_TYPE_LPAREN)
		node := p.expression()
		p.consume(token.TOKEN_TYPE_RPAREN)
		return node
	default:
		return p.variable()
	}
}

func (p *Parser) Parse() nodes.TreeNode {
	node := p.program()
	if p.CurToken.Type != token.TOKEN_TYPE_EOF {
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
