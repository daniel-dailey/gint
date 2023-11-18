package interpreter

type Compound struct {
	children []AST
}

func (c *Compound) visit() (interface{}, ReturnType) {
	for _, child := range c.children {
		child.visit()
	}
	return nil, TYPE_NIL
}

func NewCompound() *Compound {
	return &Compound{
		children: make([]AST, 0),
	}
}
