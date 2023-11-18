package interpreter

type Var struct {
	typ   TreeNodeType
	token *Token
	val   string
}

func (v *Var) getType() TreeNodeType {
	return v.typ
}

func NewVar(t *Token) *Var {
	return &Var{
		typ:   TreeNodeTypeVar,
		token: t,
		val:   t.Value.(string),
	}
}
