package nodes

type VariableDeclaration struct {
	typ      TreeNodeType
	varNode  TreeNode
	typeNode TreeNode
}

func (vd *VariableDeclaration) getType() TreeNodeType {
	return vd.typ
}

func (vd *VariableDeclaration) GetType() TreeNodeType {
	return vd.getType()
}

func (vd *VariableDeclaration) GetTypeNode() TreeNode {
	return vd.typeNode
}

func (vd *VariableDeclaration) GetVarNode() TreeNode {
	return vd.varNode
}

func InitVariableDeclaration(v, t TreeNode) *VariableDeclaration {
	return &VariableDeclaration{
		typ:      TreeNodeTypeVariableDeclaration,
		varNode:  v,
		typeNode: t,
	}
}
