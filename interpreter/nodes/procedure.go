package nodes

type ProcedureDeclaration struct {
	typ       TreeNodeType
	procName  string
	blockNode TreeNode
}

func (pd *ProcedureDeclaration) getType() TreeNodeType {
	return pd.typ
}

func (pd *ProcedureDeclaration) GetType() TreeNodeType {
	return pd.getType()
}

func InitProcedureDeclaration(procName string, b TreeNode) *ProcedureDeclaration {
	return &ProcedureDeclaration{
		typ:       TreeNodeTypeProcDec,
		procName:  procName,
		blockNode: b,
	}
}
