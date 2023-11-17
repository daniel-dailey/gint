package interpreter

type Block struct {
	declarations      []AST
	compoundStatement AST
}

func (b *Block) visit() (interface{}, ReturnType) {
	for _, declaration := range b.declarations {
		declaration.visit()
	}
	return nil, TYPE_NIL
}

func InitBlock(declarations []AST, compoundStatement AST) *Block {
	return &Block{
		declarations:      declarations,
		compoundStatement: compoundStatement,
	}
}
