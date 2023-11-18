package interpreter

/*
	UNARY OPERATOR
*/

type UnaryOperator struct {
	typ  TreeNodeType
	op   *Token
	expr TreeNode
}

func (uo *UnaryOperator) getType() TreeNodeType {
	return uo.typ
}

func NewUnaryOp(token *Token, expr TreeNode) *UnaryOperator {
	return &UnaryOperator{
		typ:  TreeNodeTypeUnaryOp,
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
	typ   TreeNodeType
	left  TreeNode
	right TreeNode
	op    *Token
}

func (bo *BinaryOperator) getType() TreeNodeType {
	return bo.typ
}

func NewBinaryOp(l, r TreeNode, op *Token) *BinaryOperator {
	return &BinaryOperator{
		typ:   TreeNodeTypeBinaryOp,
		left:  l,
		right: r,
		op:    op,
	}
}

type NoOp struct {
	typ TreeNodeType
}

func (no *NoOp) getType() TreeNodeType {
	return no.typ
}

func NewNoOp() *NoOp {
	return &NoOp{
		typ: TreeNodeTypeNoOp,
	}
}
