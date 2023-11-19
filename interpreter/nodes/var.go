package nodes

import "github.com/daniel-dailey/gint/interpreter/token"

type Var struct {
	typ   TreeNodeType
	token *token.Token
	val   string
}

func (v *Var) getType() TreeNodeType {
	return v.typ
}

func (v *Var) GetType() TreeNodeType {
	return v.getType()
}

func (v *Var) GetVal() string {
	return v.val
}

func NewVar(t *token.Token) *Var {
	return &Var{
		typ:   TreeNodeTypeVar,
		token: t,
		val:   t.Value.(string),
	}
}
