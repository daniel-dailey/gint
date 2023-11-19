package nodes

type Compound struct {
	typ      TreeNodeType
	children []TreeNode
}

func (c *Compound) getType() TreeNodeType {
	return c.typ
}

func (c *Compound) GetType() TreeNodeType {
	return c.getType()
}

func (c *Compound) GetChildren() []TreeNode {
	return c.children
}

func (c *Compound) SetChildren(children []TreeNode) {
	c.children = children
}

func NewCompound() *Compound {
	return &Compound{
		typ:      TreeNodeTypeCompound,
		children: make([]TreeNode, 0),
	}
}
