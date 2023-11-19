package interpreter

import "log"

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
	case TreeNodeTypeAssign:
		varName := node.(*Assign).left.(*Var).val
		varSymbol := stb.lookup(varName)
		if varSymbol == nil {
			log.Fatal("var symbol nil...")
		}
		stb.Visit(node.(*Assign).right)
	case TreeNodeTypeVar:
		varName := node.(*Var).val
		varSymbol := stb.lookup(varName)
		if varSymbol == nil {
			log.Fatalf("var name %s nil...", varName)
		}

	case TreeNodeTypeVariableDeclaration:
		log.Println(node)
		typeName := node.(*VariableDeclaration).typeNode.(*Type).val()
		symbol := stb.lookup(typeName.(string))
		if symbol == nil {
			log.Fatalln("treenode var declaration symbol lookup = nil")
		}
		typeSymbol := stb.lookup(typeName.(string)).(*BuiltInTypeSymbol)
		varName := node.(*VariableDeclaration).varNode.(*Var).val
		varSymbol := InitVarSymbol(varName, typeSymbol)
		stb.insert(varSymbol)
	}
}

func InitSymbolTableBuilder() *SymbolTableBuilder {
	return &SymbolTableBuilder{
		InitSymbolTable(),
	}
}
