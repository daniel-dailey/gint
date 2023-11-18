package interpreter

type Num struct {
	typ   TreeNodeType
	token *Token
	val   interface{}
}

func (n *Num) getType() TreeNodeType {
	return n.typ
}

func NewNum(t *Token) *Num {
	return &Num{
		typ:   TreeNodeTypeNum,
		token: t,
		val:   t.Value,
	}
}
