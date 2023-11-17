package interpreter

import "log"

type Block struct {
	declarations      []AST
	compoundStatement AST
}

func (b *Block) visit() (interface{}, ReturnType) {
	log.Println("visit declarations...")
	for _, declaration := range b.declarations {
		declaration.visit()
	}
	b.compoundStatement.visit()
	return nil, TYPE_NIL
}

func InitBlock(declarations []AST, compoundStatement AST) *Block {
	return &Block{
		declarations:      declarations,
		compoundStatement: compoundStatement,
	}
}
