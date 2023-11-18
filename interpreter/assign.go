package interpreter

type Assign struct {
	typ   TreeNodeType
	left  TreeNode
	right TreeNode
	op    *Token
}

func (a *Assign) getType() TreeNodeType {
	return a.typ
}

func NewAssign(l, r TreeNode, op *Token) *Assign {
	return &Assign{
		typ:   TreeNodeTypeAssign,
		left:  l,
		right: r,
		op:    op,
	}
}
