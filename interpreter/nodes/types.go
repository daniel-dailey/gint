package nodes

type TreeNodeType = int

const (
	TreeNodeTypeBinaryOp TreeNodeType = iota
	TreeNodeTypeUnaryOp
	TreeNodeTypeNoOp
	TreeNodeTypeProgram
	TreeNodeTypeBlock
	TreeNodeTypeVariableDeclaration
	TreeNodeTypeType
	TreeNodeTypeVar
	TreeNodeTypeCompound
	TreeNodeTypeAssign
	TreeNodeTypeNum
	TreeNodeTypeProcDec
)

type TreeNode interface {
	GetType() TreeNodeType
}
