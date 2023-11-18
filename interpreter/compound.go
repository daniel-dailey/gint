package interpreter

type Compound struct {
	typ      TreeNodeType
	children []TreeNode
}

func (c *Compound) getType() TreeNodeType {
	return c.typ
}

func NewCompound() *Compound {
	return &Compound{
		typ:      TreeNodeTypeCompound,
		children: make([]TreeNode, 0),
	}
}
