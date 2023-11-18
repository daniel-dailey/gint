package interpreter

type Block struct {
	typ               TreeNodeType
	declarations      []TreeNode
	compoundStatement TreeNode
}

func (b *Block) getType() TreeNodeType {
	return b.typ
}

func InitBlock(declarations []TreeNode, compoundStatement TreeNode) *Block {
	return &Block{
		typ:               TreeNodeTypeBlock,
		declarations:      declarations,
		compoundStatement: compoundStatement,
	}
}
