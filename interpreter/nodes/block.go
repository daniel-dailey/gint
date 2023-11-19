package nodes

type Block struct {
	typ               TreeNodeType
	declarations      []TreeNode
	compoundStatement TreeNode
}

func (b *Block) getType() TreeNodeType {
	return b.typ
}

func (b *Block) GetType() TreeNodeType {
	return b.getType()
}

func (b *Block) GetDeclarations() []TreeNode {
	return b.declarations
}

func (b *Block) GetCompoundStatement() TreeNode {
	return b.compoundStatement
}

func InitBlock(declarations []TreeNode, compoundStatement TreeNode) *Block {
	return &Block{
		typ:               TreeNodeTypeBlock,
		declarations:      declarations,
		compoundStatement: compoundStatement,
	}
}
