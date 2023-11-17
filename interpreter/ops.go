package interpreter

import "log"

/*
	UNARY OPERATOR
*/

type UnaryOperator struct {
	op   *Token
	expr AST
}

func (uo *UnaryOperator) visit() (interface{}, ReturnType) {
	switch uo.op.Type {
	case TOKEN_TYPE_ADDITION:
		v, _ := uo.expr.visit()
		return v.(int), TYPE_INT
	case TOKEN_TYPE_SUBTRACTION:
		v, _ := uo.expr.visit()
		return -v.(int), TYPE_INT
	}
	return nil, TYPE_NIL
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

func (bo *BinaryOperator) visit() (interface{}, ReturnType) {
	if bo.op != nil {
		log.Println("visit binop")
		l, _ := bo.left.visit()
		r, _ := bo.right.visit()
		switch bo.op.Type {
		case TOKEN_TYPE_ADDITION:
			return l.(int) + r.(int), TYPE_INT
		case TOKEN_TYPE_SUBTRACTION:
			return l.(int) - r.(int), TYPE_INT
		case TOKEN_TYPE_MULTIPLICATION:
			return l.(int) * r.(int), TYPE_INT
		case TOKEN_TYPE_INTEGER_DIV:
			return l.(int) / r.(int), TYPE_INT
		case TOKEN_TYPE_FLOAT_DIV:
			return l.(float64) - r.(float64), TYPE_FLOAT
		}
	}
	return nil, TYPE_NIL
}

func NewBinaryOp(l, r AST, op *Token) *BinaryOperator {
	return &BinaryOperator{
		left:  l,
		right: r,
		op:    op,
	}
}

type NoOp struct {
}

func (no *NoOp) visit() (interface{}, ReturnType) {
	return nil, TYPE_NIL
}

func NewNoOp() *NoOp {
	return &NoOp{}
}
