package interpreter

type Program struct {
	typ   TreeNodeType
	name  string
	block TreeNode
}

func (p *Program) getType() TreeNodeType {
	return p.typ
}

func InitProgram(name string, block TreeNode) *Program {
	return &Program{
		typ:   TreeNodeTypeProgram,
		name:  name,
		block: block,
	}
}
