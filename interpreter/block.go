package interpreter

type Block struct {
	declarations      []AST
	compoundStatement AST
}

func (b *Block) visit() interface{} {
	for _, declaration := range b.declarations {
		declaration.visit()
	}
	b.compoundStatement.visit()
	return nil
}

func InitBlock(declarations []AST, compoundStatement AST) *Block {
	return &Block{
		declarations:      declarations,
		compoundStatement: compoundStatement,
	}
}
