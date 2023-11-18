package interpreter

/*
	UNARY OPERATOR
*/

type UnaryOperator struct {
	op   *Token
	expr AST
}

func (uo *UnaryOperator) visit() interface{} {
	switch uo.op.Type {
	case TOKEN_TYPE_ADDITION:
		v := uo.expr.visit()
		return v.(int)
	case TOKEN_TYPE_SUBTRACTION:
		v := uo.expr.visit()
		return -v.(int)
	}
	return nil
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

type RealNumber interface {
	int | float64
}

type BinaryOperator struct {
	left  AST
	right AST
	op    *Token
}

func (bo *BinaryOperator) visit() interface{} {
	if bo.op != nil {
		l := bo.left.visit()
		r := bo.right.visit()
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

func (no *NoOp) visit() interface{} {
	return nil
}

func NewNoOp() *NoOp {
	return &NoOp{}
}
