package nodes

import "github.com/daniel-dailey/gint/interpreter/token"

type Assign struct {
	typ   TreeNodeType
	left  TreeNode
	right TreeNode
	op    *token.Token
}

func (a *Assign) getType() TreeNodeType {
	return a.typ
}

func (a *Assign) GetType() TreeNodeType {
	return a.getType()
}

func (a *Assign) GetLeft() TreeNode {
	return a.left
}

func (a *Assign) GetRight() TreeNode {
	return a.right
}

func NewAssign(l, r TreeNode, op *token.Token) *Assign {
	return &Assign{
		typ:   TreeNodeTypeAssign,
		left:  l,
		right: r,
		op:    op,
	}
}
