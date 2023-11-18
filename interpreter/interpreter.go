package interpreter

import (
	"fmt"
)

type Interpreter struct {
	parser *Parser
}

var GLOBAL_SCOPE = map[string]interface{}{}

func (i *Interpreter) visit(node TreeNode) interface{} {
	switch node.getType() {
	case TreeNodeTypeBlock:
		b := node.(*Block)
		for _, declaration := range b.declarations {
			i.visit(declaration)
		}
		i.visit(b.compoundStatement)
		return nil
	case TreeNodeTypeProgram:
		i.visit(node.(*Program).block)
		return nil
	case TreeNodeTypeBinaryOp:
		bo := node.(*BinaryOperator)
		if bo.op != nil {
			l := i.visit(bo.left)
			r := i.visit(bo.right)
			switch bo.op.Type {
			case TOKEN_TYPE_ADDITION:
				return l.(int) + r.(int)
			case TOKEN_TYPE_SUBTRACTION:
				return l.(int) - r.(int)
			case TOKEN_TYPE_MULTIPLICATION:
				return l.(int) * r.(int)
			case TOKEN_TYPE_INTEGER_DIV:
				return l.(int) / r.(int)
			case TOKEN_TYPE_FLOAT_DIV:
				return l.(float64) / r.(float64)
			}
		}
		return nil
	case TreeNodeTypeUnaryOp:
		uo := node.(*UnaryOperator)
		switch uo.op.Type {
		case TOKEN_TYPE_ADDITION:
			v := i.visit(uo.expr)
			return v.(int)
		case TOKEN_TYPE_SUBTRACTION:
			v := i.visit(uo.expr)
			return -v.(int)
		}
		return nil
	case TreeNodeTypeCompound:
		for _, c := range node.(*Compound).children {
			i.visit(c)
		}
		return nil
	case TreeNodeTypeVariableDeclaration:
		return nil
	case TreeNodeTypeAssign:
		varName := node.(*Assign).left.(*Var).val
		varValue := i.visit(node.(*Assign).right)
		GLOBAL_SCOPE[varName] = varValue
		return nil
	case TreeNodeTypeVar:
		return GLOBAL_SCOPE[node.(*Var).val]
	case TreeNodeTypeNum:
		return node.(*Num).val
	default:
		return nil
	}
}

func (i *Interpreter) Interpret() {
	rootNode := i.parser.Parse()
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
