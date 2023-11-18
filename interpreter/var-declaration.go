package interpreter

type VariableDeclaration struct {
	varNode  AST
	typeNode AST
}

func (vd *VariableDeclaration) visit() interface{} {
	return nil
}

func InitVariableDeclaration(v, t AST) *VariableDeclaration {
	return &VariableDeclaration{
		varNode:  v,
		typeNode: t,
	}
}
