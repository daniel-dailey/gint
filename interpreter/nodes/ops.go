package nodes

import "github.com/daniel-dailey/gint/interpreter/token"

/*
	UNARY OPERATOR
*/

type UnaryOperator struct {
	typ  TreeNodeType
	op   *token.Token
	expr TreeNode
}

func (uo *UnaryOperator) getType() TreeNodeType {
	return uo.typ
}

func (uo *UnaryOperator) GetType() TreeNodeType {
	return uo.getType()
}

func (uo *UnaryOperator) GetExpression() TreeNode {
	return uo.expr
}

func (uo *UnaryOperator) GetOp() *token.Token {
	return uo.op
}

func NewUnaryOp(token *token.Token, expr TreeNode) *UnaryOperator {
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
	op    *token.Token
}

func (bo *BinaryOperator) getType() TreeNodeType {
	return bo.typ
}

func (bo *BinaryOperator) GetType() TreeNodeType {
	return bo.getType()
}

func (bo *BinaryOperator) GetOp() *token.Token {
	return bo.op
}

func (bo *BinaryOperator) GetLeftNode() TreeNode {
	return bo.left
}

func (bo *BinaryOperator) GetRightNode() TreeNode {
	return bo.right
}

func NewBinaryOp(l, r TreeNode, op *token.Token) *BinaryOperator {
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

func (no *NoOp) GetType() TreeNodeType {
	return no.getType()
}

func NewNoOp() *NoOp {
	return &NoOp{
		typ: TreeNodeTypeNoOp,
	}
}
