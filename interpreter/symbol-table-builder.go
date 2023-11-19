package interpreter

import (
	"log"

	"github.com/daniel-dailey/gint/interpreter/nodes"
)

type SymbolTableBuilder struct {
	*SymbolTable
}

func (stb *SymbolTableBuilder) Visit(node nodes.TreeNode) {
	switch node.GetType() {
	case nodes.TreeNodeTypeBlock:
		b := node.(*nodes.Block)
		for _, declaration := range b.GetDeclarations() {
			stb.Visit(declaration)
		}
		stb.Visit(b.GetCompoundStatement())
	case nodes.TreeNodeTypeProgram:
		stb.Visit(node.(*nodes.Program).GetBlock())
	case nodes.TreeNodeTypeBinaryOp:
		bo := node.(*nodes.BinaryOperator)
		stb.Visit(bo.GetLeftNode())
		stb.Visit(bo.GetRightNode())
	case nodes.TreeNodeTypeUnaryOp:
		stb.Visit(node.(*nodes.UnaryOperator).GetExpression())
	case nodes.TreeNodeTypeCompound:
		for _, c := range node.(*nodes.Compound).GetChildren() {
			stb.Visit(c)
		}
	case nodes.TreeNodeTypeAssign:
		varName := node.(*nodes.Assign).GetLeft().(*nodes.Var).GetVal()
		varSymbol := stb.lookup(varName)
		if varSymbol == nil {
			log.Fatal("var symbol nil...")
		}
		stb.Visit(node.(*nodes.Assign).GetRight())
	case nodes.TreeNodeTypeVar:
		varName := node.(*nodes.Var).GetVal()
		varSymbol := stb.lookup(varName)
		if varSymbol == nil {
			log.Fatalf("var name %s nil...", varName)
		}

	case nodes.TreeNodeTypeVariableDeclaration:
		log.Println(node)
		typeName := node.(*nodes.VariableDeclaration).GetTypeNode().(*nodes.Type).Val()
		symbol := stb.lookup(typeName.(string))
		if symbol == nil {
			log.Fatalln("treenode var declaration symbol lookup = nil")
		}
		typeSymbol := stb.lookup(typeName.(string)).(*BuiltInTypeSymbol)
		varName := node.(*nodes.VariableDeclaration).GetVarNode().(*nodes.Var).GetVal()
		varSymbol := InitVarSymbol(varName, typeSymbol)
		stb.insert(varSymbol)
	}
}

func InitSymbolTableBuilder() *SymbolTableBuilder {
	return &SymbolTableBuilder{
		InitSymbolTable(),
	}
}
