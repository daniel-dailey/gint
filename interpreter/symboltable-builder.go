package interpreter

type SymbolTableBuilder struct {
	*SymbolTable
}

func (stb *SymbolTableBuilder) Visit(node TreeNode) {
	switch node.getType() {
	case TreeNodeTypeBlock:
		b := node.(*Block)
		for _, declaration := range b.declarations {
			stb.Visit(declaration)
		}
		stb.Visit(b.compoundStatement)
	case TreeNodeTypeProgram:
		stb.Visit(node.(*Program).block)
	case TreeNodeTypeBinaryOp:
		bo := node.(*BinaryOperator)
		stb.Visit(bo.left)
		stb.Visit(bo.right)
	case TreeNodeTypeUnaryOp:
		stb.Visit(node.(*UnaryOperator).expr)
	case TreeNodeTypeCompound:
		for _, c := range node.(*Compound).children {
			stb.Visit(c)
		}
	case TreeNodeTypeVariableDeclaration:
		typeName := node.(*VariableDeclaration).typeNode.(*Type).val()
		typeSymbol := stb.lookup(typeName.(string)).(*BuiltInTypeSymbol)
		varName := node.(*VariableDeclaration).varNode.(*Var).val
		varSymbol := InitVarSymbol(varName, typeSymbol)
		stb.insert(varSymbol)
	}
}

func InitSymbolTableBuilder() *SymbolTableBuilder {
	return &SymbolTableBuilder{
		&SymbolTable{symbols: make(map[string]SymbolIntf)},
	}
}
