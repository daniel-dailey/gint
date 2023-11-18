package interpreter

type Compound struct {
	children []AST
}

func (c *Compound) visit() interface{} {
	for _, child := range c.children {
		child.visit()
	}
	return nil
}

func NewCompound() *Compound {
	return &Compound{
		children: make([]AST, 0),
	}
}
