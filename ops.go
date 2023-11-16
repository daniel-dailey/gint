package main

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
