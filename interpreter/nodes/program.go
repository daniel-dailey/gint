package nodes

type Program struct {
	typ   TreeNodeType
	name  string
	block TreeNode
}

func (p *Program) getType() TreeNodeType {
	return p.typ
}

func (p *Program) GetType() TreeNodeType {
	return p.getType()
}

func (p *Program) GetBlock() TreeNode {
	return p.block
}

func InitProgram(name string, block TreeNode) *Program {
	return &Program{
		typ:   TreeNodeTypeProgram,
		name:  name,
		block: block,
	}
}
