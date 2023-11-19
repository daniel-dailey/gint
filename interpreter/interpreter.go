package interpreter

import (
	"fmt"

	"github.com/daniel-dailey/gint/interpreter/nodes"
	"github.com/daniel-dailey/gint/interpreter/token"
)

type Interpreter struct {
	parser *Parser
}

var GLOBAL_SCOPE = map[string]interface{}{}

func (i *Interpreter) visit(node nodes.TreeNode) interface{} {
	switch node.GetType() {
	case nodes.TreeNodeTypeBlock:
		b := node.(*nodes.Block)
		for _, declaration := range b.GetDeclarations() {
			i.visit(declaration)
		}
		i.visit(b.GetCompoundStatement())
		return nil
	case nodes.TreeNodeTypeProgram:
		i.visit(node.(*nodes.Program).GetBlock())
		return nil
	case nodes.TreeNodeTypeBinaryOp:
		bo := node.(*nodes.BinaryOperator)
		if bo.GetOp() != nil {
			l := i.visit(bo.GetLeftNode())
			r := i.visit(bo.GetRightNode())
			if l == nil || r == nil {
				return nil
			}
			switch bo.GetOp().Type {
			case token.TOKEN_TYPE_ADDITION:
				return l.(int) + r.(int)
			case token.TOKEN_TYPE_SUBTRACTION:
				return l.(int) - r.(int)
			case token.TOKEN_TYPE_MULTIPLICATION:
				return l.(int) * r.(int)
			case token.TOKEN_TYPE_INTEGER_DIV:
				return l.(int) / r.(int)
			case token.TOKEN_TYPE_FLOAT_DIV:
				return l.(float64) / r.(float64)
			}
		}
		return nil
	case nodes.TreeNodeTypeUnaryOp:
		uo := node.(*nodes.UnaryOperator)
		switch uo.GetOp().Type {
		case token.TOKEN_TYPE_ADDITION:
			v := i.visit(uo.GetExpression())
			return v.(int)
		case token.TOKEN_TYPE_SUBTRACTION:
			v := i.visit(uo.GetExpression())
			return -v.(int)
		}
		return nil
	case nodes.TreeNodeTypeCompound:
		for _, c := range node.(*nodes.Compound).GetChildren() {
			i.visit(c)
		}
		return nil
	case nodes.TreeNodeTypeVariableDeclaration:
		return nil
	case nodes.TreeNodeTypeAssign:
		varName := node.(*nodes.Assign).GetLeft().(*nodes.Var).GetVal()
		varValue := i.visit(node.(*nodes.Assign).GetRight())
		GLOBAL_SCOPE[varName] = varValue
		return nil
	case nodes.TreeNodeTypeVar:
		return GLOBAL_SCOPE[node.(*nodes.Var).GetVal()]
	case nodes.TreeNodeTypeNum:
		return node.(*nodes.Num).GetVal()
	default:
		return nil
	}
}

func (i *Interpreter) Interpret(rootNode nodes.TreeNode) {
	i.visit(rootNode)
	fmt.Println("======================================================")
	fmt.Println("=                      OUTPUT                        =")
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Println("======================================================")
	fmt.Println("=                    MEMORY MAP                      =")
	fmt.Println("======================================================")
	for k, v := range GLOBAL_SCOPE {
		fmt.Printf(">> %s = %d\n", k, v)
	}
	fmt.Println("======================================================")
}

func NewInterpreter(p *Parser) *Interpreter {
	return &Interpreter{
		parser: p,
	}
}
