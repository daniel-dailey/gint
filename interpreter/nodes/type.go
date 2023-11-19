package nodes

import "github.com/daniel-dailey/gint/interpreter/token"

type ReturnType int

const (
	TYPE_FLOAT ReturnType = iota
	TYPE_INT
	TYPE_BOOL
	TYPE_NIL
)

type Type struct {
	typ   TreeNodeType
	token *token.Token
	value interface{}
}

func (t *Type) val() interface{} {
	return t.value
}

func (t *Type) Val() interface{} {
	return t.val()
}

func (t *Type) getType() TreeNodeType {
	return t.typ
}

func (t *Type) GetType() TreeNodeType {
	return t.getType()
}

func InitType(t *token.Token) *Type {
	return &Type{
		typ:   TreeNodeTypeType,
		token: t,
		value: t.Value,
	}
}
