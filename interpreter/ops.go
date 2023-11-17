package interpreter

import (
	"log"
	"strconv"
)

/*
	UNARY OPERATOR
*/

type UnaryOperator struct {
	op   *Token
	expr AST
}

func (uo *UnaryOperator) visit() int {
	switch uo.op.Type {
	case TOKEN_TYPE_ADDITION:
		return uo.expr.visit()
	case TOKEN_TYPE_SUBTRACTION:
		return -uo.expr.visit()
	}
	return -1
}

func NewUnaryOp(token *Token, expr AST) *UnaryOperator {
	return &UnaryOperator{
		op:   token,
		expr: expr,
	}
}

/*
	BINARY OPERATOR
*/

type BinaryOperator struct {
	left  AST
	right AST
	op    *Token
}

func (bo *BinaryOperator) visit() int {
	if bo.op != nil {
		switch bo.op.Type {
		case TOKEN_TYPE_ADDITION:
			return bo.left.visit() + bo.right.visit()
		case TOKEN_TYPE_SUBTRACTION:
			return bo.left.visit() - bo.right.visit()
		case TOKEN_TYPE_MULTIPLICATION:
			return bo.left.visit() * bo.right.visit()
		case TOKEN_TYPE_DIVISION:
			return bo.left.visit() / bo.right.visit()
		}
	}
	return -1
}

func NewBinaryOp(l, r AST, op *Token) *BinaryOperator {
	return &BinaryOperator{
		left:  l,
		right: r,
		op:    op,
	}
}

type Compound struct {
	children []AST
}

func (c *Compound) visit() int {
	for _, child := range c.children {
		child.visit()
	}
	return 0
}

func NewCompound() *Compound {
	return &Compound{
		children: make([]AST, 0),
	}
}

type Assign struct {
	left  AST
	right AST
	op    *Token
}

func (a *Assign) visit() int {
	variableName := a.left.(*Var).val
	log.Println("assign visit: var name: ", variableName)
	GLOBAL_SCOPE[variableName] = a.right.visit()
	log.Println("visited...")
	return 0
}

func NewAssign(l, r AST, op *Token) *Assign {
	return &Assign{
		left:  l,
		right: r,
		op:    op,
	}
}

type Var struct {
	token *Token
	val   string
}

func (v *Var) visit() int {
	variableName := v.val
	if val, ok := GLOBAL_SCOPE[variableName]; ok {
		return val
	}
	s, _ := strconv.Atoi(v.val)
	// log.Fatalln("name err... ", variableName, v.token.Value)
	return s
}

func NewVar(t *Token) *Var {
	return &Var{
		token: t,
		val:   t.Value.(string),
	}
}

type NoOp struct {
}

func (no *NoOp) visit() int {
	return 0
}

func NewNoOp() *NoOp {
	return &NoOp{}
}
