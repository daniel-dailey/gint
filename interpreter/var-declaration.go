package interpreter

type VariableDeclaration struct {
	typ      TreeNodeType
	varNode  TreeNode
	typeNode TreeNode
}

func (vd *VariableDeclaration) getType() TreeNodeType {
	return vd.typ
}

func InitVariableDeclaration(v, t TreeNode) *VariableDeclaration {
	return &VariableDeclaration{
		typ:      TreeNodeTypeVariableDeclaration,
		varNode:  v,
		typeNode: t,
	}
}
