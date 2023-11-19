package interpreter

type ProcedureDeclaration struct {
	typ       TreeNodeType
	procName  string
	blockNode TreeNode
}

func (pd *ProcedureDeclaration) getType() TreeNodeType {
	return pd.typ
}

func InitProcedureDeclaration(procName string, b TreeNode) *ProcedureDeclaration {
	return &ProcedureDeclaration{
		typ:       TreeNodeTypeProcDec,
		procName:  procName,
		blockNode: b,
	}
}
