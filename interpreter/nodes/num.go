package nodes

import "github.com/daniel-dailey/gint/interpreter/token"

type Num struct {
	typ   TreeNodeType
	token *token.Token
	val   interface{}
}

func (n *Num) getType() TreeNodeType {
	return n.typ
}

func (n *Num) GetType() TreeNodeType {
	return n.getType()
}

func (n *Num) GetVal() interface{} {
	return n.val
}

func NewNum(t *token.Token) *Num {
	return &Num{
		typ:   TreeNodeTypeNum,
		token: t,
		val:   t.Value,
	}
}
